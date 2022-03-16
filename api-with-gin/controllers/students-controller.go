package controllers

import (
	"api-with-gin/database"
	"api-with-gin/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(200, students)
}

func Gretting(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"APY says:": "Hello " + strings.Title(strings.ToLower(name)) + ", are you ok ?",
	})
}

func Create(c *gin.Context) {
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
