package repositorio

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db: db}
}

func (repo usuarios) Criar(usurio models.Usuario) (uint, error) {
	statment, erro := repo.db.Prepare("insert into usuarios (nome,  nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statment.Close()

	resultado, erro := statment.Exec(usurio.Nome, usurio.Nick, usurio.Email, usurio.Senha)
	if erro != nil {
		return 0, erro
	}

	idInserito, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint(idInserito), nil
}

func (repo usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repo.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?", nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usurios []models.Usuario

	for linhas.Next() {
		var user models.Usuario
		if erro = linhas.Scan(&user.Id, &user.Nome, &user.Nick, &user.Email, &user.CriadoEm); erro != nil {
			return nil, erro
		}

		usurios = append(usurios, user)
	}

	return usurios, nil
}

func (repo usuarios) BuscarPorId(ID uint64) (models.Usuario, error) {
	linhas, erro := repo.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}

	var usuario models.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, nil
		}
	}

	return usuario, nil
}

func (repo usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statment, erro := repo.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repo usuarios) Deletar(ID uint64) error {
	statment, erro := repo.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

func (repo usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linhas, erro := repo.db.Query(
		"select id, senha from usuarios where email = ?",
		email,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario models.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Senha,
		); erro != nil {
			return models.Usuario{}, nil
		}
	}

	return usuario, nil
}

func (repo usuarios) Seguir(usuarioId, seguidorId uint64) error {
	statment, erro := repo.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statment.Close()
	if _, erro = statment.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (repo usuarios) PararDeSeguir(usuarioId, seguidorId uint64) error {
	statment, erro := repo.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (repo usuarios) BuscarSeguidores(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := repo.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?
	`, usuarioId)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repo usuarios) BuscarSeguindo(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := repo.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?
	`, usuarioId)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repo usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linha, erro := repo.db.Query("select senha from usuarios where id = ?", usuarioId)
	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", nil
		}
	}

	return usuario.Senha, nil
}

func (repo usuarios) AtualizarSenha(usuarioId uint64, senhaComHash string) error {
	statment, erro := repo.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(senhaComHash, usuarioId); erro != nil {
		return erro
	}

	return nil
}
