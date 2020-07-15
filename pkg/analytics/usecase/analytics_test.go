package usecase

import (
	"context"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/pagination"
	"github.com/abmid/icanvas-analytics/pkg/analytics/entity"
	mock_repoAnalytics "github.com/abmid/icanvas-analytics/pkg/analytics/repository/mock"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestFindBestCourseByFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.TODO()
	mockRepo := mock_repoAnalytics.NewMockAnalyticsRepository(ctrl)
	t.Run("course", func(t *testing.T) {
		filter := entity.FilterAnalytics{
			AnalyticsTeacher: false,
		}
		exceptedResult := []entity.AnalyticsCourse{{
			ID:                 1,
			AssigmentCount:     1,
			CourseID:           1,
			CourseName:         "Name",
			DiscussionCount:    1,
			FinalScore:         1,
			FinishGradingCount: 5,
			StudentCount:       10,
			AverageGrading:     50,
		}}
		exceptedPag := pagination.Pagination{
			CurrentPage: 1,
		}
		mockRepo.EXPECT().FindBestCourseByFilter(ctx, filter).Return(exceptedResult, exceptedPag, nil)
		UC := NewAnalyticsUseCase(mockRepo)
		res, pag, err := UC.FindBestCourseByFilter(ctx, filter)
		assert.NilError(t, err)
		assert.Equal(t, len(res), len(exceptedResult))
		assert.Equal(t, pag.CurrentPage, exceptedPag.CurrentPage)
	})
	t.Run("teacher", func(t *testing.T) {
		filter := entity.FilterAnalytics{
			AnalyticsTeacher: true,
		}
		exceptedTeacher := canvas.User{
			ID:          1,
			Name:        "Anony",
			Enrollments: "TeacherEnrollment",
			LoginID:     "name@domain.com",
		}
		exceptedResult := []entity.AnalyticsCourse{{
			ID:                 1,
			AssigmentCount:     1,
			CourseID:           1,
			CourseName:         "Name",
			DiscussionCount:    1,
			FinalScore:         1,
			FinishGradingCount: 5,
			StudentCount:       10,
			AverageGrading:     50,
			Teacher:            &exceptedTeacher,
		}}
		exceptedPag := pagination.Pagination{
			CurrentPage: 1,
		}
		mockRepo.EXPECT().FindBestCourseByFilter(ctx, filter).Return(exceptedResult, exceptedPag, nil)
		UC := NewAnalyticsUseCase(mockRepo)
		res, pag, err := UC.FindBestCourseByFilter(ctx, filter)
		assert.NilError(t, err)
		assert.Equal(t, len(res), len(exceptedResult))
		assert.Equal(t, res[0].Teacher.LoginID, exceptedResult[0].Teacher.LoginID)
		assert.Equal(t, pag.CurrentPage, exceptedPag.CurrentPage)
	})
}
