package repositories

import (
	"go_user/domain"
	"go_user/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// CreateUser implements domain.IUserRepo
func (repo *userRepo) CreateUser(user *models.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: d,
	}
}
