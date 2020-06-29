/*
 * File Created: Saturday, 27th June 2020 11:34:35 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import "github.com/abmid/icanvas-analytics/pkg/user/entity"

type LoginUseCase interface {
	Login(email, password string) (*entity.User, int, error)
}
