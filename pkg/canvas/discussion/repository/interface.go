package repository

import "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type DiscussionRepository interface {
	ListDiscussionByCourseID(courseID uint32) (res []entity.Discussion, err error)
}
