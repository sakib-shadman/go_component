package service

import (
	"errors"
	"go_user/domain"
	"go_user/models"
)

type UserService struct {
	repo domain.IUserRepo
}

func (service *UserService) CreateUser(user *models.User) error {
	if err := service.repo.CreateUser(user); err != nil {
		return errors.New("book was not created")
	}
	return nil

}

func UserServiceInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &UserService{
		repo: userRepo,
	}
}
