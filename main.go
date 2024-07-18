package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia1"},
		{Nome: "Nome 2", Historia: "Historia2"},
	}



	fmt.Println("Iniciando o servidor RESt com go")
	routes.HandleRequest()
}
