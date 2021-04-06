package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/atranimal/go-echo-boilerplate/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Middlewares
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))	

	routes.RegisterRoutes(e)

	//Start Server
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("Shutting Down...")
		}
	}()
	//Wait For interrupt signal gracefully shutdown the server after timeout.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
