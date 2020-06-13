package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gotest.tools/assert"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func RealSetup() *sql.DB {
	parse, err := pgx.ParseURI("postgres://abdulhamid:@localhost:5432/canvas_analytics_dev?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	db := stdlib.OpenDB(parse)
	return db
}

func TestFindBestCourseByFilterReal(t *testing.T) {
	ctx := context.TODO()
	repo := NewRepositoryPG(RealSetup())
	filter := entity.FilterAnalytics{
		OrderBy:          "desc",
		AnalyticsTeacher: true,
		Limit:            9,
	}
	res, err := repo.FindBestCourseByFilter(ctx, filter)
	t.Log(err)
	for _, each := range res {
		t.Log(each)
	}
	t.Fatalf("")
}

func TestT(t *testing.T) {
	rsult := (float32(1) / float32(1)) * 100
	t.Log(rsult)
	t.Logf("P")
	t.Fatalf("")
}

func TestFindBestCourseByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	ctx := context.TODO()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	t.Run("course", func(t *testing.T) {
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
		repo := NewRepositoryPG(db)
		filter := entity.FilterAnalytics{
			AnalyticsTeacher: false,
		}
		res, err := repo.FindBestCourseByFilter(ctx, filter)
		assert.NilError(t, err)
		assert.Equal(t, len(res), 1)
		averagaGrading := (float32(exceptedResult.FinishGradingCount) / float32(exceptedResult.StudentCount)) * 100
		assert.Equal(t, res[0].AverageGrading, averagaGrading)
		assert.Equal(t, res[0].CourseName, exceptedResult.CourseName)
	})
	t.Run("teacher", func(t *testing.T) {
		t.Run("course", func(t *testing.T) {
			exceptedTeacher := canvas.User{
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
			repo := NewRepositoryPG(db)
			filter := entity.FilterAnalytics{
				AnalyticsTeacher: true,
			}
			res, err := repo.FindBestCourseByFilter(ctx, filter)
			assert.NilError(t, err)
			assert.Equal(t, len(res), 1)
			averagaGrading := (float32(exceptedResult.FinishGradingCount) / float32(exceptedResult.StudentCount)) * 100
			assert.Equal(t, res[0].AverageGrading, averagaGrading)
			assert.Equal(t, res[0].CourseName, exceptedResult.CourseName)
			assert.Equal(t, res[0].Teacher.LoginID, exceptedTeacher.LoginID)
		})
	})
}
