package main

import (
	"app/pkg/config"
	"app/pkg/database"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config := config.LoadConfig()

	db, err := database.NewPostgresPool(config.DB)

	if err != nil {
		panic("Error starting up new DB Pool")
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		panic("error in ping")
	}

	router.GET("/fetch", addSomething)
	router.POST("/posting", postSomething)

	router.Run(":8080")
}

func addSomething(c *gin.Context) {

}
func postSomething(c *gin.Context) {

}
