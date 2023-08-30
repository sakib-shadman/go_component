package controllers

import (
	"fmt"
	"go_user/domain"
	"go_user/models"
	"go_user/types"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var UserService domain.IUserService

func SetUserService(uService domain.IUserService) {
	UserService = uService
}

func CreateUser(e echo.Context) error {

	reqUser := &types.CreateUserRequest{}
	error := e.Bind(reqUser)
	if error != nil {
		return e.JSON(http.StatusBadRequest, error.Error())
	}

	user := &models.User{
		FirstName:   reqUser.FirstName,
		LastName:    reqUser.LastName,
		Email:       reqUser.Email,
		PhoneNumber: reqUser.PhoneNumber,
	}

	if err := UserService.CreateUser(user); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	if error != nil {
		return e.JSON(http.StatusInternalServerError, error.Error())
	}
	return e.JSON(http.StatusCreated, "User was created successfully")

}

func LogRequest(e echo.Context) error {
	fmt.Println("The current time is:", time.Now())
	return e.JSON(http.StatusCreated, "User was created successfully")
}
