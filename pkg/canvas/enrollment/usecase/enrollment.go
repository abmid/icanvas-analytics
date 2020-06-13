package usecase

import (
	"github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type enrollUseCase struct {
	EnrollRepo repository.EnrollRepository
}

func NewEnrollUseCase(enrollRepo repository.EnrollRepository) *enrollUseCase {
	return &enrollUseCase{
		EnrollRepo: enrollRepo,
	}
}

func (EUC *enrollUseCase) ListEnrollmentByCourseID(courseID uint32) (res []entity.Enrollment, err error) {
	res, err = EUC.EnrollRepo.ListEnrollmentByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	return res, err
}
