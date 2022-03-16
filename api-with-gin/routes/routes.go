package routes

import (
	"api-with-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Gretting)
	r.POST("/students", controllers.Create)

	r.Run()
}
