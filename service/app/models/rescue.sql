-- 无缓存 goctl model mysql ddl -src rescue.sql -dir . 
-- 有缓存 goctl model mysql ddl -src rescue.sql -dir . -c
/* DROP TABLE IF EXISTS `rescue_info`;
CREATE TABLE `rescue_info`(
    `id`             INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '救援信息编号',
    `uid`            VARCHAR(20) DEFAULT '' NOT NULL COMMENT '微博用户的uid',
    `nickname`       VARCHAR(50) NOT NULL DEFAULT '' COMMENT '昵称',
    `certified`      VARCHAR(50) DEFAULT '' COMMENT '认证信息',
    `sex`            VARCHAR(200) DEFAULT '' COMMENT '性别',
    `relationship`   VARCHAR(20) DEFAULT '' COMMENT '感情状况',
    `area`           VARCHAR(500) DEFAULT '' COMMENT '地区',
    `birthday`       VARCHAR(50) DEFAULT '' COMMENT '生日',
    `education_info` VARCHAR(300) DEFAULT '' COMMENT '学习经历',
    `work_info`      VARCHAR(300) DEFAULT '' COMMENT '工作经历',
    `description`    VARCHAR(2500) DEFAULT '' COMMENT '简介',
    `weibo_address`  VARCHAR(300) DEFAULT '' COMMENT '微博地址',
    `release_time`   VARCHAR(30) DEFAULT '' COMMENT '发布时间',
    `risk_level`     VARCHAR(10) DEFAULT '' COMMENT '风险等级',
    `main_demand`    VARCHAR(30) DEFAULT '' COMMENT '主要诉求',
    `status`         INT NOT NULL DEFAULT '0' COMMENT '救援状态 0-待救援 1-救援中 2-已救援',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='救援信息表'; */

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