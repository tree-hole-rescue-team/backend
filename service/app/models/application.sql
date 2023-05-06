-- 有缓存 goctl model mysql ddl -src application.sql -dir . -c
-- 无缓存 goctl model mysql ddl -src application.sql -dir . 
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