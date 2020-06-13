package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	mock_discussion "github.com/abmid/icanvas-analytics/internal/report/discussion/repository/mock"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportDiss := mock_discussion.NewMockDisscussionRepositoryPG(ctrl)
	reportDiss := report.ReportDiscussion{
		CourseReportID: 1,
		CreatedAt:      sql.NullTime{Time: time.Now()},
		UpdatedAt:      sql.NullTime{Time: time.Now()},
		Title:          "Title",
	}
	mockReportDiss.EXPECT().Create(context.Background(), &reportDiss).DoAndReturn(
		func(ctx context.Context, m *report.ReportDiscussion) error {
			m.ID = 2
			return nil
		})
	useCase := NewReportDiscussionUseCase(mockReportDiss)
	err := useCase.Create(context.Background(), &reportDiss)
	assert.NilError(t, err, "Error Create")
	assert.Equal(t, uint32(2), reportDiss.ID)
}

func TestRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportDiss := mock_discussion.NewMockDisscussionRepositoryPG(ctrl)
	reportAss := []report.ReportDiscussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
	}
	mockReportDiss.EXPECT().Read(context.TODO()).Return(reportAss, nil)
	useCase := NewReportDiscussionUseCase(mockReportDiss)
	result, err := useCase.Read(context.TODO())
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(result), len(reportAss))
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportDiss := mock_discussion.NewMockDisscussionRepositoryPG(ctrl)
	timeUpdate := time.Now()
	reportAss := report.ReportDiscussion{
		ID:             1,
		Title:          "Title",
		CourseReportID: 1,
	}
	mockReportDiss.EXPECT().Update(context.TODO(), &reportAss).DoAndReturn(func(ctx context.Context, m *report.ReportDiscussion) error {
		m.ID = reportAss.ID
		m.Title = reportAss.Title
		m.CourseReportID = reportAss.CourseReportID
		m.UpdatedAt = sql.NullTime{Time: timeUpdate}
		return nil
	})
	useCase := NewReportDiscussionUseCase(mockReportDiss)
	err := useCase.Update(context.TODO(), &reportAss)
	assert.NilError(t, err)
	assert.Equal(t, timeUpdate, reportAss.UpdatedAt.Time)
}

func removeSlice(report []report.ReportDiscussion, index int) []report.ReportDiscussion {
	return append(report[:index], report[index+1:]...)
}
func TestRemoveSlice(t *testing.T) {
	reportAss := []report.ReportDiscussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
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
	mockReportDiss := mock_discussion.NewMockDisscussionRepositoryPG(ctrl)
	reportAss := []report.ReportDiscussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
	}
	mockReportDiss.EXPECT().Delete(context.TODO(), uint32(1)).DoAndReturn(func(ctx context.Context, id uint32) error {
		removeSlice := func(report []report.ReportDiscussion, index int) []report.ReportDiscussion {
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
	useCase := NewReportDiscussionUseCase(mockReportDiss)
	err := useCase.Delete(context.TODO(), 1)
	assert.NilError(t, err)
	assert.Equal(t, len(reportAss), 1)
	assert.Equal(t, reportAss[0].ID, uint32(2))
}

func TestCreateOrUpdateByFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReportDiss := mock_discussion.NewMockDisscussionRepositoryPG(ctrl)
	t.Run("create", func(t *testing.T) {
		filter := report.ReportDiscussion{
			DiscussionID: 1,
		}
		resMockReportDiss := report.ReportDiscussion{}
		mockReportDiss.EXPECT().FindFirstByFilter(context.Background(), filter).Return(resMockReportDiss, nil)
		mockReportDiss.EXPECT().Create(context.Background(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, m *report.ReportDiscussion) error {
				m.ID = 2
				return nil
			})
		uc := NewReportDiscussionUseCase(mockReportDiss)
		submitReport := report.ReportDiscussion{
			Title:          "Title",
			DiscussionID:   1,
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
		filter := report.ReportDiscussion{
			DiscussionID: 1,
		}
		resMockReportDiss := report.ReportDiscussion{
			ID:           1,
			DiscussionID: 1,
			Title:        "Title",
		}
		mockReportDiss.EXPECT().FindFirstByFilter(context.Background(), filter).Return(resMockReportDiss, nil)
		submitReport := report.ReportDiscussion{
			ID:             1,
			Title:          "Update",
			DiscussionID:   1,
			CourseReportID: 1,
		}
		mockReportDiss.EXPECT().Update(context.Background(), &submitReport).DoAndReturn(
			func(ctx context.Context, m *report.ReportDiscussion) error {
				m.ID = submitReport.ID
				m.DiscussionID = submitReport.DiscussionID
				m.Title = submitReport.Title
				m.UpdatedAt.Time = timeUpdate
				return nil
			})
		uc := NewReportDiscussionUseCase(mockReportDiss)
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
	mockReportDiss := mock_discussion.NewMockDisscussionRepositoryPG(ctrl)
	filter := report.ReportDiscussion{
		CourseReportID: 1,
	}
	reportAss := []report.ReportDiscussion{
		{ID: 1, Title: "One", CourseReportID: 1},
		{ID: 2, Title: "Two", CourseReportID: 1},
	}
	mockReportDiss.EXPECT().FindFilter(context.TODO(), filter).Return(reportAss, nil)
	useCase := NewReportDiscussionUseCase(mockReportDiss)
	result, err := useCase.FindFilter(context.TODO(), filter)
	assert.NilError(t, err, "Error Read")
	assert.Equal(t, len(result), len(reportAss))
}
