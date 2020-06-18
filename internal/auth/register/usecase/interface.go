/*
 * File Created: Thursday, 18th June 2020 4:26:43 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */
package usecase

import "github.com/abmid/icanvas-analytics/internal/user/entity"

type RegisterUseCase interface {
	Register(user *entity.User) error
}
