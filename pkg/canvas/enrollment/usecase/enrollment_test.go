package usecase

import (
	"context"
	"testing"

	mock_enrollment "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/repository/mock"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestListEnrollmentByCourseID(t *testing.T) {
	EnrollmentGrade := entity.EnrollmentGrade{
		HtmlURL:      "html_url",
		CurrentGrade: 80.5,
		CurrentScore: 13,
		FinalScore:   12.2,
		FinalGrade:   11.2,
	}
	ListEnroll := []entity.Enrollment{
		{ID: 1, Grades: EnrollmentGrade},
	}
	ctrl, _ := gomock.WithContext(context.Background(), t)
	mockEnrollRepo := mock_enrollment.NewMockEnrollRepository(ctrl)
	mockEnrollRepo.EXPECT().ListEnrollmentByCourseID(uint32(1)).Return(ListEnroll, nil)
	enrollUseCase := NewEnrollUseCase(mockEnrollRepo)
	res, err := enrollUseCase.ListEnrollmentByCourseID(1)
	assert.NilError(t, err, "Error List Enroll")
	assert.Equal(t, len(res), len(ListEnroll))
}
