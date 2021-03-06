/*
 * File Created: Monday, 6th July 2020 2:48:10 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/canvas/account/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type accountUC struct {
	AccountRepo repository.AccountRepository
	Log         *logger.LoggerWrap
}

func NewUseCase(accountRepo repository.AccountRepository) *accountUC {

	logger := logger.New()

	return &accountUC{
		AccountRepo: accountRepo,
		Log:         logger,
	}
}

func (UC *accountUC) ListAccount(accountID uint32) (res []entity.Account, err error) {
	res, err = UC.AccountRepo.ListAccount(accountID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
