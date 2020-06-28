/*
 * File Created: Thursday, 18th June 2020 5:24:07 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */
package usecase

import (
	"github.com/abmid/icanvas-analytics/pkg/user/entity"
	user_uc "github.com/abmid/icanvas-analytics/pkg/user/usecase"
)

type registerUC struct {
	UserUC user_uc.UserUseCase
}

func New(UserUC user_uc.UserUseCase) *registerUC {
	return &registerUC{
		UserUC: UserUC,
	}
}

func (UC *registerUC) Register(user *entity.User) error {
	err := UC.UserUC.Create(user)
	if err != nil {
		return err
	}
	return nil
}
