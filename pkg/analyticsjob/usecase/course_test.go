package usecase

import (
	"context"
	"sync"
	"testing"

	report_assigment_uc "github.com/abmid/icanvas-analytics/pkg/report/assigment/usecase/mock"
	report_course_uc "github.com/abmid/icanvas-analytics/pkg/report/course/usecase/mock"
	report_discussion_uc "github.com/abmid/icanvas-analytics/pkg/report/discussion/usecase/mock"
	report_enrollment_uc "github.com/abmid/icanvas-analytics/pkg/report/enrollment/usecase/mock"
	report "github.com/abmid/icanvas-analytics/pkg/report/entity"
	report_result_uc "github.com/abmid/icanvas-analytics/pkg/report/result/usecase/mock"
	canvas_assigment_uc "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/usecase/mock"
	canvas_course_uc "github.com/abmid/icanvas-analytics/pkg/canvas/course/usecase/mock"
	canvas_discussion_uc "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/usecase/mock"
	canvas_enrollment_uc "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/usecase/mock"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCreateReportCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	canvasCourse := canvas.Course{
		ID:        1,
		AccountID: 1,
		Name:      "Course Name",
	}
	// Mock Canvas Assigment UseCase
	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
	// Mock Canvas Course UseCase
	canvasCourseUC := canvas_course_uc.NewMockCourseUseCase(ctrl)
	// Mock Canvas Discussion UseCase
	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
	// Mock Canvas Enrollment UseCase
	canvasEnrollmentUC := canvas_enrollment_uc.NewMockEnrollmentUseCase(ctrl)

	// Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	// Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	reportCourse := report.ReportCourse{
		CourseID:   canvasCourse.ID,
		AccountID:  canvasCourse.AccountID,
		CourseName: canvasCourse.Name,
	}
	reportCourseUC.EXPECT().CreateOrUpdateCourseID(ctx, &reportCourse).DoAndReturn(func(ctx context.Context, m *report.ReportCourse) error {
		m.ID = 1
		return nil
	})
	// Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	// Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	reportCourseID, err := AUC.createReportCourse(canvasCourse)
	assert.NilError(t, err)
	assert.Equal(t, reportCourseID, uint32(1))
}

func TestListReportCourse(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// TODO Mock Canvas Assigment UseCase
	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
	// TODO Mock Canvas Course UseCase
	canvasCourseUC := canvas_course_uc.NewMockCourseUseCase(ctrl)
	// TODO Mock Canvas Discussion UseCase
	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
	// TODO Mock Canvas Enrollment UseCase
	canvasEnrollmentUC := canvas_enrollment_uc.NewMockEnrollmentUseCase(ctrl)

	// TODO Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	listReportCourse := []report.ReportCourse{
		{ID: 1, AccountID: 39},
		{ID: 2, AccountID: 39},
	}
	filter := report.ReportCourse{}
	reportCourseUC.EXPECT().FindFilter(ctx, filter).Return(listReportCourse, nil)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	ch := make(chan report.ReportCourse)
	wg := new(sync.WaitGroup)
	result := []report.ReportCourse{}
	go func() {
		for rCourse := range ch {
			// t.Log(rCourse)
			result = append(result, rCourse)
			wg.Done()
		}
	}()
	AUC.listReportCourse(filter, ch, wg)
	assert.Equal(t, len(result), len(listReportCourse))
}
