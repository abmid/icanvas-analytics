package usecase

import (
	"context"
	"testing"

	mock_assigment "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/repository/mock"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestListAssigmentByCourseID(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	assigmentRepo := mock_assigment.NewMockAssigmentRepository(ctrl)
	ListAssigment := []entity.Assigment{
		{ID: 1, CourseID: 1, Name: "Title Assigment"},
	}
	assigmentRepo.EXPECT().ListAssigmentByCourseID(uint32(1)).Return(ListAssigment, nil)
	assigmentUC := NewAssigmentUseCase(assigmentRepo)
	res, err := assigmentUC.ListAssigmentByCourseID(uint32(1))
	assert.NilError(t, err, "Error list assigment")
	assert.Equal(t, len(res), len(ListAssigment))
}
