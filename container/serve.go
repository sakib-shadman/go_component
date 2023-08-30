package container

import (
	"bytes"
	"fmt"
	"go_user/config"
	"go_user/connection"
	"go_user/controllers"
	"go_user/repositories"
	"go_user/routes"
	"go_user/service"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
)

func Serve() {

	e := echo.New()

	e.Use(logBodyMiddleware)

	config.SetConfig()

	db := connection.GetDB()

	userRepo := repositories.UserDBInstance(db)

	userService := service.UserServiceInstance(userRepo)

	controllers.SetUserService(userService)

	routes.UserRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}

func logBodyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Content-Type") != echo.MIMEApplicationJSON {
			return next(c)
		}

		data, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		c.Logger().Error(string(data))
		c.Request().Body = ioutil.NopCloser(bytes.NewReader(data))

		return next(c)
	}
}
