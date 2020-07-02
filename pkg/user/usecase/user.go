package usecase

import (
	"github.com/abmid/icanvas-analytics/pkg/user/entity"
	"github.com/abmid/icanvas-analytics/pkg/user/repository"
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

func (uc *userUseCase) Find(email string) (*entity.User, error) {
	res, err := uc.userRepo.Find(email)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (uc *userUseCase) All() ([]entity.User, error) {
	res, err := uc.userRepo.All()
	if err != nil {
		return nil, err
	}
	return res, err
}
