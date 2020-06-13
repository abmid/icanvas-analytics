package repository

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type AssigmentRepository interface {
	ListAssigmentByCourseID(CourseID uint32) (res []entity.Assigment, err error)
}
