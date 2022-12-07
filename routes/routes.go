package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-cotrim/gin-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.AllStudents)
	r.GET("/:name", controllers.Salutation)

	r.Run()
}
