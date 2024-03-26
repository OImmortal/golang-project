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

func (repo publicacoes) BuscarPublicacao(publicacaoId uint64) (models.Publicacao, error) {
	linha, erro := repo.db.Query(`
		select p.*, u.nick from publicacoes p 
		inner join usuarios u
		on u.id = p.user_id where p.id = ?
	`, publicacaoId)
	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repo publicacoes) BuscarPublicacoes(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := repo.db.Query("SELECT DISTINCT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id = p.user_id INNER JOIN seguidores s ON p.user_id = s.usuario_id WHERE u.id = ? OR s.seguidor_id = ? order by 1 desc", usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {

		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repo publicacoes) Atualizar(publicacaoId uint64, publicacao models.Publicacao) error {
	statment, erro := repo.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return nil
	}

	return nil
}

func (repo publicacoes) Deletar(publicacaoId uint64) error {
	statment, erro := repo.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro := statment.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repo publicacoes) BuscarPublicacaoPorUsuario(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := repo.db.Query("SELECT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id = p.user_id WHERE p.user_id = ?", usuarioId)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repo publicacoes) Curtir(publicacaoId uint64) error {
	statment, erro := repo.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro := statment.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repo publicacoes) Descurtir(publicacoesId uint64) error {
	statment, erro := repo.db.Prepare(`
		update publicacoes set curtidas =
		CASE WHEN curtidas > 0 THEN curtidas - 1
		ELSE 0 END
		where id = ?
	`)
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro := statment.Exec(publicacoesId); erro != nil {
		return erro
	}

	return nil
}
