package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/respostas"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuarios, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	response, erro := http.Post("http://10.1.1.33:5000/usuarios", "application/json", bytes.NewBuffer(usuarios))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
