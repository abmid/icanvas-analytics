/*
 * File Created: Thursday, 18th June 2020 4:58:48 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import "github.com/abmid/icanvas-analytics/pkg/user/entity"

type UserRepository interface {
	Create(user *entity.User) error
	Find(email string) (*entity.User, error)
	All() ([]entity.User, error)
}
