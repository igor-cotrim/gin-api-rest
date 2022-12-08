package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-cotrim/gin-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/:name", controllers.Salutation)
	r.GET("/alunos", controllers.AllStudents)
	r.GET("/alunos/:id", controllers.SearchStudentForId)
	r.POST("/alunos", controllers.CreateNewStudent)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.PATCH("/alunos/:id", controllers.EditStudent)

	r.Run()
}
