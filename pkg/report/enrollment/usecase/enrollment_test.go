package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	mock_enrollment "github.com/abmid/icanvas-analytics/pkg/report/enrollment/repository/mock"
	report "github.com/abmid/icanvas-analytics/pkg/report/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportEnroll := mock_enrollment.NewMockEnrollmentRepositoryPG(ctrl)
	reportEnroll := report.ReportEnrollment{
		CourseReportID: 1,
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
	}
	mockReportEnroll.EXPECT().Create(context.Background(), &reportEnroll).DoAndReturn(
		func(ctx context.Context, m *report.ReportEnrollment) error {
			m.ID = 2
			return nil
		})
	useCase := NewReportEnrollUseCase(mockReportEnroll)
	err := useCase.Create(context.Background(), &reportEnroll)
	assert.NilError(t, err, "Error Create")
	assert.Equal(t, uint32(2), reportEnroll.ID)
}

func TestRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportEnroll := mock_enrollment.NewMockEnrollmentRepositoryPG(ctrl)
	reportEnroll := []report.ReportEnrollment{
		{ID: 1, Role: "One"},
		{ID: 2, Role: "Two"},
	}
	mockReportEnroll.EXPECT().Read(context.TODO()).Return(reportEnroll, nil)
	useCase := NewReportEnrollUseCase(mockReportEnroll)
	result, err := useCase.Read(context.TODO())
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(result), len(reportEnroll))
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportEnroll := mock_enrollment.NewMockEnrollmentRepositoryPG(ctrl)
	timeUpdate := time.Now()
	reportEnroll := report.ReportEnrollment{
		ID:             1,
		Role:           "Title",
		CourseReportID: 1,
	}
	mockReportEnroll.EXPECT().Update(context.TODO(), &reportEnroll).DoAndReturn(func(ctx context.Context, m *report.ReportEnrollment) error {
		m.ID = reportEnroll.ID
		m.Role = reportEnroll.Role
		m.CourseReportID = reportEnroll.CourseReportID
		m.UpdatedAt = sql.NullTime{Time: timeUpdate}
		return nil
	})
	useCase := NewReportEnrollUseCase(mockReportEnroll)
	err := useCase.Update(context.TODO(), &reportEnroll)
	assert.NilError(t, err)
	assert.Equal(t, timeUpdate, reportEnroll.UpdatedAt.Time)
}

func removeSlice(report []report.ReportEnrollment, index int) []report.ReportEnrollment {
	return append(report[:index], report[index+1:]...)
}
func TestRemoveSlice(t *testing.T) {
	reportEnroll := []report.ReportEnrollment{
		{ID: 1, Role: "One"},
		{ID: 2, Role: "Two"},
	}
	for index, ass := range reportEnroll {
		if ass.ID == 1 {
			reportEnroll = removeSlice(reportEnroll, index)
		}
	}
	assert.Equal(t, len(reportEnroll), 1)
	assert.Equal(t, reportEnroll[0].ID, uint32(2))
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportEnroll := mock_enrollment.NewMockEnrollmentRepositoryPG(ctrl)
	reportEnroll := []report.ReportEnrollment{
		{ID: 1, Role: "One"},
		{ID: 2, Role: "Two"},
	}
	mockReportEnroll.EXPECT().Delete(context.TODO(), uint32(1)).DoAndReturn(func(ctx context.Context, id uint32) error {
		removeSlice := func(report []report.ReportEnrollment, index int) []report.ReportEnrollment {
			return append(report[:index], report[index+1:]...)

		}
		for key, ass := range reportEnroll {
			if ass.ID == id {
				reportEnroll = removeSlice(reportEnroll, key)
				break
			}
		}
		return nil
	})
	useCase := NewReportEnrollUseCase(mockReportEnroll)
	err := useCase.Delete(context.TODO(), 1)
	assert.NilError(t, err)
	assert.Equal(t, len(reportEnroll), 1)
	assert.Equal(t, reportEnroll[0].ID, uint32(2))
}

func TestCreateOrUpdateByFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportEnroll := mock_enrollment.NewMockEnrollmentRepositoryPG(ctrl)
	t.Run("create", func(t *testing.T) {
		filter := report.ReportEnrollment{
			EnrollmentID: 1,
		}
		resMockReportDiss := report.ReportEnrollment{}
		mockReportEnroll.EXPECT().FindFirstByFilter(context.Background(), filter).Return(resMockReportDiss, nil)
		mockReportEnroll.EXPECT().Create(context.Background(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, m *report.ReportEnrollment) error {
				m.ID = 2
				return nil
			})
		uc := NewReportEnrollUseCase(mockReportEnroll)
		submitReport := report.ReportEnrollment{
			EnrollmentID:   1,
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
		filter := report.ReportEnrollment{
			EnrollmentID: 1,
		}
		resMockReportDiss := report.ReportEnrollment{
			ID:           1,
			EnrollmentID: 1,
		}
		mockReportEnroll.EXPECT().FindFirstByFilter(context.Background(), filter).Return(resMockReportDiss, nil)
		submitReport := report.ReportEnrollment{
			ID:             1,
			EnrollmentID:   1,
			CourseReportID: 1,
		}
		mockReportEnroll.EXPECT().Update(context.Background(), &submitReport).DoAndReturn(
			func(ctx context.Context, m *report.ReportEnrollment) error {
				m.ID = submitReport.ID
				m.EnrollmentID = submitReport.EnrollmentID
				m.UpdatedAt.Time = timeUpdate
				return nil
			})
		uc := NewReportEnrollUseCase(mockReportEnroll)
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
	mockReportEnroll := mock_enrollment.NewMockEnrollmentRepositoryPG(ctrl)
	filter := report.ReportEnrollment{
		CourseReportID: 1,
	}
	reportEnroll := []report.ReportEnrollment{
		{ID: 1, Role: "One", CourseReportID: 1},
		{ID: 2, Role: "Two", CourseReportID: 1},
	}
	mockReportEnroll.EXPECT().FindFilter(context.TODO(), filter).Return(reportEnroll, nil)
	useCase := NewReportEnrollUseCase(mockReportEnroll)
	result, err := useCase.FindFilter(context.TODO(), filter)
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(result), len(reportEnroll))
}
