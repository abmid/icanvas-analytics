package usecase

import "github.com/abmid/icanvas-analytics/internal/user/entity"

type UserUseCase interface {
	Create(user *entity.User) error
}
