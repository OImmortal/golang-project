package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositorio"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuarios models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuarios); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuarios.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repositorio.NovoRepositorioDeUsuarios(db)
	usuarioId, erro := rep.Criar(usuarios)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	usuarios.Id = usuarioId

	respostas.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar por todos os usuarios"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar por 1 usuário"))
}

func AutalizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
