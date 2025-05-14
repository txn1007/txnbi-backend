package model

import "time"

const (
	TableNameOperationLog   = "operation_log"
	TableNameRole           = "role"
	TableNamePermission     = "permission"
	TableNameRolePermission = "role_permission"
	TableNameUserRole       = "user_role"
	TableNameInviteCode     = "invite_code"
)

// OperationLog 操作日志表
type OperationLog struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id"`                                                                // ID
	UserID      int64     `gorm:"column:userId;type:bigint;not null;comment:用户ID;index:idx_userId" json:"userId"`                                         // 用户ID
	UserName    string    `gorm:"column:userName;type:varchar(64);comment:用户名" json:"userName"`                                                           // 用户名
	UserAccount string    `gorm:"column:userAccount;type:varchar(64);comment:用户账号" json:"userAccount"`                                                    // 用户账号
	Operation   string    `gorm:"column:operation;type:varchar(128);not null;comment:操作;index:idx_operation" json:"operation"`                            // 操作
	Method      string    `gorm:"column:method;type:varchar(16);not null;comment:请求方法" json:"method"`                                                     // 请求方法
	Path        string    `gorm:"column:path;type:varchar(128);comment:请求路径" json:"path"`                                                                 // 请求路径
	Params      string    `gorm:"column:params;type:text;comment:请求参数" json:"params"`                                                                     // 请求参数
	IP          string    `gorm:"column:ip;type:varchar(64);comment:IP地址" json:"ip"`                                                                      // IP地址
	Status      int8      `gorm:"column:status;type:tinyint;not null;comment:操作状态;index:idx_status" json:"status"`                                        // 操作状态：0-成功，1-失败
	CreateTime  time.Time `gorm:"column:createTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间;index:idx_createTime" json:"createTime"` // 创建时间
}

// TableName OperationLog's table name
func (*OperationLog) TableName() string {
	return TableNameOperationLog
}

// Role 角色表
type Role struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id"`                                           // ID
	Name        string    `gorm:"column:name;type:varchar(32);not null;comment:角色名称;uniqueIndex:uk_name" json:"name"`                // 角色名称
	Description string    `gorm:"column:description;type:varchar(128);comment:角色描述" json:"description"`                              // 角色描述
	CreateTime  time.Time `gorm:"column:createTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime  time.Time `gorm:"column:updateTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"` // 更新时间
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}

// Permission 权限表
type Permission struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id"`                                           // ID
	Name        string    `gorm:"column:name;type:varchar(64);not null;comment:权限名称;uniqueIndex:uk_name" json:"name"`                // 权限名称
	Description string    `gorm:"column:description;type:varchar(128);comment:权限描述" json:"description"`                              // 权限描述
	Type        string    `gorm:"column:type;type:varchar(16);not null;comment:权限类型;index:idx_type" json:"type"`                     // 权限类型：view-查看，edit-编辑，delete-删除
	CreateTime  time.Time `gorm:"column:createTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime  time.Time `gorm:"column:updateTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"` // 更新时间
}

// TableName Permission's table name
func (*Permission) TableName() string {
	return TableNamePermission
}

// RolePermission 角色权限关联表
type RolePermission struct {
	ID           uint64 `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id"`                                                                                             // ID
	RoleID       uint64 `gorm:"column:roleId;type:bigint unsigned;not null;comment:角色ID;uniqueIndex:uk_role_permission,priority:1" json:"roleId"`                                    // 角色ID
	PermissionID uint64 `gorm:"column:permissionId;type:bigint unsigned;not null;comment:权限ID;uniqueIndex:uk_role_permission,priority:2;index:idx_permissionId" json:"permissionId"` // 权限ID
}

// TableName RolePermission's table name
func (*RolePermission) TableName() string {
	return TableNameRolePermission
}

// UserRole 用户角色关联表
type UserRole struct {
	ID     uint64 `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id"`                                                                     // ID
	UserID uint64 `gorm:"column:userId;type:bigint unsigned;not null;comment:用户ID;uniqueIndex:uk_user_role,priority:1" json:"userId"`                  // 用户ID
	RoleID uint64 `gorm:"column:roleId;type:bigint unsigned;not null;comment:角色ID;uniqueIndex:uk_user_role,priority:2;index:idx_roleId" json:"roleId"` // 角色ID
}

// TableName UserRole's table name
func (*UserRole) TableName() string {
	return TableNameUserRole
}

// InviteCode 邀请码表
type InviteCode struct {
	ID         uint64     `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id"`                                           // ID
	Code       string     `gorm:"column:code;type:varchar(32);not null;comment:邀请码;uniqueIndex:uk_code" json:"code"`                 // 邀请码
	MaxUses    int        `gorm:"column:maxUses;type:int;not null;default:1;comment:最大使用次数" json:"maxUses"`                          // 最大使用次数，0表示不限制
	UsedCount  int        `gorm:"column:usedCount;type:int;not null;default:0;comment:已使用次数" json:"usedCount"`                       // 已使用次数
	Status     int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:状态;index:idx_status" json:"status"`           // 状态：0-有效，1-无效
	ExpireTime *time.Time `gorm:"column:expireTime;comment:过期时间;index:idx_expireTime" json:"expireTime"`                             // 过期时间，NULL表示永不过期
	CreateTime time.Time  `gorm:"column:createTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateTime time.Time  `gorm:"column:updateTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"` // 更新时间
}

// TableName InviteCode's table name
func (*InviteCode) TableName() string {
	return TableNameInviteCode
}
