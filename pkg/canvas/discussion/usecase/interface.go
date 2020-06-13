package usecase

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type DiscussionUseCase interface {
	ListDiscussionByCourseID(courseID uint32) (res []entity.Discussion, err error)
}
