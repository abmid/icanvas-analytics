package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	repository_real "github.com/abmid/icanvas-analytics/internal/report/course/repository"
	mock_assigment "github.com/abmid/icanvas-analytics/internal/report/course/repository/mock"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportCourse := mock_assigment.NewMockCourseRepositoryPG(ctrl)
	reportCourse := report.ReportCourse{
		CourseID: 1,
	}
	mockReportCourse.EXPECT().Create(context.Background(), &reportCourse).DoAndReturn(
		func(ctx context.Context, m *report.ReportCourse) error {
			m.ID = 2
			return nil
		})

	useCase := NewReportCourseUseCase(mockReportCourse)
	err := useCase.Create(context.Background(), &reportCourse)
	assert.NilError(t, err, "Error Create")
	assert.Equal(t, uint32(2), reportCourse.ID)
}

func TestRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportCourse := mock_assigment.NewMockCourseRepositoryPG(ctrl)
	reportCourse := []report.ReportCourse{
		{ID: 1, CourseName: "One"},
		{ID: 2, CourseName: "Two"},
	}
	mockReportCourse.EXPECT().Read(context.TODO()).Return(reportCourse, nil)
	useCase := NewReportCourseUseCase(mockReportCourse)
	result, err := useCase.Read(context.TODO())
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(result), len(reportCourse))
}

func TestFindFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportCourse := mock_assigment.NewMockCourseRepositoryPG(ctrl)
	filter := report.ReportCourse{}
	reportCourse := []report.ReportCourse{
		{ID: 1, CourseName: "One"},
		{ID: 2, CourseName: "Two"},
	}
	mockReportCourse.EXPECT().FindFilter(context.TODO(), filter).Return(reportCourse, nil)
	useCase := NewReportCourseUseCase(mockReportCourse)
	result, err := useCase.FindFilter(context.TODO(), filter)
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(result), len(reportCourse))
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportAss := mock_assigment.NewMockCourseRepositoryPG(ctrl)
	timeUpdate := time.Now()
	reportCourse := report.ReportCourse{
		ID:         1,
		CourseName: "Name",
		AccountID:  1,
	}
	mockReportAss.EXPECT().Update(context.TODO(), &reportCourse).DoAndReturn(func(ctx context.Context, m *report.ReportCourse) error {
		m.ID = reportCourse.ID
		m.CourseName = reportCourse.CourseName
		m.AccountID = reportCourse.AccountID
		m.UpdatedAt = sql.NullTime{Time: timeUpdate}
		return nil
	})
	useCase := NewReportCourseUseCase(mockReportAss)
	err := useCase.Update(context.TODO(), &reportCourse)
	assert.NilError(t, err)
	assert.Equal(t, timeUpdate, reportCourse.UpdatedAt.Time)
}

func removeSlice(report []report.ReportCourse, index int) []report.ReportCourse {
	return append(report[:index], report[index+1:]...)
}
func TestRemoveSlice(t *testing.T) {
	reportCourse := []report.ReportCourse{
		{ID: 1, CourseName: "One"},
		{ID: 2, CourseName: "Two"},
	}
	for index, ass := range reportCourse {
		if ass.ID == 1 {
			reportCourse = removeSlice(reportCourse, index)
		}
	}
	assert.Equal(t, len(reportCourse), 1)
	assert.Equal(t, reportCourse[0].ID, uint32(2))
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockreportCourse := mock_assigment.NewMockCourseRepositoryPG(ctrl)
	reportCourse := []report.ReportCourse{
		{ID: 1, CourseName: "One"},
		{ID: 2, CourseName: "Two"},
	}
	mockreportCourse.EXPECT().Delete(context.TODO(), uint32(1)).DoAndReturn(func(ctx context.Context, id uint32) error {
		removeSlice := func(report []report.ReportCourse, index int) []report.ReportCourse {
			return append(report[:index], report[index+1:]...)

		}
		for key, ass := range reportCourse {
			if ass.ID == id {
				reportCourse = removeSlice(reportCourse, key)
				break
			}
		}
		return nil
	})
	useCase := NewReportCourseUseCase(mockreportCourse)
	err := useCase.Delete(context.TODO(), 1)
	assert.NilError(t, err)
	assert.Equal(t, len(reportCourse), 1)
	assert.Equal(t, reportCourse[0].ID, uint32(2))
}

func TestCreateOrUpdateCourseID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportCourse := mock_assigment.NewMockCourseRepositoryPG(ctrl)
	reportCourse := report.ReportCourse{
		CourseID: 1,
	}
	mockReportCourse.EXPECT().FindByCourseIDDateNow(context.Background(), uint32(1)).Return(report.ReportCourse{}, nil)
	mockReportCourse.EXPECT().Create(context.Background(), &reportCourse).DoAndReturn(
		func(ctx context.Context, m *report.ReportCourse) error {
			m.ID = 2
			return nil
		})

	useCase := NewReportCourseUseCase(mockReportCourse)
	err := useCase.CreateOrUpdateCourseID(context.Background(), &reportCourse)
	assert.NilError(t, err, "Error Create")
	assert.Equal(t, uint32(2), reportCourse.ID)
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

func TestCreateOrUpdateCourseIDReal(t *testing.T) {
	mockReportCourse := repository_real.NewCoursePG(RealSetup())
	reportCourse := report.ReportCourse{
		CourseID: 1,
	}
	useCase := NewReportCourseUseCase(mockReportCourse)
	err := useCase.CreateOrUpdateCourseID(context.Background(), &reportCourse)
	assert.NilError(t, err, "Error Create")
	assert.Equal(t, uint32(2), reportCourse.ID)
	t.Log(err)
	t.Fatalf("P")
}
