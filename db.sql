-- 如果报错Error 1292 (22007): Incorrect datetime value: '0000-00-00' for column 'delete_time' at row 1
-- 去删除sql_mode中的NO_ZERO_DATE和NO_ZERO_IN_DATE
-- 查看当前sql_mode select @@sql_mode;
-- 查看全局sql_mode select @@global.sql_mode;
-- 修改当前sql_mode set sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
-- 修改全局sql_mode set global sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
-- 注意，docker上的mysql在配置后不要restart的，否则配置全部重置
-- 下面是docker上的mysql的时区的问题，默认是system，即北京时间-8小时
-- 查看数据库时区 select @@global.time_zone;
-- 设置时区为中国 set global time_zone='Asia/Shanghai';
DROP DATABASE IF EXISTS `managementsystem`;
CREATE DATABASE `managementsystem`;
USE `managementsystem`;

-- 用户
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户编号',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `open_id`     varchar(200) NOT NULL DEFAULT '' COMMENT '微信用户id',
    `mobile`      char(11) NOT NULL DEFAULT '' COMMENT '电话',
    `username`    varchar(50) NOT NULL DEFAULT '' COMMENT '用户名称',
    `password`    varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码',
    `sex`         int NOT NULL DEFAULT '0' COMMENT '性别 0:男 1:女',
    `email`       varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
    `role`        int NOT NULL DEFAULT '0' COMMENT '身份代号 0-申请队员 1-岗前培训 2-见习队员 3-正式队员 4-督导老师 30-普通队员 31-核心队员 32-区域负责人 33-组委会成员 34-组委会主任',
    `rolename`    varchar(255) NOT NULL DEFAULT '申请队员' COMMENT '身份名称',
    `avatar`      varchar(255) NOT NULL DEFAULT '0' COMMENT '头像',
    `address`     varchar(255) NOT NULL DEFAULT '0' COMMENT '地址',
    `birthday`    varchar(20) NOT NULL DEFAULT '0000-00-00' COMMENT '生日 xxxx-xx-xx',
    PRIMARY KEY(`id`),
    UNIQUE KEY `idx_mobile` (`mobile`) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- 用户授权
DROP TABLE IF EXISTS `users_auth`;
CREATE TABLE `users_auth`(
    `id`          int unsigned NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_id`     int unsigned NOT NULL DEFAULT '0',
    `auth_key`    varchar(64) NOT NULL DEFAULT '' COMMENT '平台唯一id',
    `auth_type`   varchar(12) NOT NULL DEFAULT '' COMMENT '平台类型',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_type_key` (`auth_type`,`auth_key`) USING BTREE, -- 复合索引，索引的存储类型为BTREE
    UNIQUE KEY `idx_userId_key` (`user_id`,`auth_type`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户授权表';

-- 权限变动
DROP TABLE IF EXISTS `roles_change`;
CREATE TABLE `roles_change`(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_id`     INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '身份变动用户',
    `username`    varchar(50) NOT NULL DEFAULT '' COMMENT '身份变动用户姓名',
    `operator_id` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '操作人',
    `operator_name` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人姓名',
    `new_role`    INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '新身份',
    `new_rolename` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '新身份名称',
    `old_role`    INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '旧身份',
    `old_rolename` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '旧身份名称',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='身份变动表';

-- 救援信息
DROP TABLE IF EXISTS `rescue_info`;
CREATE TABLE `rescue_info`(
    `id`                    INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '救援信息编号',
    `weibo_account`         VARCHAR(200) DEFAULT '' NOT NULL COMMENT '微博账号',
    `area`                  VARCHAR(500) DEFAULT '' COMMENT '所在城市',
    `sex`                   VARCHAR(200) DEFAULT '' COMMENT '性别',
    `age`                   VARCHAR(10)  DEFAULT '' COMMENT '年龄',
    `release_time`          VARCHAR(30) DEFAULT '' COMMENT '发布时间',
    `main_demand`           VARCHAR(30) DEFAULT '' COMMENT '主要诉求',
    `risk_level`            VARCHAR(10) DEFAULT '' COMMENT '风险等级',
    `status`                INT NOT NULL DEFAULT '0' COMMENT '救援状态 0-待救援 1-救援中 2-已救援',
    `rescue_teacher_id`     INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '救助老师id',
    `rescue_teacher_name`   VARCHAR(20) DEFAULT '' COMMENT '救助老师姓名',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='救援信息表';

-- 申请表
DROP TABLE IF EXISTS `application_forms`;
CREATE TABLE `application_forms`(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '申请表编号',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_id`     int NOT NULL DEFAULT '0' COMMENT '申请人id',
    `mobile`      char(11) NOT NULL DEFAULT '' COMMENT '电话',
    `username`    varchar(50) NOT NULL DEFAULT '' COMMENT '用户名称',
    `sex`         int NOT NULL DEFAULT '0' COMMENT '性别 0-男 1-女',
    `address`     varchar(255) NOT NULL DEFAULT '0' COMMENT '地址',
    `birthday`    varchar(20) NOT NULL DEFAULT '0000-00-00' COMMENT '生日 xxxx-xx-xx',
    `email`       varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
    `status`      int NOT NULL DEFAULT '0' COMMENT '申请表状态 0-待审批 1-已通过 2-不合格',
    `operator_id` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '操作人id',
    `operator_name` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人姓名',
    PRIMARY KEY(`id`),
    UNIQUE KEY`idx_userId` (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='申请表';