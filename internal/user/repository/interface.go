/*
 * File Created: Thursday, 18th June 2020 4:58:48 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import "github.com/abmid/icanvas-analytics/internal/user/entity"

type UserRepository interface {
	Create(user *entity.User) error
}
