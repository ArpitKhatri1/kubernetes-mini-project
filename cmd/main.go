package main

import (
	"app/pkg/config"
	"app/pkg/database"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App Struct {
	DB *pgxpool.Pool
}

func main() {

	router := gin.Default()

	config := config.LoadConfig()

	if err := database.RunMigrations(config.DB); err != nil {
		panic("failed to migrate db")
	}

	//setup database Connection Pool
	db, err := database.NewPostgresPool(config.DB)

	if err != nil {
		panic("Error starting up new DB Pool")
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		panic("error in ping")
	}

	queries := database.New()

	app := &App{
		DB: pool
	}

	router.GET("/fetch", app.getSomething)
	router.POST("/posting", app.postSomething)

	router.Run(":8080")
}

func (app *App)addSomething(c *gin.Context) {

}	
func postSomething(c *gin.Context) {

}
