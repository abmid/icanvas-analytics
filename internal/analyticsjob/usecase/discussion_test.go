package usecase

import (
	"context"
	"sync"
	"testing"

	report_assigment_uc "github.com/abmid/icanvas-analytics/internal/report/assigment/usecase/mock"
	report_course_uc "github.com/abmid/icanvas-analytics/internal/report/course/usecase/mock"
	report_discussion_uc "github.com/abmid/icanvas-analytics/internal/report/discussion/usecase/mock"
	report_enrollment_uc "github.com/abmid/icanvas-analytics/internal/report/enrollment/usecase/mock"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"
	report_result_uc "github.com/abmid/icanvas-analytics/internal/report/result/usecase/mock"
	canvas_assigment_uc "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/usecase/mock"
	canvas_course_uc "github.com/abmid/icanvas-analytics/pkg/canvas/course/usecase/mock"
	canvas_discussion_uc "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/usecase/mock"
	canvas_enrollment_uc "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/usecase/mock"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestListDiscussion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// TODO Mock Canvas Assigment UseCase
	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
	// TODO Mock Canvas Course UseCase
	canvasCourseUC := canvas_course_uc.NewMockCourseUseCase(ctrl)
	// TODO Mock Canvas Discussion UseCase
	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
	discussions := []entity.Discussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
	}
	canvasDiscussionUC.EXPECT().ListDiscussionByCourseID(uint32(1)).Return(discussions, nil)
	// TODO Mock Canvas Enrollment UseCase
	canvasEnrollmentUC := canvas_enrollment_uc.NewMockEnrollmentUseCase(ctrl)

	// TODO Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	res, err := AUC.listDiscussion(uint32(1))
	assert.NilError(t, err)
	assert.Equal(t, len(discussions), len(res))
}

func TestDispatchWorkerCreateReportDiscussion(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.TODO(), t)
	defer ctrl.Finish()
	// TODO Mock Canvas Assigment UseCase
	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
	// TODO Mock Canvas Course UseCase
	canvasCourseUC := canvas_course_uc.NewMockCourseUseCase(ctrl)
	// TODO Mock Canvas Discussion UseCase
	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
	discussions := []canvas.Discussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
	}
	// TODO Mock Canvas Enrollment UseCase
	canvasEnrollmentUC := canvas_enrollment_uc.NewMockEnrollmentUseCase(ctrl)

	// TODO Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	reportDiss := []report.ReportDiscussion{
		{CourseReportID: 1, Title: discussions[0].Title, DiscussionID: discussions[0].ID},
		{CourseReportID: 1, Title: discussions[1].Title, DiscussionID: discussions[1].ID},
	}
	reportDiscussionUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), reportDiss[0], &reportDiss[0]).Return(nil).AnyTimes()
	reportDiscussionUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), reportDiss[1], &reportDiss[1]).Return(nil).AnyTimes()
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	wg := new(sync.WaitGroup)
	in := make(chan canvas.Discussion)
	go AUC.dispatchWorkerCreateReportDiscussion(context.TODO(), wg, in, uint32(1))
	for _, discussion := range discussions {
		wg.Add(1)
		in <- discussion
	}
	wg.Wait()
}

func TestCreateReportDiscussion(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.TODO(), t)
	defer ctrl.Finish()
	// TODO Mock Canvas Assigment UseCase
	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
	// TODO Mock Canvas Course UseCase
	canvasCourseUC := canvas_course_uc.NewMockCourseUseCase(ctrl)
	// TODO Mock Canvas Discussion UseCase
	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
	discussions := []canvas.Discussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
	}
	canvasDiscussionUC.EXPECT().ListDiscussionByCourseID(uint32(1)).Return(discussions, nil)
	// TODO Mock Canvas Enrollment UseCase
	canvasEnrollmentUC := canvas_enrollment_uc.NewMockEnrollmentUseCase(ctrl)

	// TODO Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	reportDiss := []report.ReportDiscussion{
		{CourseReportID: 1, Title: discussions[0].Title, DiscussionID: discussions[0].ID},
		{CourseReportID: 1, Title: discussions[1].Title, DiscussionID: discussions[1].ID},
	}
	reportDiscussionUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), reportDiss[0], &reportDiss[0]).Return(nil).AnyTimes()
	reportDiscussionUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), reportDiss[1], &reportDiss[1]).Return(nil).AnyTimes()
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// TODO Mock Report Result UseCase
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	wg := new(sync.WaitGroup)
	out := make(chan []canvas.Discussion)
	wg.Add(1)
	go AUC.createReportDiscussion(wg, out, context.TODO(), uint32(1), uint32(1))
	for i := 0; i < 1; i++ {
		select {
		case d := <-out:
			for key, discuss := range d {
				assert.Equal(t, discuss.ID, discussions[key].ID)
			}
		}
	}
	wg.Wait()
}

func TestListReportDiscussion(t *testing.T) {
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
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	filter := report.ReportDiscussion{
		CourseReportID: 1,
	}
	listReportDiscussion := []report.ReportDiscussion{
		{CourseReportID: 1, ID: 1},
		{CourseReportID: 1, ID: 2},
	}
	reportDiscussionUC.EXPECT().FindFilter(ctx, filter).Return(listReportDiscussion, nil)
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// TODO Mock Report Result UseCase
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)

	res, err := AUC.listReportDiscussion(ctx, filter)
	assert.NilError(t, err)
	assert.Equal(t, len(listReportDiscussion), len(res))
}
