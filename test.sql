-- 存储一些需要在本地修改的sql语句

-- 修改权限
UPDATE `users` SET `role` = 34, `rolename`= "组委会主任" WHERE `id` = 1;

-- 连接服务器数据库
mysql -h zxjing.cn -P 8036 -u root -p