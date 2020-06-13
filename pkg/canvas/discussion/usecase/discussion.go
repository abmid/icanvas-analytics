package usecase

import (
	"github.com/abmid/icanvas-analytics/pkg/canvas/discussion/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type discussionUseCase struct {
	DiscussRepo repository.DiscussionRepository
}

func NewDiscussUseCase(discussRepo repository.DiscussionRepository) *discussionUseCase {
	return &discussionUseCase{
		DiscussRepo: discussRepo,
	}
}

func (DUC *discussionUseCase) ListDiscussionByCourseID(courseID uint32) (res []entity.Discussion, err error) {
	discussions, err := DUC.DiscussRepo.ListDiscussionByCourseID(courseID)
	if err != nil {
		return nil, err
	}
	res = discussions
	return res, nil
}
