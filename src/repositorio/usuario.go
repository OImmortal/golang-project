package repositorio

import (
	"api/src/models"
	"database/sql"
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
