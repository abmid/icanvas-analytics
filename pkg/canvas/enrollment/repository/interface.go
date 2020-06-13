package repository

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type EnrollRepository interface {
	ListEnrollmentByCourseID(courseID uint32) (res []entity.Enrollment, err error)
}
