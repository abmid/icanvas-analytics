package usecase

import (
	"sync"

	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type CourseUseCase interface {
	Courses(accountId, page uint32) (res []entity.Course, err error)
	ListUserInCourse(courseID uint32, enrollmentRole string) (res []entity.User, err error)
	AllCourse(accountId, pool uint32) (res []entity.Course, err error)
	GoAllCourse(accountID uint32, ch chan<- []entity.Course, wg *sync.WaitGroup)
}
