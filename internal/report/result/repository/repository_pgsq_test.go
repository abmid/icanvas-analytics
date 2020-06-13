package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/report/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"gotest.tools/assert"
)

func RealSetup() *sql.DB {
	parse, err := pgx.ParseURI("postgres://abdulhamid:@localhost:5432/canvas_analytics_dev?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	db := stdlib.OpenDB(parse)
	return db
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ctx := context.Background()
	result := entity.ReportResult{
		AssigmentCount: 1,
	}
	mock.ExpectQuery("INSERT INTO " + DBTABLE).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	repo := NewResultPG(db)
	err = repo.Create(ctx, &result)
	assert.NilError(t, err)
	assert.Equal(t, uint32(2), result.ID)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	mock.ExpectQuery("UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(3))
	repo := NewResultPG(db)
	result := entity.ReportResult{
		AssigmentCount: 1,
	}
	err = repo.Update(context.TODO(), &result)
	assert.NilError(t, err)
	assert.Equal(t, uint32(3), result.ID)
}

func TestCreateOrUpdateByCourseReportID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ctx := context.Background()
	t.Run("exist", func(t *testing.T) {
		mock.ExpectQuery("SELECT id").WithArgs(uint32(1)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(3))
		mock.ExpectQuery("UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(3))

		repo := NewResultPG(db)
		result := entity.ReportResult{
			ReportCourseID: 1,
		}
		err := repo.CreateOrUpdateByCourseReportID(ctx, &result)
		assert.NilError(t, err)
		assert.Equal(t, uint32(3), result.ID)
	})
	t.Run("not-exist", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WithArgs(uint32(1)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(0))
		mock.ExpectQuery("INSERT INTO " + DBTABLE).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

		repo := NewResultPG(db)
		result := entity.ReportResult{
			ReportCourseID: 1,
		}
		err := repo.CreateOrUpdateByCourseReportID(ctx, &result)
		assert.NilError(t, err)
		assert.Equal(t, uint32(2), result.ID)
	})

}
