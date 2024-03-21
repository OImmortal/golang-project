package repositorio

import (
	"api/src/models"
	"database/sql"
)

type publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *publicacoes {
	return &publicacoes{db: db}
}

func (repo publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statment, erro := repo.db.Prepare("insert into publicacoes (titulo, conteudo, user_id) value (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statment.Close()

	resultado, erro := statment.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoId, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoId), erro
}
