INSERT INTO usuarios (nome, nick, email, senha) VALUES 
("usuario 1", "usuario_1", "usuario1@email.com", "$2a$10$EDYMpUQX9Fv3tTOEBnyzpeSNGJFw7RyqhdYKB0Umuac0iolj30.5C"),
("usuario 2", "usuario_2", "usuario2@email.com", "$2a$10$EDYMpUQX9Fv3tTOEBnyzpeSNGJFw7RyqhdYKB0Umuac0iolj30.5C"),
("usuario 3", "usuario_3", "usuario3@email.com", "$2a$10$EDYMpUQX9Fv3tTOEBnyzpeSNGJFw7RyqhdYKB0Umuac0iolj30.5C");

INSERT INTO seguidores (usuario_id, seguidor_id) VALUES
(1, 2),
(3, 1),
(1, 3);