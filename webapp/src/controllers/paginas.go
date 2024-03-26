package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplete(w, "login.html", nil)
}

func CarregarPaginaCadUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplete(w, "cadastro.html", nil)
}
