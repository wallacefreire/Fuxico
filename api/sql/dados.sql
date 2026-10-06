insert into usuarios(nome, nick, email, senha) 
values
("Usuário 1", "usuario_1", "usuario_1@gmail.com", "$2a$10$yDhbKgqDF.0oFfgJeUC7wui5hSrMH2wFkkRMaNYmecyO5mA3ULTiO"),
("Usuário 2", "usuario_2", "usuario_2@gmail.com", "$2a$10$yDhbKgqDF.0oFfgJeUC7wui5hSrMH2wFkkRMaNYmecyO5mA3ULTiO"),
("Usuário 3", "usuario_3", "usuario_3@gmail.com", "$2a$10$yDhbKgqDF.0oFfgJeUC7wui5hSrMH2wFkkRMaNYmecyO5mA3ULTiO");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação 1", "Conteúdo da publicação 1", 1),
("Publicação 2", "Conteúdo da publicação 2", 2),
("Publicação 3", "Conteúdo da publicação 3", 3);