package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"
	mock_repository "github.com/abmid/icanvas-analytics/pkg/report/user/repository/mock"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCreateOrUpdateByCourseReportID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	mockRepo := mock_repository.NewMockReportResultRepository(ctrl)
	mockRepo.EXPECT().CreateOrUpdateByCourseReportID(ctx, gomock.Any()).DoAndReturn(
		func(ctx context.Context, m *entity.ReportUser) error {
			m.ID = 2
			m.CreatedAt = sql.NullTime{Time: time.Now()}

			return nil
		})
	uc := NewReportUserUseCase(mockRepo)
	rUser := entity.ReportUser{
		ReportCourseID: 1,
	}
	err := uc.CreateOrUpdateByCourseReportID(ctx, &rUser)
	assert.NilError(t, err)
	assert.Equal(t, uint32(2), rUser.ID)
}
