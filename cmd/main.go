package main

import (
	"app/pkg/config"
	"app/pkg/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/pkg/app"

	"github.com/gin-gonic/gin"
)

func main() {

	config := config.LoadConfig()
	// ------------------------------------------------
	// Initialize application dependencies
	// ------------------------------------------------

	app, err := app.NewApp(config)

	if err != nil {
		log.Fatal(err)
	}

	defer app.Db.Close()

	// ------------------------------------------------
	// Create gin router
	// ------------------------------------------------
	router := gin.New()

	// middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// ------------------------------------------------
	// Register application routes
	// ------------------------------------------------

	routes.RegisterRoutes(router, app)

	// ------------------------------------------------
	// Custom http server with read and write timeouts
	// ------------------------------------------------

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Server.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("server running on port %s", config.Server.Port)

		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// ------------------------------------------------
	// Graceful shutdown
	// ------------------------------------------------

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("shutting down server...")

	// shutdown must happen in 10 seconds
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server exited cleanly")
}
