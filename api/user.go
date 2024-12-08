package api

import (
	"time"
)

type UserLoginReq struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserLoginResp struct {
	StatusCode int    `json:"statusCode" form:"statusCode"`
	Message    string `json:"message" form:"message"`
	Token      string `json:"token" form:"token"`
}

type UserRegisterReq struct {
	Account    string `json:"account" form:"account" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	InviteCode string `json:"inviteCode" form:"inviteCode" binding:"required"`
}

type UserRegisterResp struct {
	// required: true
	StatusCode int `json:"statusCode" form:"statusCode"`
	// required: true
	Message string `json:"message" form:"message"`
}

type CurrentUserDetailReq struct {
	Token string `json:"token" form:"token" binding:"required"`
}

type CurrentUserDetailResp struct {
	// required: true
	StatusCode int `json:"statusCode" form:"statusCode"`
	// required: true
	Message string `json:"message" form:"message"`

	UserInfoV0 UserInfoV0 `json:"userInfoV0" form:"userInfoV0"`
}

type UserInfoV0 struct {
	ID          int64     `json:"id"`          // id
	UserAccount string    `json:"userAccount"` // 账号
	UserName    string    `json:"userName"`    // 用户昵称
	UserAvatar  string    `json:"userAvatar"`  // 用户头像
	UserRole    string    `json:"userRole"`    // 用户角色：user/admin
	CreateTime  time.Time `json:"createTime"`  // 创建时间
	UpdateTime  time.Time `json:"updateTime"`  // 更新时间
}

type UserLoginOutReq struct {
	Token string `json:"token" form:"token" binding:"required"`
}

type UserLoginOutResp struct {
	// required: true
	StatusCode int `json:"statusCode" form:"statusCode"`
	// required: true
	Message string `json:"message" form:"message"`
}
