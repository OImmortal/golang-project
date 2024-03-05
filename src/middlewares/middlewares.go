package middlewares

import (
	"api/src/auth"
	"api/src/respostas"
	"log"
	"net/http"
)

func Logger(proxiamFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proxiamFuncao(w, r)
	}
}

func Autenticar(proxiamFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proxiamFuncao(w, r)
	}
}
