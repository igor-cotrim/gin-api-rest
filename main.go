package main

import (
	"github.com/igor-cotrim/gin-api-rest/database"
	"github.com/igor-cotrim/gin-api-rest/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
