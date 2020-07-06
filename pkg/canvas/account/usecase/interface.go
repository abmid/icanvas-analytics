package usecase

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type AccountUseCase interface {
	ListAccount(accountID uint32) ([]entity.Account, error)
}
