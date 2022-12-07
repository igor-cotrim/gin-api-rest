package controllers

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

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API diz:": "E ai " + name + ", tudo beleza?",
	})
}
