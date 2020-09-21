/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/canvas/discussion/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type discussionUseCase struct {
	DiscussRepo repository.DiscussionRepository
	Log         *logger.LoggerWrap
}

func NewDiscussUseCase(discussRepo repository.DiscussionRepository) *discussionUseCase {
	logger := logger.New()

	return &discussionUseCase{
		DiscussRepo: discussRepo,
		Log:         logger,
	}
}

func (DUC *discussionUseCase) ListDiscussionByCourseID(courseID uint32) (res []entity.Discussion, err error) {
	discussions, err := DUC.DiscussRepo.ListDiscussionByCourseID(courseID)
	if err != nil {
		DUC.Log.Error(err)
		return nil, err
	}
	res = discussions
	return res, nil
}
