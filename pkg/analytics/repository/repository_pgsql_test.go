package repository

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/abmid/icanvas-analytics/internal/pagination"
	paginate_mock "github.com/abmid/icanvas-analytics/internal/pagination/mock"
	"github.com/abmid/icanvas-analytics/pkg/analytics/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"github.com/golang/mock/gomock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gotest.tools/assert"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestFindBestCourseByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	ctx := context.TODO()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	t.Run("course", func(t *testing.T) {
		// Create Mock Query
		exceptedResult := entity.AnalyticsCourse{
			ID:                 1,
			AccountID:          1,
			AssigmentCount:     1,
			CourseID:           1,
			CourseName:         "Name",
			DiscussionCount:    1,
			FinalScore:         1,
			FinishGradingCount: 5,
			StudentCount:       10,
		}

		rows := sqlmock.NewRows([]string{"id", "account_id", "course_id", "course_name", "assigment_count", "discussion_count", "student_count", "finish_grading_count", "final_score"}).AddRow(
			exceptedResult.ID,
			exceptedResult.AccountID,
			exceptedResult.CourseID,
			exceptedResult.CourseName,
			exceptedResult.AssigmentCount,
			exceptedResult.DiscussionCount,
			exceptedResult.StudentCount,
			exceptedResult.FinishGradingCount,
			exceptedResult.FinalScore,
		)

		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		// Paginate Mock
		exceptedPaginate := pagination.Pagination{
			Total: 1,
		}
		ctrl := gomock.NewController(t)
		paginationMock := paginate_mock.NewMockPaginationInterface(ctrl)
		paginationMock.EXPECT().BuildPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(exceptedPaginate, nil)

		// Init Repo
		repo := NewRepositoryPG(db, paginationMock)

		filter := entity.FilterAnalytics{
			AnalyticsTeacher: false,
		}

		res, pag, err := repo.FindBestCourseByFilter(ctx, filter)
		assert.NilError(t, err)
		assert.Equal(t, len(res), 1)
		assert.Equal(t, uint32(1), pag.Total)
		averagaGrading := (float32(exceptedResult.FinishGradingCount) / float32(exceptedResult.StudentCount)) * 100
		assert.Equal(t, res[0].AverageGrading, averagaGrading, "error course average grade")
		assert.Equal(t, res[0].CourseName, exceptedResult.CourseName, "error course courseName")
	})
	t.Run("teacher", func(t *testing.T) {
		t.Run("course", func(t *testing.T) {
			exceptedTeacher := &canvas.User{
				ID:          1,
				Name:        "Anony",
				LoginID:     "login@domain.com",
				Enrollments: "TeacherEnrollment",
			}
			exceptedResult := entity.AnalyticsCourse{
				ID:                 1,
				AccountID:          1,
				AssigmentCount:     1,
				CourseID:           1,
				CourseName:         "Name",
				DiscussionCount:    1,
				FinalScore:         1,
				FinishGradingCount: 5,
				StudentCount:       10,
				Teacher:            exceptedTeacher,
			}
			rows := sqlmock.NewRows([]string{"id", "account_id", "course_id", "course_name", "assigment_count", "discussion_count", "student_count", "finish_grading_count", "final_score", "full_name", "login_id", "user_id", "role"}).AddRow(
				exceptedResult.ID,
				exceptedResult.AccountID,
				exceptedResult.CourseID,
				exceptedResult.CourseName,
				exceptedResult.AssigmentCount,
				exceptedResult.DiscussionCount,
				exceptedResult.StudentCount,
				exceptedResult.FinishGradingCount,
				exceptedResult.FinalScore,
				exceptedTeacher.Name,
				exceptedTeacher.LoginID,
				exceptedTeacher.ID,
				exceptedTeacher.Enrollments,
			)
			mock.ExpectQuery("SELECT").WillReturnRows(rows)
			// Paginate Mock
			exceptedPaginate := pagination.Pagination{
				Total: 1,
			}
			ctrl := gomock.NewController(t)
			paginationMock := paginate_mock.NewMockPaginationInterface(ctrl)
			paginationMock.EXPECT().BuildPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(exceptedPaginate, nil)

			// Init Repo
			repo := NewRepositoryPG(db, paginationMock)

			filter := entity.FilterAnalytics{
				AnalyticsTeacher: true,
			}

			res, pag, err := repo.FindBestCourseByFilter(ctx, filter)
			assert.NilError(t, err)
			assert.Equal(t, len(res), 1)
			assert.Equal(t, uint32(1), pag.Total)
			averagaGrading := (float32(exceptedResult.FinishGradingCount) / float32(exceptedResult.StudentCount)) * 100
			assert.Equal(t, res[0].AverageGrading, averagaGrading, "Error teacher Average Grade")
			assert.Equal(t, res[0].CourseName, exceptedResult.CourseName, "Error teacher CourseName")
			assert.Equal(t, res[0].Teacher.LoginID, exceptedTeacher.LoginID, "Error teacher LoginID")
		})
	})
}
