package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", func(c echo.Context) error{
		return c.String(http.StatusOK, "Healthy")
	})
}