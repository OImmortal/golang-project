CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(50) not null,
    criadoEm timestamp default current_timestamp()
);

CREATE TABLE seguidores (
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
);

CREATE TABLE publicacoes(
    id int auto_increment PRIMARY KEY,
    titulo varchar(50) not NULL,
    conteudo text NOT NULL,

    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp()
);

USE mysql;

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'admin';

GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost';