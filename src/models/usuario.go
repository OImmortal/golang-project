package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	Id       uint      `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"Senha,omitempty"`
	CriadoEm time.Time `json:"dataDeCriacao,omitempty"`
}

func (user *Usuario) Preparar(etapa string) error {
	if erro := user.validar(etapa); erro != nil {
		return erro
	}

	if erro := user.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (user *Usuario) validar(etapa string) error {
	if user.Nome == "" {
		return errors.New("Nome inválido")
	}

	if user.Email == "" {
		return errors.New("E-mail inválido")
	}

	if user.Nick == "" {
		return errors.New("Nick inválido")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("E-mail inserido é inválido")
	}

	if user.Senha == "" && etapa == "cadastro" {
		return errors.New("Senha inválido")
	}

	return nil
}

func (user *Usuario) formatar(etapa string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(user.Senha)
		if erro != nil {
			return erro
		}

		user.Senha = string(senhaComHash)
	}
	return nil
}
