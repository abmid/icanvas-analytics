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
	assigment := entity.ReportAssigment{
		CourseReportID: 1,
		AssigmentID:    1,
		Name:           "Name Assigment",
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	mock.ExpectQuery("INSERT INTO "+DBTABLE).
		WithArgs(assigment.CourseReportID, assigment.AssigmentID, assigment.Name, AnyTime{}, AnyTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	repo := NewAssigmentPG(db)
	err = repo.Create(context.TODO(), &assigment)
	assert.NilError(t, err, "Error Test Create")
	assert.Equal(t, uint32(2), assigment.ID)
}

func TestRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	assigments := []entity.ReportAssigment{
		{
			ID:             1,
			CourseReportID: 1,
			AssigmentID:    1,
			Name:           "Name Assigment",
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "course_report_id", "assigment_id", "name", "created_at", "updated_at"}).AddRow(
		assigments[0].ID,
		assigments[0].CourseReportID,
		assigments[0].AssigmentID,
		assigments[0].Name,
		assigments[0].CreatedAt.Time,
		assigments[0].UpdatedAt.Time,
	)
	mock.ExpectQuery("SELECT id, course_report_id, assigment_id, name, created_at, updated_at").WillReturnRows(rows)
	repo := NewAssigmentPG(db)
	results, err := repo.Read(context.Background())
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(assigments), len(results), "Not same len result")
	assert.Equal(t, assigments[0].ID, results[0].ID, "Not same result")
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	assigment := entity.ReportAssigment{
		ID:             2,
		CourseReportID: 1,
		AssigmentID:    1,
		Name:           "Update Assigment",
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	mock.ExpectExec("UPDATE report_assigment").
		WillReturnResult(sqlmock.NewResult(2, 1))
	repo := NewAssigmentPG(db)
	err = repo.Update(context.Background(), &assigment)
	assert.NilError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	mock.ExpectExec("DELETE FROM report_assigments WHERE id = $1").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewAssigmentPG(db)
	err = repo.Delete(context.Background(), uint32(1))
	assert.NilError(t, err, "Error Delete")
}

func TestCreateOrUpdateByCourseReportID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	assigment := entity.ReportAssigment{
		ID:             1,
		CourseReportID: 1,
		AssigmentID:    1,
		Name:           "Name Assigment",
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	t.Run("assigment-exists", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
		mock.ExpectQuery("SELECT id").WithArgs(uint32(1)).WillReturnRows(rows)

		repo := NewAssigmentPG(db)
		err = repo.CreateOrUpdateByCourseReportID(context.Background(), &assigment)
		assert.NilError(t, err)
		assert.Equal(t, uint32(3), assigment.ID)
	})
	t.Run("assigment-not-exists", func(t *testing.T) {
		assigment.ID = 1
		rows := sqlmock.NewRows([]string{"id"}).AddRow(0)
		mock.ExpectQuery("SELECT id").WithArgs(uint32(1)).WillReturnRows(rows)
		mock.ExpectQuery("INSERT INTO "+DBTABLE).
			WithArgs(assigment.CourseReportID, assigment.AssigmentID, assigment.Name, AnyTime{}, AnyTime{}).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

		repo := NewAssigmentPG(db)
		err = repo.CreateOrUpdateByCourseReportID(context.Background(), &assigment)
		assert.NilError(t, err)
		assert.Equal(t, uint32(2), assigment.ID)
	})
}

func TestFindByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	filter := entity.ReportAssigment{
		ID:          1,
		AssigmentID: 1,
	}
	assigments := []entity.ReportAssigment{
		{
			ID:             1,
			CourseReportID: 1,
			Name:           "Name Assigment",
			AssigmentID:    1,
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	t.Run("find-exists", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "course_report_id", "assigment_id", "name", "created_at", "updated_at"}).AddRow(
			assigments[0].ID,
			assigments[0].CourseReportID,
			assigments[0].AssigmentID,
			assigments[0].Name,
			assigments[0].CreatedAt.Time,
			assigments[0].UpdatedAt.Time,
		)
		mock.ExpectQuery("SELECT").WithArgs(filter.ID, filter.AssigmentID).WillReturnRows(rows)
		repo := NewAssigmentPG(db)
		res, err := repo.FindFirstByFilter(context.Background(), filter)
		assert.NilError(t, err)
		assert.Equal(t, assigments[0].ID, res.ID)
	})
	t.Run("not-exists", func(t *testing.T) {
		nullValue := entity.ReportAssigment{}
		rows := sqlmock.NewRows([]string{"id", "course_report_id", "assigment_id", "name", "created_at", "updated_at"}).AddRow(
			nullValue.ID,
			nullValue.CourseReportID,
			nullValue.AssigmentID,
			nullValue.Name,
			nullValue.CreatedAt.Time,
			nullValue.UpdatedAt.Time,
		)
		mock.ExpectQuery("SELECT").WithArgs(filter.ID, filter.AssigmentID).WillReturnRows(rows)
		repo := NewAssigmentPG(db)
		res, err := repo.FindFirstByFilter(context.Background(), filter)
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
	assigments := []entity.ReportAssigment{
		{
			ID:             1,
			CourseReportID: 1,
			AssigmentID:    1,
			Name:           "Name Assigment",
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "course_report_id", "assigment_id", "name", "created_at", "updated_at"}).AddRow(
		assigments[0].ID,
		assigments[0].CourseReportID,
		assigments[0].AssigmentID,
		assigments[0].Name,
		assigments[0].CreatedAt.Time,
		assigments[0].UpdatedAt.Time,
	)
	mock.ExpectQuery("SELECT id, course_report_id, assigment_id, name, created_at, updated_at").WillReturnRows(rows)
	repo := NewAssigmentPG(db)
	filter := entity.ReportAssigment{}
	results, err := repo.FindFilter(context.Background(), filter)
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(assigments), len(results), "Not same len result")
	assert.Equal(t, assigments[0].ID, results[0].ID, "Not same result")
}
