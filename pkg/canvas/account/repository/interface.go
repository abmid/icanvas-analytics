package repository

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type AccountRepository interface {
	ListAccount(accountID uint32) ([]entity.Account, error)
}
