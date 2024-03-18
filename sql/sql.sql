CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(50) not null,
    criadoEm timestamp default current_timestamp()
)

USE mysql;

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'admin';

GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost';