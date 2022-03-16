package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/students", ShowAllStudents)

	r.Run()
}

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Davi Silva",
	})
}
