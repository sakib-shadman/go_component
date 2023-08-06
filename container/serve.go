package container

import (
	"go_user/config"
	"go_user/connection"
	"go_user/controllers"
	"go_user/repositories"
	"go_user/service"
)

func Serve() {

	config.SetConfig()

	db := connection.GetDB()

	userRepo := repositories.UserDBInstance(db)

	userService := service.UserServiceInstance(userRepo)

	controllers.SetUserService(userService)

	//routes.UserR(e)

	//log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
