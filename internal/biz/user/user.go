package biz

import (
	"context"
	"errors"
	"fmt"
	"time"
	"txnbi-backend/conf"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store"
	myRedis2 "txnbi-backend/internal/store/myRedis"
	"txnbi-backend/pkg/jwt"
	"txnbi-backend/tool/encry"

	"gorm.io/gorm"
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
	// 检查用户是否被禁用
	if ac.UserStatus == 1 {
		return "", fmt.Errorf("account is disabled")
	}
	// 更新最后登录时间
	ac.LastLogin = time.Now()
	store.UpdateUser(context.Background(), ac)
	// 将token记录到redis中，标记为最新的token
	token = jwt.SignForUser(ac.ID, ac.UserAccount, conf.JWTCfg.SignKey)
	myRedis2.SetUserToken(ac.ID, token)
	return token, nil
}

func UserRegister(ctx context.Context, account, password, inviteCode string) error {
	// 检查邀请码是否存在
	is, err := store.IsInviteCode(inviteCode)
	if !is {
		return fmt.Errorf("邀请码不存在！")
	}
	if err != nil {
		return err
	}

	// 更新邀请码使用次数
	err = myRedis2.UpdateInviteCodeUsage(inviteCode)
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
	user := model.User{
		UserAccount:  account,
		UserPassword: encry.EncodeByMd5(password),
		UserRole:     "user",
		UserStatus:   0,
		LastLogin:    time.Now(),
	}
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
	myRedis2.DeleteUserToken(userID)
	return nil
}

// 运营端接口实现方案

// 根据您的需求，我需要为系统添加运营端的接口，包含用户管理、图表管理、日志管理、用户/运营权限管理和邀请码功能。下面我将为您实现这些功能。

// 首先，我们需要在 `user.go` 文件中添加相关的运营端接口。我将按照您的要求实现各个功能模块。

// 用户管理接口

// GetUserList 获取用户列表
func GetUserList(ctx context.Context, page, pageSize int, keyword string) ([]model.User, int64, error) {
	return store.GetUserList(ctx, page, pageSize, keyword)
}

// CreateUser 创建用户
func CreateUser(ctx context.Context, user model.User) (int64, error) {
	// 检查账号是否已存在
	ac, err := store.GetUserByAccount(user.UserAccount)
	if ac != nil && err == nil {
		return 0, fmt.Errorf("账号已存在")
	}

	// 加密用户密码
	user.UserPassword = encry.EncodeByMd5(user.UserPassword)

	// 创建用户
	return store.CreateUser(user)
}

// UpdateUser 更新用户信息
func UpdateUser(ctx context.Context, user model.User) error {
	// 检查用户是否存在
	_, err := store.GetUserByID(user.ID)
	if err != nil {
		return fmt.Errorf("用户不存在")
	}

	// 如果更新密码，需要加密
	if user.UserPassword != "" {
		user.UserPassword = encry.EncodeByMd5(user.UserPassword)
	}

	return store.UpdateUser(ctx, &user)
}

// DeleteUser 删除用户
func DeleteUser(ctx context.Context, userID int64) error {
	// 检查用户是否存在
	_, err := store.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("用户不存在")
	}

	return store.DeleteUser(ctx, userID)
}

// DisableUser 禁用/启用用户
func DisableUser(ctx context.Context, userID int64, isDisabled bool) error {
	// 检查用户是否存在
	user, err := store.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("用户不存在")
	}

	// 更新用户状态
	var status int8
	if isDisabled {
		status = 1
	} else {
		status = 0
	}
	user.UserStatus = status
	return store.UpdateUser(ctx, user)
}

// 图表管理接口

// GetChartList 获取图表列表
func GetChartList(ctx context.Context, page, pageSize int, keyword string) ([]model.Chart, int64, error) {
	return store.GetChartList(ctx, page, pageSize, keyword)
}

// GetChartDetail 获取图表详情
func GetChartDetail(ctx context.Context, chartID int64) (*model.Chart, error) {
	return store.GetChartByID(ctx, chartID)
}

// CreateChartByAdmin 管理员创建图表
func CreateChartByAdmin(ctx context.Context, chart model.Chart) error {
	chart.CreateTime = time.Now()
	chart.UpdateTime = time.Now()
	return store.CreateChart(ctx, chart)
}

// UpdateChartByAdmin 管理员更新图表
func UpdateChartByAdmin(ctx context.Context, chart model.Chart) error {
	// 检查图表是否存在
	_, err := store.GetChartByID(ctx, chart.ID)
	if err != nil {
		return fmt.Errorf("图表不存在")
	}

	chart.UpdateTime = time.Now()
	return store.UpdateChart(ctx, &chart)
}

// DeleteChartByAdmin 管理员删除图表
func DeleteChartByAdmin(ctx context.Context, chartID int64) error {
	// 检查图表是否存在
	_, err := store.GetChartByID(ctx, chartID)
	if err != nil {
		return fmt.Errorf("图表不存在")
	}

	return store.DeleteChart(ctx, chartID)
}

// 日志管理接口

// GetLogList 获取日志列表
func GetLogList(ctx context.Context, page, pageSize int, keyword string, startTime, endTime time.Time) ([]model.OperationLog, int64, error) {
	return store.GetLogList(ctx, page, pageSize, keyword, startTime, endTime)
}

// DeleteLog 删除日志
func DeleteLog(ctx context.Context, logID int64) error {
	return store.DeleteLog(ctx, logID)
}

// DeleteLogBatch 批量删除日志
func DeleteLogBatch(ctx context.Context, logIDs []int64) error {
	return store.DeleteLogBatch(ctx, logIDs)
}

// 运营权限管理接口

// GetRoleList 获取角色列表
func GetRoleList(ctx context.Context) ([]model.Role, error) {
	return store.GetRoleList(ctx)
}

// GetPermissionList 获取权限列表
func GetPermissionList(ctx context.Context) ([]model.Permission, error) {
	return store.GetPermissionList(ctx)
}

// AssignRoleToUser 为用户分配角色
func AssignRoleToUser(ctx context.Context, userID int64, roleID int64) error {
	// 检查用户是否存在
	_, err := store.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("用户不存在")
	}

	// 检查角色是否存在
	_, err = store.GetRoleByID(ctx, roleID)
	if err != nil {
		return fmt.Errorf("角色不存在")
	}

	return store.AssignRoleToUser(ctx, userID, roleID)
}

// UpdateUserRole 更新用户角色
func UpdateUserRole(ctx context.Context, userID int64, roleID int64) error {
	// 检查用户是否存在
	_, err := store.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("用户不存在")
	}

	// 检查角色是否存在
	_, err = store.GetRoleByID(ctx, roleID)
	if err != nil {
		return fmt.Errorf("角色不存在")
	}

	return store.UpdateUserRole(ctx, userID, roleID)
}

// 邀请码管理接口

// CreateInviteCode 创建邀请码
func CreateInviteCode(ctx context.Context, code string, maxUses int, expireTime time.Time) error {
	// 检查邀请码是否已存在
	exists, _ := store.IsInviteCode(code)
	if exists {
		return fmt.Errorf("邀请码已存在")
	}

	return store.SetInviteCode(code, maxUses, expireTime)
}

// GetInviteCodeList 获取邀请码列表
func GetInviteCodeList(ctx context.Context) ([]model.InviteCode, error) {
	return myRedis2.GetAllInviteCodes()
}

// UpdateInviteCode 更新邀请码
func UpdateInviteCode(ctx context.Context, code string, maxUses int, expireTime time.Time) error {
	// 检查邀请码是否存在
	exists, _ := myRedis2.IsInviteCode(code)
	if !exists {
		return fmt.Errorf("邀请码不存在")
	}

	// 删除旧的邀请码
	err := myRedis2.DeleteInviteCode(code)
	if err != nil {
		return err
	}

	// 创建新的邀请码
	return myRedis2.SetInviteCode(code, maxUses, expireTime)
}

// DeleteInviteCode 删除邀请码
func DeleteInviteCode(ctx context.Context, code string) error {
	// 检查邀请码是否存在
	exists, _ := myRedis2.IsInviteCode(code)
	if !exists {
		return fmt.Errorf("邀请码不存在")
	}

	return myRedis2.DeleteInviteCode(code)
}

// GetInviteCodeUsage 获取邀请码使用情况
func GetInviteCodeUsage(ctx context.Context, code string) (int, error) {
	// 检查邀请码是否存在
	exists, _ := myRedis2.IsInviteCode(code)
	if !exists {
		return 0, fmt.Errorf("邀请码不存在")
	}

	return myRedis2.GetInviteCodeUsage(code)
}
