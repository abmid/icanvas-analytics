/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type EnrollRepository interface {
	ListEnrollmentByCourseID(courseID uint32) (res []entity.Enrollment, err error)
}
