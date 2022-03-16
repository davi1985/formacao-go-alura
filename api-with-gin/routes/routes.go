package routes

import (
	"api-with-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.Index)
	r.GET("/:name", controllers.Gretting)
	r.GET("/students/:id", controllers.FindById)

	r.POST("/students", controllers.Create)

	r.Run()
}
