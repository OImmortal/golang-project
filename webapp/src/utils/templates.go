package utils

import (
	"net/http"
	"text/template"
)

var templetes *template.Template

func CarregarTempletes() {
	templetes = template.Must(template.ParseGlob("views/*.html"))
}

func ExecutarTemplete(w http.ResponseWriter, template string, dados interface{}) {
	templetes.ExecuteTemplate(w, template, dados)
}
