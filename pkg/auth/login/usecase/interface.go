package usecase

import "github.com/abmid/icanvas-analytics/pkg/user/entity"

type LoginUseCase interface {
	Login(email, password string) (*entity.User, int, error)
}
