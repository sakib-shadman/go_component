package controllers

import (
	"go_user/domain"

	"github.com/labstack/echo/v4"
)

var UserService domain.IUserService

func SetUserService(uService domain.IUserService) {
	UserService = uService
}

func CreateUser(e echo.Context) {

}
