package routes

import (
	"app/pkg/app"
	"app/pkg/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	app *app.App,
) {

	userHandler := handler.NewUserHandler(app)

	{
		api := router.Group("/api")
		api.GET("/get-user", userHandler.GetUser)
		api.POST("/add-user", userHandler.AddUser)
	}

}
