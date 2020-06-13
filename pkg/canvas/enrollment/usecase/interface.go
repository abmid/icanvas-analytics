package usecase

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type EnrollmentUseCase interface {
	ListEnrollmentByCourseID(courseID uint32) (res []entity.Enrollment, err error)
}