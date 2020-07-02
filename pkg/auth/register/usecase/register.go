/*
 * File Created: Thursday, 18th June 2020 5:24:07 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */
package usecase

import (
	"errors"

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
	users, err := UC.UserUC.All()
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errors.New("Failed to create user admin ! user admin already exists in database, please create user from dashboard")
	}

	err = UC.UserUC.Create(user)
	if err != nil {
		return err
	}
	return nil
}

// RegisterCheck is function to check register can do or not
// This is for help web client to redirect page welcome
func (UC *registerUC) RegisterCheck() (bool, error) {
	users, err := UC.UserUC.All()
	if err != nil {
		return false, err
	}
	if len(users) > 0 {
		return false, nil
	}
	return true, nil
}
