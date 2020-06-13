package usecase

import (
	mock_discuss "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/repository/mock"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestListDiscussionByCourseID(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	mockRepoDiscuss := mock_discuss.NewMockDiscussionRepository(ctrl)
	ListDiscuss := []entity.Discussion{
		{ID: 1, Published: true, Title: "Title Discussion"},
	}
	mockRepoDiscuss.EXPECT().ListDiscussionByCourseID(uint32(1)).Return(ListDiscuss, nil)
	discussUseCase := NewDiscussUseCase(mockRepoDiscuss)
	res, err := discussUseCase.ListDiscussionByCourseID(1)
	assert.NilError(t, err, "Error List Discussion")
	assert.Equal(t, len(res), len(ListDiscuss))
}
