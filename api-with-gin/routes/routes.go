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

	r.DELETE("/students/:id", controllers.Delete)

	r.POST("/students", controllers.Create)
	r.PATCH("/students/:id", controllers.Edit)
	r.GET("/students/cpf/:cpf", controllers.GetByCPF)

	r.Run()
}
