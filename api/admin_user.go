package api

import "time"

// AdminUserListReq 管理员获取用户列表请求
type AdminUserListReq struct {
	Page     int    `form:"page" json:"page" binding:"required,min=1"`
	PageSize int    `form:"pageSize" json:"pageSize" binding:"required,min=1,max=100"`
	Keyword  string `form:"keyword" json:"keyword"`
}

// AdminUserListResp 管理员获取用户列表响应
type AdminUserListResp struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	Total      int64        `json:"total"`
	Users      []UserInfoV1 `json:"users"`
}

// UserInfoV1 用户信息（管理员视角）
type UserInfoV1 struct {
	ID          int64     `json:"id"`
	UserAccount string    `json:"userAccount"`
	UserName    string    `json:"userName"`
	UserAvatar  string    `json:"userAvatar"`
	UserRole    string    `json:"userRole"`
	UserStatus  int8      `json:"userStatus"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

// AdminUserDetailReq 管理员获取用户详情请求
type AdminUserDetailReq struct {
	UserID int64 `form:"userId" json:"userId" binding:"required,min=1"`
}

// AdminUserDetailResp 管理员获取用户详情响应
type AdminUserDetailResp struct {
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	User       UserInfoV1 `json:"user"`
}

// AdminCreateUserReq 管理员创建用户请求
type AdminCreateUserReq struct {
	UserAccount  string `json:"userAccount" binding:"required,min=6,max=16"`
	UserPassword string `json:"userPassword" binding:"required,min=8,max=24"`
	UserRole     string `json:"userRole" binding:"required,oneof=admin user"`
}

// AdminCreateUserResp 管理员创建用户响应
type AdminCreateUserResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	UserID     int64  `json:"userId"`
}

// AdminUpdateUserReq 管理员更新用户请求
type AdminUpdateUserReq struct {
	UserID       int64  `json:"userId" binding:"required,min=1"`
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	UserName     string `json:"userName"`
	UserAvatar   string `json:"userAvatar"`
	UserRole     string `json:"userRole" binding:"omitempty,oneof=admin user"`
}

// AdminUpdateUserResp 管理员更新用户响应
type AdminUpdateUserResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// AdminDeleteUserReq 管理员删除用户请求
type AdminDeleteUserReq struct {
	UserID int64 `json:"userId" binding:"required,min=1"`
}

// AdminDeleteUserResp 管理员删除用户响应
type AdminDeleteUserResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// AdminDisableUserReq 管理员禁用/启用用户请求
type AdminDisableUserReq struct {
	UserID int64 `json:"userId" binding:"required,min=1"`
	Status int   `json:"status" binding:"required,oneof=0 1"` // 0-正常, 1-禁用
}

// AdminDisableUserResp 管理员禁用/启用用户响应
type AdminDisableUserResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
