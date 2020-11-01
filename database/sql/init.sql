---- create auth database ----
DROP DATABASE IF EXISTS auth;
CREATE DATABASE auth;

DROP TABLE IF EXISTS `auth`.`user`;
CREATE TABLE `auth`.`user`
(
 `id`       VARCHAR(256),
 `password` VARCHAR(256) NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

---- create user ----
CREATE USER IF NOT EXISTS 'auth_user'@'%' IDENTIFIED BY 'password';
GRANT ALL ON auth.* TO auth_user;
