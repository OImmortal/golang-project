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
