/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/canvas/assigment/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type assigmentUseCase struct {
	AssigmentRepo repository.AssigmentRepository
	Log           *logger.LoggerWrap
}

func NewAssigmentUseCase(assigmentRepo repository.AssigmentRepository) *assigmentUseCase {

	logger := logger.New()

	return &assigmentUseCase{
		AssigmentRepo: assigmentRepo,
		Log:           logger,
	}
}

func (AUC *assigmentUseCase) ListAssigmentByCourseID(courseID uint32) (res []entity.Assigment, err error) {
	assigments, err := AUC.AssigmentRepo.ListAssigmentByCourseID(courseID)
	if err != nil {
		return nil, err
	}
	res = assigments
	return res, nil
}
