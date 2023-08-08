package container

import (
	"fmt"
	"go_user/config"
	"go_user/connection"
	"go_user/controllers"
	"go_user/repositories"
	"go_user/routes"
	"go_user/service"
	"log"

	"github.com/labstack/echo/v4"
)

func Serve() {

	e := echo.New()

	config.SetConfig()

	db := connection.GetDB()

	userRepo := repositories.UserDBInstance(db)

	userService := service.UserServiceInstance(userRepo)

	controllers.SetUserService(userService)

	routes.UserRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
