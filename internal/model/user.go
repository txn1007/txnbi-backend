package model

import "time"

const TableNameUser = "users"

// User 用户表
type User struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id"`                                                                // ID
	UserAccount  string    `gorm:"column:userAccount;type:varchar(64);not null;comment:账号;uniqueIndex:uk_userAccount" json:"userAccount"`                  // 账号
	UserPassword string    `gorm:"column:userPassword;type:varchar(128);not null;comment:密码" json:"userPassword"`                                          // 密码
	UserName     string    `gorm:"column:userName;type:varchar(64);comment:用户昵称" json:"userName"`                                                          // 用户昵称
	UserAvatar   string    `gorm:"column:userAvatar;type:varchar(255);comment:用户头像" json:"userAvatar"`                                                     // 用户头像
	UserRole     string    `gorm:"column:userRole;type:varchar(32);not null;default:user;comment:用户角色;index:idx_userRole" json:"userRole"`                 // 用户角色：user/admin/operator
	UserStatus   int8      `gorm:"column:userStatus;type:tinyint;not null;default:0;comment:用户状态;index:idx_userStatus" json:"userStatus"`                  // 用户状态：0-正常，1-禁用
	LastLogin    time.Time `gorm:"column:lastLogin;comment:最后登录时间" json:"lastLogin"`                                                                       // 最后登录时间
	CreateTime   time.Time `gorm:"column:createTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间;index:idx_createTime" json:"createTime"` // 创建时间
	UpdateTime   time.Time `gorm:"column:updateTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"`                      // 更新时间
	IsDelete     int8      `gorm:"column:isDelete;type:tinyint;not null;default:0;comment:是否删除" json:"isDelete"`                                           // 是否删除
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
