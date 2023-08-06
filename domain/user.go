package domain

import "go_user/models"

type IUserRepo interface {
	CreateUser(user models.User)
}

type IUserService interface {
	CreateUser(user models.User)
}
