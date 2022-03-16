package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Davi Silva",
	})
}

func Gretting(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"APY says:": "Hello " + strings.Title(strings.ToLower(name)) + ", are you ok ?",
	})
}
