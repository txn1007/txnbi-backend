package biz

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"txnbi-backend/conf"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store"
	"txnbi-backend/pkg/jwt"
	"txnbi-backend/pkg/myRedis"
	"txnbi-backend/tool/encry"
)

func UserLogin(account string, password string) (token string, err error) {
	// 根据账号获取用户
	ac, err := store.GetUserByAccount(account)
	// 检查账号是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("account not exist")
	}
	// 检查密码是否正确
	if ac.UserPassword != encry.EncodeByMd5(password) {
		return "", fmt.Errorf("password error")
	}
	// 将token记录到redis中，标记为最新的token
	token = jwt.SignForUser(ac.ID, ac.UserAccount, conf.JWTCfg.SignKey)
	myRedis.SetUserToken(ac.ID, token)
	return token, nil
}

func UserRegister(ctx context.Context, account, password, inviteCode string) error {
	// 检查邀请码是否存在
	is, err := myRedis.IsInviteCode(inviteCode)
	if !is {
		return fmt.Errorf("邀请码不存在！")
	}
	if err != nil {
		return err
	}
	ac, err := store.GetUserByAccount(account)
	if ac != nil && err == nil {
		return fmt.Errorf("account exist")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 创建用户,并加密用户密码
	user := model.User{UserAccount: account, UserPassword: encry.EncodeByMd5(password), UserRole: "user"}
	userID, err := store.CreateUser(user)
	if err != nil {
		return err
	}

	// 为新用户创建一张测试表
	chart := model.Chart{
		Goal:       "分析一下网站用户量数据趋势",
		Name:       "xx网站用户量表(创建账号时生成)",
		UserID:     userID,
		ChartType:  "折线图",
		GenChart:   `{"legend": {"data": []},"grid": {"left": "3%","right": "4%","bottom": "3%","containLabel": true},"xAxis": {"type": "category","boundaryGap": false,"data": ["1 号","2 号","3 号","4 号","5 号","6 号","7 号"]},"yAxis": {"type": "value"},"series": [{"name": "用户数","type": "line","data": [10,20,30,90,0,10,20]}]}`,
		GenResult:  "通过分析可知，网站用户量整体呈波动趋势。其中 4 号用户数达到峰值 90，而 5 号骤降至 0。其余日期用户数相对较为平稳，整体变化较为明显，需要进一步分析 4 号用户数暴增以及 5 号骤降的原因，以便更好地优化网站运营策略。",
		Status:     "succeed",
		IsDelete:   0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err = store.CreateChart(ctx, chart)
	if err != nil {
		return err
	}
	return nil
}

func CurrentUserDetail(userID int64) (*model.User, error) {
	return store.GetUserByID(userID)
}

func UserLoginOut(userID int64) error {
	myRedis.DeleteUserToken(userID)
	return nil
}
