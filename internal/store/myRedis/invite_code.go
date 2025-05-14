package myRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"txnbi-backend/internal/model"
)

const (
	inviteCodePrefix = "invite_code:"
	inviteCodeList   = "invite_code:list"
)

// SetInviteCode 设置邀请码
func SetInviteCode(code string, maxUses int, expireTime time.Time) error {
	inviteCode := model.InviteCode{
		Code:       code,
		MaxUses:    maxUses,
		UsedCount:  0,
		ExpireTime: &expireTime,
		CreateTime: time.Now(),
	}

	// 序列化邀请码
	data, err := json.Marshal(inviteCode)
	if err != nil {
		return err
	}

	// 存储邀请码
	key := inviteCodePrefix + code
	err = Cli.Set(context.Background(), key, string(data), time.Until(expireTime)).Err()
	if err != nil {
		return err
	}

	// 将邀请码添加到列表中
	return Cli.SAdd(context.Background(), inviteCodeList, code).Err()
}

// IsInviteCode 检查邀请码是否存在
func IsInviteCode(code string) (bool, error) {
	key := inviteCodePrefix + code
	exists, err := Cli.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

// GetInviteCode 获取邀请码信息
func GetInviteCode(code string) (*model.InviteCode, error) {
	key := inviteCodePrefix + code
	data, err := Cli.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	var inviteCode model.InviteCode
	err = json.Unmarshal([]byte(data), &inviteCode)
	if err != nil {
		return nil, err
	}

	return &inviteCode, nil
}

// UpdateInviteCodeUsage 更新邀请码使用次数
func UpdateInviteCodeUsage(code string) error {
	inviteCode, err := GetInviteCode(code)
	if err != nil {
		return err
	}

	// 检查使用次数是否达到上限
	if inviteCode.MaxUses > 0 && inviteCode.UsedCount >= inviteCode.MaxUses {
		return fmt.Errorf("邀请码已达到最大使用次数")
	}

	// 更新使用次数
	inviteCode.UsedCount++

	// 序列化邀请码
	data, err := json.Marshal(inviteCode)
	if err != nil {
		return err
	}

	// 存储邀请码
	key := inviteCodePrefix + code
	return Cli.Set(context.Background(), key, string(data), time.Until(*inviteCode.ExpireTime)).Err()
}

// DeleteInviteCode 删除邀请码
func DeleteInviteCode(code string) error {
	key := inviteCodePrefix + code

	// 删除邀请码
	err := Cli.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}

	// 从列表中删除邀请码
	return Cli.SRem(context.Background(), inviteCodeList, code).Err()
}

// GetAllInviteCodes 获取所有邀请码
func GetAllInviteCodes() ([]model.InviteCode, error) {
	// 获取所有邀请码
	codes, err := Cli.SMembers(context.Background(), inviteCodeList).Result()
	if err != nil {
		return nil, err
	}

	var inviteCodes []model.InviteCode
	for _, code := range codes {
		inviteCode, err := GetInviteCode(code)
		if err != nil {
			continue
		}

		inviteCodes = append(inviteCodes, *inviteCode)
	}

	return inviteCodes, nil
}

// GetInviteCodeUsage 获取邀请码使用情况
func GetInviteCodeUsage(code string) (int, error) {
	inviteCode, err := GetInviteCode(code)
	if err != nil {
		return 0, err
	}

	return inviteCode.UsedCount, nil
}
