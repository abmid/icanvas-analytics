package usecase

import (
	"github.com/abmid/icanvas-analytics/pkg/canvas/assigment/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type assigmentUseCase struct {
	AssigmentRepo repository.AssigmentRepository
}

func NewAssigmentUseCase(assigmentRepo repository.AssigmentRepository) *assigmentUseCase {
	return &assigmentUseCase{
		AssigmentRepo: assigmentRepo,
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
