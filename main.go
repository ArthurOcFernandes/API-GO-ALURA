package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia bem legal"},
		{Id: 2, Nome: "Nome2", Historia: "Historia bem legal 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor rest com Go")

	routes.HandleRequest()
}
