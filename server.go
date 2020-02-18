package main

import (
	"context"
	"github.com/hellofresh/hfcp-rack-backend/app"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()

	app.RegisterRoutes(e)

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