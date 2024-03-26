package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTempletes()
	r := router.Gerar()

	fmt.Println("Rodando WebApp http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
