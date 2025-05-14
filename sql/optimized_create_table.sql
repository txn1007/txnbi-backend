-- 创建库
CREATE DATABASE IF NOT EXISTS txnbi DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 切换库
USE txnbi;

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
    `id`           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `userAccount`  VARCHAR(64)                           NOT NULL COMMENT '账号',
    `userPassword` VARCHAR(128)                          NOT NULL COMMENT '密码',
    `userName`     VARCHAR(64)                           NULL COMMENT '用户昵称',
    `userAvatar`   VARCHAR(255)                          NULL COMMENT '用户头像',
    `userRole`     VARCHAR(32) DEFAULT 'user'            NOT NULL COMMENT '用户角色：user/admin/operator',
    `userStatus`   TINYINT     DEFAULT 0                 NOT NULL COMMENT '用户状态：0-正常，1-禁用',
    `lastLogin`    DATETIME                              NULL COMMENT '最后登录时间',
    `createTime`   DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    `updateTime`   DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `isDelete`     TINYINT     DEFAULT 0                 NOT NULL COMMENT '是否删除',
    UNIQUE KEY `uk_userAccount` (`userAccount`),
    INDEX `idx_userRole` (`userRole`),
    INDEX `idx_userStatus` (`userStatus`),
    INDEX `idx_createTime` (`createTime`)
) ENGINE=InnoDB COMMENT '用户表' COLLATE=utf8mb4_unicode_ci;

-- 图表表
CREATE TABLE IF NOT EXISTS `chart` (
    `id`              BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `goal`            TEXT                                 NULL COMMENT '分析目标',
    `name`            VARCHAR(128)                         NULL COMMENT '图表名称',
    `chartTableName`  VARCHAR(64)                          NULL COMMENT '用户原始数据的表名',
    `chartType`       VARCHAR(32)                          NULL COMMENT '图表类型',
    `genChart`        TEXT                                 NULL COMMENT '生成的图表数据',
    `genResult`       TEXT                                 NULL COMMENT '生成的分析结论',
    `status`          VARCHAR(16) DEFAULT 'wait'           NOT NULL COMMENT '状态：wait,running,succeed,failed',
    `execMessage`     TEXT                                 NULL COMMENT '执行信息',
    `userId`          BIGINT UNSIGNED                      NULL COMMENT '创建用户ID',
    `createTime`      DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    `updateTime`      DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `isDelete`        TINYINT     DEFAULT 0                NOT NULL COMMENT '是否删除',
    INDEX `idx_userId` (`userId`),
    INDEX `idx_status` (`status`),
    INDEX `idx_createTime` (`createTime`),
    CONSTRAINT `fk_chart_user` FOREIGN KEY (`userId`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB COMMENT '图表信息表' COLLATE=utf8mb4_unicode_ci;

-- 操作日志表
CREATE TABLE IF NOT EXISTS `operation_log` (
    `id`           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `userId`       BIGINT UNSIGNED                      NOT NULL COMMENT '用户ID',
    `userName`     VARCHAR(64)                          NULL COMMENT '用户名',
    `userAccount`  VARCHAR(64)                          NULL COMMENT '用户账号',
    `operation`    VARCHAR(128)                         NOT NULL COMMENT '操作',
    `method`       VARCHAR(16)                          NOT NULL COMMENT '请求方法',
    `path`         VARCHAR(128)                         NULL COMMENT '请求路径',
    `params`       TEXT                                 NULL COMMENT '请求参数',
    `ip`           VARCHAR(64)                          NULL COMMENT 'IP地址',
    `status`       TINYINT                              NOT NULL COMMENT '操作状态：0-成功，1-失败',
    `createTime`   DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    INDEX `idx_userId` (`userId`),
    INDEX `idx_createTime` (`createTime`),
    INDEX `idx_operation` (`operation`),
    INDEX `idx_status` (`status`),
    CONSTRAINT `fk_log_user` FOREIGN KEY (`userId`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB COMMENT '操作日志表' COLLATE=utf8mb4_unicode_ci;

-- 角色表
CREATE TABLE IF NOT EXISTS `role` (
    `id`           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `name`         VARCHAR(32)                          NOT NULL COMMENT '角色名称',
    `description`  VARCHAR(128)                         NULL COMMENT '角色描述',
    `createTime`   DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    `updateTime`   DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB COMMENT '角色表' COLLATE=utf8mb4_unicode_ci;

-- 权限表
CREATE TABLE IF NOT EXISTS `permission` (
    `id`           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `name`         VARCHAR(64)                          NOT NULL COMMENT '权限名称',
    `description`  VARCHAR(128)                         NULL COMMENT '权限描述',
    `type`         VARCHAR(16)                          NOT NULL COMMENT '权限类型：view-查看，edit-编辑，delete-删除',
    `createTime`   DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    `updateTime`   DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_name` (`name`),
    INDEX `idx_type` (`type`)
) ENGINE=InnoDB COMMENT '权限表' COLLATE=utf8mb4_unicode_ci;

-- 角色权限关联表
CREATE TABLE IF NOT EXISTS `role_permission` (
    `id`           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `roleId`       BIGINT UNSIGNED                      NOT NULL COMMENT '角色ID',
    `permissionId` BIGINT UNSIGNED                      NOT NULL COMMENT '权限ID',
    UNIQUE KEY `uk_role_permission` (`roleId`, `permissionId`),
    INDEX `idx_permissionId` (`permissionId`),
    CONSTRAINT `fk_rp_role` FOREIGN KEY (`roleId`) REFERENCES `role` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_rp_permission` FOREIGN KEY (`permissionId`) REFERENCES `permission` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB COMMENT '角色权限关联表' COLLATE=utf8mb4_unicode_ci;

-- 用户角色关联表
CREATE TABLE IF NOT EXISTS `user_role` (
    `id`           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `userId`       BIGINT UNSIGNED                      NOT NULL COMMENT '用户ID',
    `roleId`       BIGINT UNSIGNED                      NOT NULL COMMENT '角色ID',
    UNIQUE KEY `uk_user_role` (`userId`, `roleId`),
    INDEX `idx_roleId` (`roleId`),
    CONSTRAINT `fk_ur_user` FOREIGN KEY (`userId`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_ur_role` FOREIGN KEY (`roleId`) REFERENCES `role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB COMMENT '用户角色关联表' COLLATE=utf8mb4_unicode_ci;

-- 邀请码表
CREATE TABLE IF NOT EXISTS `invite_code` (
    `id`           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID' PRIMARY KEY,
    `code`         VARCHAR(32)                          NOT NULL COMMENT '邀请码',
    `maxUses`      INT          DEFAULT 1               NOT NULL COMMENT '最大使用次数，0表示不限制',
    `usedCount`    INT          DEFAULT 0               NOT NULL COMMENT '已使用次数',
    `status`       TINYINT      DEFAULT 0               NOT NULL COMMENT '状态：0-有效，1-无效',
    `expireTime`   DATETIME                             NULL COMMENT '过期时间，NULL表示永不过期',
    `createTime`   DATETIME     DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    `updateTime`   DATETIME     DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_code` (`code`),
    INDEX `idx_status` (`status`),
    INDEX `idx_expireTime` (`expireTime`)
) ENGINE=InnoDB COMMENT '邀请码表' COLLATE=utf8mb4_unicode_ci;

-- 初始化管理员账号
INSERT INTO `users` (`userAccount`, `userPassword`, `userName`, `userRole`, `userStatus`, `createTime`, `updateTime`)
VALUES ('admin123', MD5('admin123456'), '系统管理员', 'admin', 0, NOW(), NOW());

-- 初始化角色
INSERT INTO `role` (`name`, `description`, `createTime`, `updateTime`)
VALUES 
    ('admin', '管理员', NOW(), NOW()),
    ('operator', '运营人员', NOW(), NOW()),
    ('user', '普通用户', NOW(), NOW());

-- 初始化权限
INSERT INTO `permission` (`name`, `description`, `type`, `createTime`, `updateTime`)
VALUES 
    ('用户查看', '查看用户列表和详情', 'view', NOW(), NOW()),
    ('用户编辑', '创建和编辑用户信息', 'edit', NOW(), NOW()),
    ('用户删除', '删除用户', 'delete', NOW(), NOW()),
    ('图表查看', '查看图表列表和详情', 'view', NOW(), NOW()),
    ('图表编辑', '创建和编辑图表', 'edit', NOW(), NOW()),
    ('图表删除', '删除图表', 'delete', NOW(), NOW()),
    ('日志查看', '查看操作日志', 'view', NOW(), NOW()),
    ('日志删除', '删除操作日志', 'delete', NOW(), NOW()),
    ('邀请码管理', '管理邀请码', 'edit', NOW(), NOW()),
    ('角色管理', '管理角色和权限', 'edit', NOW(), NOW());

-- 为管理员角色分配所有权限
INSERT INTO `role_permission` (`roleId`, `permissionId`)
SELECT 1, id FROM `permission`;

-- 为运营角色分配部分权限
INSERT INTO `role_permission` (`roleId`, `permissionId`)
SELECT 2, id FROM `permission` WHERE `name` IN ('用户查看', '图表查看', '图表编辑', '日志查看', '邀请码管理');

-- 为管理员用户分配管理员角色
INSERT INTO `user_role` (`userId`, `roleId`)
VALUES (1, 1);

-- 添加系统初始邀请码
INSERT INTO `invite_code` (`code`, `maxUses`, `status`, `createTime`, `updateTime`)
VALUES ('TXNBI2023', 10, 0, NOW(), NOW());