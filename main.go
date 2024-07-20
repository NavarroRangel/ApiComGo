package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia2"},
	}
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor RESt com go")
	routes.HandleRequest()
}
