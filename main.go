package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "1",
		"name": "Igor Cotrim",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", AllStudents)
	r.Run()
}
