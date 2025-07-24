package repository

import (
	"fiber-clean-arch-demo/internal/domain"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetAll() ([]domain.User, error) {
	var users []domain.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepo) Create(user *domain.User) error {
	return r.db.Create(user).Error
}