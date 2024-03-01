package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	Id       uint      `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"Senha,omitempty"`
	CriadoEm time.Time `json:"dataDeCriacao,omitempty"`
}

func (user *Usuario) Preparar() error {
	if erro := user.validar(); erro != nil {
		return erro
	}

	user.formatar()
	return nil
}

func (user *Usuario) validar() error {
	if user.Nome == "" {
		return errors.New("Nome inválido")
	}

	if user.Email == "" {
		return errors.New("E-mail inválido")
	}

	if user.Nick == "" {
		return errors.New("Nick inválido")
	}

	if user.Senha == "" {
		return errors.New("Senha inválido")
	}

	return nil
}

func (user *Usuario) formatar() {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.Senha = strings.TrimSpace(user.Senha)
}
