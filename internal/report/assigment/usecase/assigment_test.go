package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	mock_assigment "github.com/abmid/icanvas-analytics/internal/report/assigment/repository/mock"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportAss := mock_assigment.NewMockAssigmentRepositoryPG(ctrl)
	reportAss := report.ReportAssigment{
		CourseReportID: 1,
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
		Name:           "Name Report Ass",
	}
	mockReportAss.EXPECT().Create(context.Background(), &reportAss).DoAndReturn(
		func(ctx context.Context, m *report.ReportAssigment) error {
			m.ID = 2
			return nil
		})
	useCase := NewReportAssigmentUseCase(mockReportAss)
	err := useCase.Create(context.Background(), &reportAss)
	assert.NilError(t, err, "Error Create")
	assert.Equal(t, uint32(2), reportAss.ID)
}

func TestRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportAss := mock_assigment.NewMockAssigmentRepositoryPG(ctrl)
	reportAss := []report.ReportAssigment{
		{ID: 1, Name: "One"},
		{ID: 2, Name: "Two"},
	}
	mockReportAss.EXPECT().Read(context.TODO()).Return(reportAss, nil)
	useCase := NewReportAssigmentUseCase(mockReportAss)
	result, err := useCase.Read(context.TODO())
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(result), len(reportAss))
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportAss := mock_assigment.NewMockAssigmentRepositoryPG(ctrl)
	timeUpdate := time.Now()
	reportAss := report.ReportAssigment{
		ID:             1,
		Name:           "Title",
		CourseReportID: 1,
	}
	mockReportAss.EXPECT().Update(context.TODO(), &reportAss).DoAndReturn(func(ctx context.Context, m *report.ReportAssigment) error {
		m.ID = reportAss.ID
		m.Name = reportAss.Name
		m.CourseReportID = reportAss.CourseReportID
		m.UpdatedAt = sql.NullTime{Time: timeUpdate}
		return nil
	})
	useCase := NewReportAssigmentUseCase(mockReportAss)
	err := useCase.Update(context.TODO(), &reportAss)
	assert.NilError(t, err)
	assert.Equal(t, timeUpdate, reportAss.UpdatedAt.Time)
}

func removeSlice(report []report.ReportAssigment, index int) []report.ReportAssigment {
	return append(report[:index], report[index+1:]...)
}
func TestRemoveSlice(t *testing.T) {
	reportAss := []report.ReportAssigment{
		{ID: 1, Name: "One"},
		{ID: 2, Name: "Two"},
	}
	for index, ass := range reportAss {
		if ass.ID == 1 {
			reportAss = removeSlice(reportAss, index)
		}
	}
	assert.Equal(t, len(reportAss), 1)
	assert.Equal(t, reportAss[0].ID, uint32(2))
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportAss := mock_assigment.NewMockAssigmentRepositoryPG(ctrl)
	reportAss := []report.ReportAssigment{
		{ID: 1, Name: "One"},
		{ID: 2, Name: "Two"},
	}
	mockReportAss.EXPECT().Delete(context.TODO(), uint32(1)).DoAndReturn(func(ctx context.Context, id uint32) error {
		removeSlice := func(report []report.ReportAssigment, index int) []report.ReportAssigment {
			return append(report[:index], report[index+1:]...)

		}
		for key, ass := range reportAss {
			if ass.ID == id {
				reportAss = removeSlice(reportAss, key)
				break
			}
		}
		return nil
	})
	useCase := NewReportAssigmentUseCase(mockReportAss)
	err := useCase.Delete(context.TODO(), 1)
	assert.NilError(t, err)
	assert.Equal(t, len(reportAss), 1)
	assert.Equal(t, reportAss[0].ID, uint32(2))
}

func TestCreateOrUpdateByFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportAss := mock_assigment.NewMockAssigmentRepositoryPG(ctrl)
	t.Run("create", func(t *testing.T) {
		filter := report.ReportAssigment{
			AssigmentID: 1,
		}
		resMockReportAss := report.ReportAssigment{}
		mockReportAss.EXPECT().FindFirstByFilter(context.Background(), filter).Return(resMockReportAss, nil)
		mockReportAss.EXPECT().Create(context.Background(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, m *report.ReportAssigment) error {
				m.ID = 2
				return nil
			})
		uc := NewReportAssigmentUseCase(mockReportAss)
		submitReport := report.ReportAssigment{
			Name:           "Title",
			AssigmentID:    1,
			CourseReportID: 1,
		}
		err := uc.CreateOrUpdateByFilter(context.Background(), filter, &submitReport)
		t.Log(err)
		t.Log(submitReport)
		assert.NilError(t, err)
		assert.Equal(t, submitReport.ID, uint32(2))
	})
	t.Run("update", func(t *testing.T) {
		timeUpdate := time.Now()
		filter := report.ReportAssigment{
			AssigmentID: 1,
		}
		resMockReportAss := report.ReportAssigment{
			ID:          1,
			AssigmentID: 1,
			Name:        "Title",
		}
		mockReportAss.EXPECT().FindFirstByFilter(context.Background(), filter).Return(resMockReportAss, nil)
		submitReport := report.ReportAssigment{
			ID:             1,
			Name:           "Update",
			AssigmentID:    1,
			CourseReportID: 1,
		}
		mockReportAss.EXPECT().Update(context.Background(), &submitReport).DoAndReturn(
			func(ctx context.Context, m *report.ReportAssigment) error {
				m.ID = submitReport.ID
				m.AssigmentID = submitReport.AssigmentID
				m.Name = submitReport.Name
				m.UpdatedAt.Time = timeUpdate
				return nil
			})
		uc := NewReportAssigmentUseCase(mockReportAss)
		err := uc.CreateOrUpdateByFilter(context.Background(), filter, &submitReport)
		t.Log(err)
		t.Log(submitReport)
		assert.NilError(t, err)
		assert.Equal(t, submitReport.UpdatedAt.Time, timeUpdate)
	})
}

func TestFindFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportAss := mock_assigment.NewMockAssigmentRepositoryPG(ctrl)
	filter := report.ReportAssigment{}
	reportAss := []report.ReportAssigment{
		{ID: 1, Name: "One"},
		{ID: 2, Name: "Two"},
	}
	mockReportAss.EXPECT().FindFilter(context.TODO(), filter).Return(reportAss, nil)
	useCase := NewReportAssigmentUseCase(mockReportAss)
	result, err := useCase.FindFilter(context.TODO(), filter)
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(result), len(reportAss))
}
