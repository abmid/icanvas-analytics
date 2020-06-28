package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"
	"gotest.tools/assert"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func RealSetup() *sql.DB {
	parse, err := pgx.ParseURI("postgres://abdulhamid:@localhost:5432/canvas_analytics_go?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	db := stdlib.OpenDB(parse)
	return db
}

func TestCreateReal(t *testing.T) {
	repo := NewDiscussionPG(RealSetup())
	reportDiss := entity.ReportDiscussion{
		CourseReportID: 1,
		DiscussionID:   1,
		Title:          "Title Diss",
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	err := repo.Create(context.TODO(), &reportDiss)
	logrus.Error(err)
	t.Log(reportDiss)
	t.Fatalf("P")
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	reportDiss := entity.ReportDiscussion{
		CourseReportID: 1,
		DiscussionID:   1,
		Title:          "Title Diss",
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	mock.ExpectQuery("INSERT INTO "+DBTABLE).
		WithArgs(reportDiss.CourseReportID, reportDiss.DiscussionID, reportDiss.Title, AnyTime{}, AnyTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

	repo := NewDiscussionPG(db)
	err = repo.Create(context.TODO(), &reportDiss)
	assert.NilError(t, err, "Error Test Create")
	assert.Equal(t, uint32(2), reportDiss.ID)
}

func TestRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	reportDiss := []entity.ReportDiscussion{
		{
			ID:             1,
			CourseReportID: 1,
			DiscussionID:   1,
			Title:          "Title Diss",
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "course_report_id", "discussion_id", "title", "created_at", "updated_at"}).AddRow(
		reportDiss[0].ID,
		reportDiss[0].CourseReportID,
		reportDiss[0].DiscussionID,
		reportDiss[0].Title,
		reportDiss[0].CreatedAt.Time,
		reportDiss[0].UpdatedAt.Time,
	)
	mock.ExpectQuery("SELECT id, course_report_id, discussion_id, title, created_at, updated_at").WillReturnRows(rows)
	repo := NewDiscussionPG(db)
	results, err := repo.Read(context.Background())
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(reportDiss), len(results), "Not same len result")
	assert.Equal(t, reportDiss[0].ID, results[0].ID, "Not same result")
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	reportDiss := entity.ReportDiscussion{
		ID:             2,
		CourseReportID: 1,
		DiscussionID:   1,
		Title:          "Update Diss",
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	mock.ExpectQuery("UPDATE "+DBTABLE).WithArgs(
		reportDiss.CourseReportID, reportDiss.DiscussionID, reportDiss.Title, AnyTime{}, reportDiss.ID,
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
	repo := NewDiscussionPG(db)
	err = repo.Update(context.Background(), &reportDiss)
	assert.NilError(t, err, "Error update")
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln("Error")
	}
	defer db.Close()
	mock.ExpectExec("DELETE FROM report_discussions WHERE id = $1").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewDiscussionPG(db)
	err = repo.Delete(context.Background(), uint32(1))
	assert.NilError(t, err, "Error Delete")
}

func TestFindByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	filter := entity.ReportDiscussion{
		ID:           1,
		DiscussionID: 1,
	}
	discussions := []entity.ReportDiscussion{
		{
			ID:             1,
			CourseReportID: 1,
			Title:          "Title Assigment",
			DiscussionID:   1,
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	t.Run("find-exists", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "course_report_id", "discussion_id", "title", "created_at", "updated_at"}).AddRow(
			discussions[0].ID,
			discussions[0].CourseReportID,
			discussions[0].DiscussionID,
			discussions[0].Title,
			discussions[0].CreatedAt.Time,
			discussions[0].UpdatedAt.Time,
		)
		mock.ExpectQuery("SELECT").WithArgs(filter.ID, filter.DiscussionID).WillReturnRows(rows)
		repo := NewDiscussionPG(db)
		res, err := repo.FindFirstByFilter(context.Background(), filter)
		t.Log(err)
		t.Log(res)
		assert.NilError(t, err)
		assert.Equal(t, discussions[0].ID, res.ID)
	})
	t.Run("not-exists", func(t *testing.T) {
		nullValue := entity.ReportDiscussion{}
		rows := sqlmock.NewRows([]string{"id", "course_report_id", "discussion_id", "title", "created_at", "updated_at"}).AddRow(
			nullValue.ID,
			nullValue.CourseReportID,
			nullValue.DiscussionID,
			nullValue.Title,
			nullValue.CreatedAt.Time,
			nullValue.UpdatedAt.Time,
		)
		mock.ExpectQuery("SELECT").WithArgs(filter.ID, filter.DiscussionID).WillReturnRows(rows)
		repo := NewDiscussionPG(db)
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
	reportDiss := []entity.ReportDiscussion{
		{
			ID:             1,
			CourseReportID: 1,
			DiscussionID:   1,
			Title:          "Title Diss",
			CreatedAt:      sql.NullTime{Time: time.Now()},
			UpdatedAt:      sql.NullTime{Time: time.Now()},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "course_report_id", "discussion_id", "title", "created_at", "updated_at"}).AddRow(
		reportDiss[0].ID,
		reportDiss[0].CourseReportID,
		reportDiss[0].DiscussionID,
		reportDiss[0].Title,
		reportDiss[0].CreatedAt.Time,
		reportDiss[0].UpdatedAt.Time,
	)
	filter := entity.ReportDiscussion{
		CourseReportID: 1,
	}
	mock.ExpectQuery("SELECT id, course_report_id, discussion_id, title, created_at, updated_at").WithArgs(filter.CourseReportID).WillReturnRows(rows)
	repo := NewDiscussionPG(db)
	results, err := repo.FindFilter(context.Background(), filter)
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(reportDiss), len(results), "Not same len result")
	assert.Equal(t, reportDiss[0].ID, results[0].ID, "Not same result")
}
