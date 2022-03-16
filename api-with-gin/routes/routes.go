package routes

import (
	"api-with-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowAllStudents)

	r.Run()
}
