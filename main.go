package main

import (
	"github.com/igor-cotrim/gin-api-rest/database"
	"github.com/igor-cotrim/gin-api-rest/models"
	"github.com/igor-cotrim/gin-api-rest/routes"
)

func main() {
	database.ConnectToDatabase()
	models.Students = []models.Student{
		{
			Name: "Igor Cotrim",
			CPF:  "000.000.000-40",
			RG:   "4532123",
		},
		{
			Name: "Bianca Cotrim",
			CPF:  "123.123.123-40",
			RG:   "4532123",
		},
	}

	routes.HandleRequests()
}
