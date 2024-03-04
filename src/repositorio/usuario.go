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
