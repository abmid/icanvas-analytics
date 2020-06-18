/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type AssigmentUseCase interface {
	ListAssigmentByCourseID(CourseID uint32) (res []entity.Assigment, err error)
}
