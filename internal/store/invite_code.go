package store

import (
	"fmt"
	"time"
	"txnbi-backend/internal/model"
)

// SetInviteCode 设置邀请码
func SetInviteCode(code string, maxUses int, expireTime time.Time) error {
	inviteCode := model.InviteCode{
		Code:       code,
		MaxUses:    maxUses,
		UsedCount:  0,
		Status:     0,
		ExpireTime: &expireTime,
		CreateTime: time.Now(),
	}

	// 存储邀请码
	result := DB.Create(&inviteCode)
	return result.Error
}

// IsInviteCode 检查邀请码是否存在
func IsInviteCode(code string) (bool, error) {
	var count int64
	result := DB.Model(&model.InviteCode{}).Where("code = ?", code).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

// GetInviteCode 获取邀请码信息
func GetInviteCode(code string) (*model.InviteCode, error) {
	var inviteCode model.InviteCode
	result := DB.Where("code = ?", code).First(&inviteCode)
	if result.Error != nil {
		return nil, result.Error
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
	result := DB.Model(&model.InviteCode{}).Where("code = ?", code).Update("used_count", inviteCode.UsedCount+1)
	return result.Error
}

// DeleteInviteCode 删除邀请码
func DeleteInviteCode(code string) error {
	result := DB.Where("code = ?", code).Delete(&model.InviteCode{})
	return result.Error
}

// GetAllInviteCodes 获取所有邀请码
func GetAllInviteCodes() ([]model.InviteCode, error) {
	var inviteCodes []model.InviteCode
	result := DB.Find(&inviteCodes)
	if result.Error != nil {
		return nil, result.Error
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
