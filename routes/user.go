package routes

import (
	"go_user/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	users := e.Group("/users")
	users.POST("/create-user", controllers.CreateUser)
}
