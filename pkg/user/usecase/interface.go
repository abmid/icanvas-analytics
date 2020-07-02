package usecase

import "github.com/abmid/icanvas-analytics/pkg/user/entity"

type UserUseCase interface {
	Create(user *entity.User) error
	Find(email string) (*entity.User, error)
	All() ([]entity.User, error)
}
