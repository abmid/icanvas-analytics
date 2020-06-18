package usecase

import (
	"github.com/abmid/icanvas-analytics/internal/user/entity"
	"github.com/abmid/icanvas-analytics/internal/user/repository"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func New(userRepo repository.UserRepository) *userUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) Create(user *entity.User) error {
	err := uc.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
