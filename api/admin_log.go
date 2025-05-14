package api

import "time"

// AdminLogListReq 管理员获取日志列表请求
type AdminLogListReq struct {
	Page      int       `form:"page" json:"page" binding:"required,min=1"`
	PageSize  int       `form:"pageSize" json:"pageSize" binding:"required,min=1,max=100"`
	Keyword   string    `form:"keyword" json:"keyword"`
	StartTime time.Time `form:"startTime" json:"startTime"`
	EndTime   time.Time `form:"endTime" json:"endTime"`
}

// AdminLogListResp 管理员获取日志列表响应
type AdminLogListResp struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Total      int64          `json:"total"`
	Logs       []OperationLog `json:"logs"`
}

// OperationLog 操作日志信息
type OperationLog struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	UserName    string    `json:"userName"`
	UserAccount string    `json:"userAccount"`
	Operation   string    `json:"operation"`
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	IP          string    `json:"ip"`
	CreateTime  time.Time `json:"createTime"`
}

// AdminLogDetailReq 管理员获取日志详情请求
type AdminLogDetailReq struct {
	LogID int64 `form:"logId" json:"logId" binding:"required,min=1"`
}

// AdminLogDetailResp 管理员获取日志详情响应
type AdminLogDetailResp struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	Log        OperationLog `json:"log"`
}

// AdminCreateLogReq 管理员创建日志请求
type AdminCreateLogReq struct {
	UserID      int64  `json:"userId" binding:"required,min=1"`
	Operation   string `json:"operation" binding:"required"`
	Method      string `json:"method" binding:"required"`
	Path        string `json:"path" binding:"required"`
	IP          string `json:"ip" binding:"required"`
}

// AdminCreateLogResp 管理员创建日志响应
type AdminCreateLogResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	LogID      int64  `json:"logId"`
}

// AdminUpdateLogReq 管理员更新日志请求
type AdminUpdateLogReq struct {
	LogID     int64  `json:"logId" binding:"required,min=1"`
	Operation string `json:"operation"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	IP        string `json:"ip"`
}

// AdminUpdateLogResp 管理员更新日志响应
type AdminUpdateLogResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// AdminDeleteLogReq 管理员删除日志请求
type AdminDeleteLogReq struct {
	LogID int64 `json:"logId" binding:"required,min=1"`
}

// AdminDeleteLogResp 管理员删除日志响应
type AdminDeleteLogResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// AdminBatchDeleteLogReq 管理员批量删除日志请求
type AdminBatchDeleteLogReq struct {
	LogIDs []int64 `json:"logIds" binding:"required,min=1"`
}

// AdminBatchDeleteLogResp 管理员批量删除日志响应
type AdminBatchDeleteLogResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}