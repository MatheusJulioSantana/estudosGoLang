package main

import (
	"Loja/routes"
	"net/http"
)

// toda requisição vai para / e o index vai atender
func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
