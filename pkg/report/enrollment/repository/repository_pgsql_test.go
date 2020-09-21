package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"
	"testing"
	"time"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"gotest.tools/assert"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	reportEnroll := entity.ReportEnrollment{
		CourseReportID: 1,
		EnrollmentID:   1,
		UserID:         1,
		RoleID:         1,
		Role:           "TeacherEnrollment",
		CurrentScore:   1.2,
		CurrentGrade:   1.2,
		FinalScore:     1.3,
		FinalGrade:     1.3,
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	mock.ExpectQuery("INSERT INTO " + DBTABLE).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	repo := NewEnrollmentPG(db)
	err = repo.Create(context.TODO(), &reportEnroll)
	assert.NilError(t, err, "Error Test Create")
	assert.Equal(t, uint32(2), reportEnroll.ID)
}

func TestRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	reportEnroll := []entity.ReportEnrollment{
		{
			ID:             1,
			CourseReportID: 1,
			EnrollmentID:   1,
			UserID:         1,
			RoleID:         1,
			Role:           "TeacherEnrollment",
			CurrentScore:   1.2,
			CurrentGrade:   1.2,
			FinalScore:     1.3,
			FinalGrade:     1.3,
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "course_report_id", "enrollment_id", "user_id", "login_id", "full_name", "role_id", "role", "current_score", "current_grade", "final_score", "final_grade", "created_at", "updated_at"}).AddRow(
		reportEnroll[0].ID,
		reportEnroll[0].CourseReportID,
		reportEnroll[0].EnrollmentID,
		reportEnroll[0].UserID,
		reportEnroll[0].LoginID,
		reportEnroll[0].FullName,
		reportEnroll[0].RoleID,
		reportEnroll[0].Role,
		reportEnroll[0].CurrentScore,
		reportEnroll[0].CurrentGrade,
		reportEnroll[0].FinalScore,
		reportEnroll[0].FinalGrade,
		reportEnroll[0].CreatedAt.Time,
		reportEnroll[0].UpdatedAt.Time,
	)
	mock.ExpectQuery("SELECT id, course_report_id, enrollment_id, user_id, login_id, full_name, role_id, role, current_score, current_grade, final_score, final_grade, created_at, updated_at").WillReturnRows(rows)
	repo := NewEnrollmentPG(db)
	results, err := repo.Read(context.Background())
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(reportEnroll), len(results), "Not same len result")
	assert.Equal(t, reportEnroll[0].ID, results[0].ID, "Not same result")
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	assigment := entity.ReportEnrollment{
		ID:             2,
		CourseReportID: 1,
		EnrollmentID:   1,
		UserID:         1,
		RoleID:         1,
		Role:           "TeacherEnrollment",
		CurrentScore:   1.2,
		CurrentGrade:   1.2,
		FinalScore:     1.3,
		FinalGrade:     1.3,
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	mock.ExpectQuery("UPDATE " + DBTABLE).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
	repo := NewEnrollmentPG(db)
	err = repo.Update(context.Background(), &assigment)
	assert.NilError(t, err, "Error update")
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	mock.ExpectExec("DELETE FROM report_enrollments WHERE id = $1").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewEnrollmentPG(db)
	err = repo.Delete(context.Background(), uint32(1))
	assert.NilError(t, err, "Error Delete")
}

func TestFindByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	filter := entity.ReportEnrollment{
		ID:           1,
		EnrollmentID: 1,
	}
	reportEnroll := []entity.ReportEnrollment{
		{
			ID:             1,
			CourseReportID: 1,
			EnrollmentID:   1,
			UserID:         1,
			RoleID:         1,
			Role:           "TeacherEnrollment",
			CurrentScore:   1.2,
			CurrentGrade:   1.2,
			FinalScore:     1.3,
			FinalGrade:     1.3,
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	t.Run("find-exists", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "course_report_id", "enrollment_id", "user_id", "login_id", "full_name", "role_id", "role", "current_score", "current_grade", "final_score", "final_grade", "created_at", "updated_at"}).AddRow(
			reportEnroll[0].ID,
			reportEnroll[0].CourseReportID,
			reportEnroll[0].EnrollmentID,
			reportEnroll[0].UserID,
			reportEnroll[0].LoginID,
			reportEnroll[0].FullName,
			reportEnroll[0].RoleID,
			reportEnroll[0].Role,
			reportEnroll[0].CurrentScore,
			reportEnroll[0].CurrentGrade,
			reportEnroll[0].FinalScore,
			reportEnroll[0].FinalGrade,
			reportEnroll[0].CreatedAt.Time,
			reportEnroll[0].UpdatedAt.Time,
		)
		mock.ExpectQuery("SELECT").WithArgs(filter.ID, filter.EnrollmentID).WillReturnRows(rows)
		repo := NewEnrollmentPG(db)
		res, err := repo.FindFirstByFilter(context.Background(), filter)
		t.Log(err)
		t.Log(res)
		assert.NilError(t, err)
		assert.Equal(t, reportEnroll[0].ID, res.ID)
	})
	t.Run("not-exists", func(t *testing.T) {
		nullValue := entity.ReportEnrollment{}
		rows := sqlmock.NewRows([]string{"id", "course_report_id", "enrollment_id", "user_id", "login_id", "full_name", "role_id", "role", "current_score", "current_grade", "final_score", "final_grade", "created_at", "updated_at"}).AddRow(
			nullValue.ID,
			nullValue.CourseReportID,
			nullValue.EnrollmentID,
			nullValue.UserID,
			nullValue.LoginID,
			nullValue.FullName,
			nullValue.RoleID,
			nullValue.Role,
			nullValue.CurrentScore,
			nullValue.CurrentGrade,
			nullValue.FinalScore,
			nullValue.FinalGrade,
			nullValue.CreatedAt.Time,
			nullValue.UpdatedAt.Time,
		)
		mock.ExpectQuery("SELECT").WithArgs(filter.ID, filter.EnrollmentID).WillReturnRows(rows)
		repo := NewEnrollmentPG(db)
		res, err := repo.FindFirstByFilter(context.Background(), filter)
		t.Log(err)
		t.Log(res)
		assert.NilError(t, err)
		assert.Equal(t, res.ID, uint32(0))
	})
}

func TestFindFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	reportEnroll := []entity.ReportEnrollment{
		{
			ID:             1,
			CourseReportID: 1,
			EnrollmentID:   1,
			UserID:         1,
			RoleID:         1,
			Role:           "TeacherEnrollment",
			CurrentScore:   1.2,
			CurrentGrade:   1.2,
			FinalScore:     1.3,
			FinalGrade:     1.3,
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "course_report_id", "enrollment_id", "user_id", "login_id", "full_name", "role_id", "role", "current_score", "current_grade", "final_score", "final_grade", "created_at", "updated_at"}).AddRow(
		reportEnroll[0].ID,
		reportEnroll[0].CourseReportID,
		reportEnroll[0].EnrollmentID,
		reportEnroll[0].UserID,
		reportEnroll[0].LoginID,
		reportEnroll[0].FullName,
		reportEnroll[0].RoleID,
		reportEnroll[0].Role,
		reportEnroll[0].CurrentScore,
		reportEnroll[0].CurrentGrade,
		reportEnroll[0].FinalScore,
		reportEnroll[0].FinalGrade,
		reportEnroll[0].CreatedAt.Time,
		reportEnroll[0].UpdatedAt.Time,
	)
	filter := entity.ReportEnrollment{
		CourseReportID: 1,
	}
	mock.ExpectQuery("SELECT id, course_report_id, enrollment_id, user_id, login_id, full_name, role_id, role, current_score, current_grade, final_score, final_grade, created_at, updated_at").WithArgs(filter.CourseReportID).WillReturnRows(rows)
	repo := NewEnrollmentPG(db)
	results, err := repo.FindFilter(context.Background(), filter)
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(reportEnroll), len(results), "Not same len result")
	assert.Equal(t, reportEnroll[0].ID, results[0].ID, "Not same result")
}
