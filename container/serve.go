package container

import (
	"go_user/config"
	"go_user/connection"
)

func Serve() {

	config.SetConfig()

	connection.GetDB()

	//routes.BookRoutes(e)

	//log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
