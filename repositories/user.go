package repositories

import (
	"go_user/domain"
	"go_user/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func (repo *userRepo) CreateUser(user models.User) {

}

func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: d,
	}
}
