package repository

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type CourseRepository interface {
	Courses(accountId, page uint32) (res []entity.Course, err error)
	ListUserInCourse(courseID uint32, enrollmentRole string) (res []entity.User, err error)
}
