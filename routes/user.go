package routes

import "github.com/labstack/echo/v4"

func userRoutes(e *echo.Echo) {
	users := e.Group("/users")
}
