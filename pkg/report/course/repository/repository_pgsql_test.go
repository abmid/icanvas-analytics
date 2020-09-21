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
	reportCourse := entity.ReportCourse{
		AccountID:  1,
		CourseID:   1,
		CourseName: "Name Course",
	}
	mock.ExpectQuery("INSERT INTO "+DBTABLE).
		WithArgs(reportCourse.CourseID, reportCourse.CourseName, reportCourse.AccountID, AnyTime{}, AnyTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	repo := NewCoursePG(db)
	err = repo.Create(context.TODO(), &reportCourse)
	assert.NilError(t, err, "Error Test Create")
	assert.Equal(t, uint32(2), reportCourse.ID)
}

func TestRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	courseReports := []entity.ReportCourse{
		{
			ID:         1,
			AccountID:  1,
			CourseID:   1,
			CourseName: "Name Course",
			CreatedAt:  sql.NullTime{Time: time.Now()},
			UpdatedAt:  sql.NullTime{Time: time.Now()},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "course_id", "course_name", "account_id", "created_at", "updated_at", "deleted_at"}).AddRow(
		courseReports[0].ID,
		courseReports[0].CourseID,
		courseReports[0].CourseName,
		courseReports[0].AccountID,
		courseReports[0].CreatedAt.Time,
		courseReports[0].UpdatedAt.Time,
		courseReports[0].DeletedAt,
	)
	mock.ExpectQuery("SELECT id, course_id, course_name, account_id, created_at, updated_at").WillReturnRows(rows)
	repo := NewCoursePG(db)
	results, err := repo.Read(context.Background())
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(courseReports), len(results))
	assert.Equal(t, courseReports[0].ID, results[0].ID)
}

func TestFindFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	courseReports := []entity.ReportCourse{
		{
			ID:         1,
			AccountID:  1,
			CourseID:   1,
			CourseName: "Name Course",
			CreatedAt:  sql.NullTime{Time: time.Now()},
			UpdatedAt:  sql.NullTime{Time: time.Now()},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "course_id", "course_name", "account_id", "created_at", "updated_at", "deleted_at"}).AddRow(
		courseReports[0].ID,
		courseReports[0].CourseID,
		courseReports[0].CourseName,
		courseReports[0].AccountID,
		courseReports[0].CreatedAt.Time,
		courseReports[0].UpdatedAt.Time,
		courseReports[0].DeletedAt,
	)
	mock.ExpectQuery("SELECT id, course_id, course_name, account_id, created_at, updated_at").WillReturnRows(rows)
	repo := NewCoursePG(db)
	filter := entity.ReportCourse{}
	results, err := repo.FindFilter(context.Background(), filter)
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(courseReports), len(results))
	assert.Equal(t, courseReports[0].ID, results[0].ID)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	reportCourse := entity.ReportCourse{
		ID:         1,
		AccountID:  1,
		CourseID:   1,
		CourseName: "Name Course",
		CreatedAt:  sql.NullTime{Time: time.Now()},
		UpdatedAt:  sql.NullTime{Time: time.Now()},
	}
	mock.ExpectQuery("UPDATE "+DBTABLE).WithArgs(
		reportCourse.CourseID, reportCourse.CourseName, reportCourse.AccountID, AnyTime{}, reportCourse.ID,
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
	repo := NewCoursePG(db)
	err = repo.Update(context.Background(), &reportCourse)
	assert.NilError(t, err, "Error update")
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	mock.ExpectExec("DELETE FROM report_courses WHERE id = $1").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewCoursePG(db)
	err = repo.Delete(context.Background(), uint32(1))
	assert.NilError(t, err, "Error Delete")
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("error")
	}
	defer db.Close()
	courseReports := entity.ReportCourse{
		ID:         1,
		AccountID:  1,
		CourseID:   1,
		CourseName: "Name Course",
		CreatedAt:  sql.NullTime{Time: time.Now()},
		UpdatedAt:  sql.NullTime{Time: time.Now()},
	}
	rows := sqlmock.NewRows([]string{"id", "course_id", "course_name", "account_id", "created_at", "updated_at", "deleted_at"}).AddRow(
		courseReports.ID,
		courseReports.CourseID,
		courseReports.CourseName,
		courseReports.AccountID,
		courseReports.CreatedAt.Time,
		courseReports.UpdatedAt.Time,
		courseReports.DeletedAt.Time,
	)
	mock.ExpectQuery("SELECT").WithArgs(uint32(1)).WillReturnRows(rows, nil)
	uc := NewCoursePG(db)
	res, err := uc.FindByID(context.Background(), uint32(1))
	assert.NilError(t, err)
	assert.Equal(t, uint32(1), res.ID)
}

func TestFindByCourseIDDateNow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("error")
	}
	defer db.Close()
	courseReports := entity.ReportCourse{
		ID:         1,
		AccountID:  1,
		CourseID:   1,
		CourseName: "Name Course",
		CreatedAt:  sql.NullTime{Time: time.Now()},
		UpdatedAt:  sql.NullTime{Time: time.Now()},
	}
	rows := sqlmock.NewRows([]string{"id", "course_id", "course_name", "account_id", "created_at", "updated_at", "deleted_at"}).AddRow(
		courseReports.ID,
		courseReports.CourseID,
		courseReports.CourseName,
		courseReports.AccountID,
		courseReports.CreatedAt.Time,
		courseReports.UpdatedAt.Time,
		courseReports.DeletedAt.Time,
	)
	mock.ExpectQuery("SELECT").WithArgs(uint32(1)).WillReturnRows(rows, nil)
	uc := NewCoursePG(db)
	res, err := uc.FindByCourseIDDateNow(context.Background(), uint32(1))
	assert.NilError(t, err)
	assert.Equal(t, uint32(1), res.ID)
}
