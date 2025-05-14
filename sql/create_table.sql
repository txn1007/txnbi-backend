-- 创建库
create database if not exists txnbi;

-- 切换库
use txnbi;

-- 用户表
create table if not exists user
(
    id           bigint auto_increment comment 'id' primary key,
    userAccount  varchar(256)                           not null comment '账号',
    userPassword varchar(512)                           not null comment '密码',
    userName     varchar(256)                           null comment '用户昵称',
    userAvatar   varchar(1024)                          null comment '用户头像',
    userRole     varchar(256) default 'user'            not null comment '用户角色：user/admin',
    userStatus   tinyint      default 0                 not null comment '用户状态：0-正常，1-禁用',
    lastLogin    datetime                               null comment '最后登录时间',
    createTime   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    updateTime   datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    isDelete     tinyint      default 0                 not null comment '是否删除',
    index idx_userAccount (userAccount)
) comment '用户' collate = utf8mb4_unicode_ci;

-- 图表表
create table if not exists chart
(
    id              bigint auto_increment comment 'id' primary key,
    goal            text                                 null comment '分析目标',
    `name`          varchar(128)                         null comment '图表名称',
    chartTableName  varchar(64)                          null comment '用户原始数据的表名',
    chartType       varchar(128)                         null comment '图表类型',
    genChart        text                                 null comment '生成的图表数据',
    genResult       text                                 null comment '生成的分析结论',
    status          varchar(128) default 'wait'          not null comment 'wait,running,succeed,failed',
    execMessage     text                                 null comment '执行信息',
    userId          bigint                               null comment '创建用户 id',
    createTime      datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    updateTime      datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    isDelete        tinyint      default 0               not null comment '是否删除'
) comment '图表信息表' collate = utf8mb4_unicode_ci;

-- 操作日志表
create table if not exists operation_log
(
    id           bigint auto_increment comment 'id' primary key,
    userId       bigint                               not null comment '用户id',
    userName     varchar(256)                         null comment '用户名',
    operation    varchar(256)                         not null comment '操作',
    method       varchar(128)                         not null comment '请求方法',
    params       text                                 null comment '请求参数',
    ip           varchar(64)                          null comment 'IP地址',
    status       tinyint                              not null comment '操作状态：0-成功，1-失败',
    createTime   datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
) comment '操作日志表' collate = utf8mb4_unicode_ci;

-- 角色表
create table if not exists role
(
    id           bigint auto_increment comment 'id' primary key,
    name         varchar(64)                          not null comment '角色名称',
    description  varchar(256)                         null comment '角色描述',
    createTime   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    updateTime   datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '角色表' collate = utf8mb4_unicode_ci;

-- 权限表
create table if not exists permission
(
    id           bigint auto_increment comment 'id' primary key,
    name         varchar(64)                          not null comment '权限名称',
    description  varchar(256)                         null comment '权限描述',
    type         varchar(64)                          not null comment '权限类型：view-查看，edit-编辑，delete-删除',
    createTime   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    updateTime   datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '权限表' collate = utf8mb4_unicode_ci;

-- 角色权限关联表
create table if not exists role_permission
(
    id           bigint auto_increment comment 'id' primary key,
    roleId       bigint                               not null comment '角色id',
    permissionId bigint                               not null comment '权限id',
    index idx_roleId (roleId),
    index idx_permissionId (permissionId)
) comment '角色权限关联表' collate = utf8mb4_unicode_ci;

-- 用户角色关联表
create table if not exists user_role
(
    id           bigint auto_increment comment 'id' primary key,
    userId       bigint                               not null comment '用户id',
    roleId       bigint                               not null comment '角色id',
    index idx_userId (userId),
    index idx_roleId (roleId)
) comment '用户角色关联表' collate = utf8mb4_unicode_ci;

-- 邀请码表
create table if not exists invite_code
(
    id           bigint auto_increment comment 'id' primary key,
    code         varchar(64)                          not null comment '邀请码',
    maxUses      int          default 1               not null comment '最大使用次数，0表示不限制',
    usedCount    int          default 0               not null comment '已使用次数',
    status       tinyint      default 0               not null comment '状态：0-有效，1-无效',
    expireTime   datetime                             null comment '过期时间，null表示永不过期',
    createTime   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    updateTime   datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    index idx_code (code)
) comment '邀请码表' collate = utf8mb4_unicode_ci;

-- 初始化管理员账号
INSERT INTO user (userAccount, userPassword, userName, userRole, userStatus, createTime, updateTime)
VALUES ('admin', MD5('admin123'), '系统管理员', 'admin', 0, NOW(), NOW());

-- 初始化角色
INSERT INTO role (name, description, createTime, updateTime)
VALUES ('admin', '管理员', NOW(), NOW()),
       ('operator', '运营人员', NOW(), NOW()),
       ('user', '普通用户', NOW(), NOW());

-- 初始化权限
INSERT INTO permission (name, description, type, createTime, updateTime)
VALUES ('用户查看', '查看用户列表和详情', 'view', NOW(), NOW()),
       ('用户编辑', '创建和编辑用户信息', 'edit', NOW(), NOW()),
       ('用户删除', '删除用户', 'delete', NOW(), NOW()),
       ('图表查看', '查看图表列表和详情', 'view', NOW(), NOW()),
       ('图表编辑', '创建和编辑图表', 'edit', NOW(), NOW()),
       ('图表删除', '删除图表', 'delete', NOW(), NOW()),
       ('日志查看', '查看操作日志', 'view', NOW(), NOW()),
       ('日志删除', '删除操作日志', 'delete', NOW(), NOW()),
       ('邀请码管理', '管理邀请码', 'edit', NOW(), NOW());

-- 为管理员角色分配所有权限
INSERT INTO role_permission (roleId, permissionId)
SELECT 1, id FROM permission;

-- 为运营角色分配部分权限
INSERT INTO role_permission (roleId, permissionId)
SELECT 2, id FROM permission WHERE name IN ('用户查看', '图表查看', '图表编辑', '日志查看', '邀请码管理');

-- 为管理员用户分配管理员角色
INSERT INTO user_role (userId, roleId)
VALUES (1, 1);
