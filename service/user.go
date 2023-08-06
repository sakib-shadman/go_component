package service

import (
	"go_user/domain"
	"go_user/models"
)

type UserService struct {
	repo domain.IUserRepo
}

func (service *UserService) CreateUser(user models.User) {

}

func UserServiceInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &UserService{
		repo: userRepo,
	}
}
