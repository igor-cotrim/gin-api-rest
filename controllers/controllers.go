package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igor-cotrim/gin-api-rest/database"
	"github.com/igor-cotrim/gin-api-rest/models"
)

func AllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.Students)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API diz:": "E ai " + name + ", tudo beleza?",
	})
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
