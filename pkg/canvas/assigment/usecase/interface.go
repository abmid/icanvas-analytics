package usecase

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type AssigmentUseCase interface {
	ListAssigmentByCourseID(CourseID uint32) (res []entity.Assigment, err error)
}
