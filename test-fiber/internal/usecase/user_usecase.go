package usecase

import (
	"fiber-clean-arch-demo/internal/domain"

	"github.com/google/uuid"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) GetAll() ([]domain.User, error) {
	return u.repo.GetAll()
}

func (u *userUsecase) Create(request *domain.CreateUserRequest) (*domain.User, error) {
	// 生成 uuid7
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:    id,
		Name:  request.Name,
		Email: request.Email,
	}

	if err := u.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}