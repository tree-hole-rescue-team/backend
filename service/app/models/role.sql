-- 有缓存 goctl model mysql ddl -src role.sql -dir . -c
-- 无缓存 goctl model mysql ddl -src role.sql -dir . 
DROP TABLE IF EXISTS `roles_change`;
CREATE TABLE `roles_change`(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_id`     INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '身份变动用户id',
    `username`    varchar(50) NOT NULL DEFAULT '' COMMENT '身份变动用户姓名',
    `operator_id` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '操作人id',
    `operator_name` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人姓名',
    `new_role`    INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '新身份',
    `new_rolename` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '新身份名称',
    `old_role`    INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '旧身份',
    `old_rolename` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '旧身份名称',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='身份变动表';