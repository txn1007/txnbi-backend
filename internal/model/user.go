package model

import (
	"time"
)

const TableNameUser = "user"

// User 用户
type User struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                        // id
	UserAccount  string    `gorm:"column:userAccount;not null;comment:账号" json:"userAccount"`                           // 账号
	UserPassword string    `gorm:"column:userPassword;not null;comment:密码" json:"userPassword"`                         // 密码
	UserName     string    `gorm:"column:userName;comment:用户昵称" json:"userName"`                                        // 用户昵称
	UserAvatar   string    `gorm:"column:userAvatar;comment:用户头像" json:"userAvatar"`                                    // 用户头像
	UserRole     string    `gorm:"column:userRole;not null;default:user;comment:用户角色：user/admin" json:"userRole"`       // 用户角色：user/admin
	CreateTime   time.Time `gorm:"column:createTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime   time.Time `gorm:"column:updateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"` // 更新时间
	IsDelete     int32     `gorm:"column:isDelete;not null;comment:是否删除" json:"isDelete"`                               // 是否删除
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
