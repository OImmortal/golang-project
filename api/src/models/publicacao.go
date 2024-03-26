package models

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorid,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func (pub *Publicacao) Preparar() error {
	if erro := pub.validar(); erro != nil {
		return erro
	}

	pub.formatar()
	return nil
}

func (pub *Publicacao) validar() error {
	if pub.Titulo == "" {
		return errors.New("O é obrigatório informar um título")
	}

	if pub.Conteudo == "" {
		return errors.New("O é obrigatório informar um conteudo")
	}

	return nil
}

func (pub *Publicacao) formatar() {
	pub.Titulo = strings.TrimSpace(pub.Titulo)
	pub.Conteudo = strings.TrimSpace(pub.Conteudo)
}
