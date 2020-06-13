/*
 * File Created: Saturday, 6th June 2020 11:37:01 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Your Company
 */
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

func TestT(t *testing.T) {
	repo := NewUserPG(RealSetup())
	ReportUser := entity.ReportUser{
		ReportCourseID: 12,
		FullName:       "a",
		LoginID:        "asdasdas",
		UserID:         12312,
	}
	err := repo.CreateOrUpdateByCourseReportID(context.Background(), &ReportUser)
	t.Log(err)
	t.Log(ReportUser)
	t.Fatalf("")
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	mock.ExpectQuery("INSERT INTO " + DBTABLE).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	repo := NewUserPG(db)
	rUser := entity.ReportUser{
		ReportCourseID: 1,
		FullName:       "ABD",
	}
	err = repo.Create(context.Background(), &rUser)
	assert.NilError(t, err)
	assert.Equal(t, uint32(2), rUser.ID)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	mock.ExpectQuery("UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	repo := NewUserPG(db)
	rUser := entity.ReportUser{
		ID:             1,
		ReportCourseID: 1,
		FullName:       "ABD",
	}
	err = repo.Update(context.Background(), &rUser)
	assert.NilError(t, err)
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

		repo := NewUserPG(db)
		result := entity.ReportUser{
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

		repo := NewUserPG(db)
		result := entity.ReportUser{
			ReportCourseID: 1,
		}
		err := repo.CreateOrUpdateByCourseReportID(ctx, &result)
		assert.NilError(t, err)
		assert.Equal(t, uint32(2), result.ID)
	})

}
