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
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestListAssigment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// TODO Mock Canvas Assigment UseCase
	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
	listAssigment := []canvas.Assigment{
		{ID: 1, CourseID: 1, Name: "One"},
		{ID: 2, CourseID: 1, Name: "Two"},
	}
	canvasAssigmentUC.EXPECT().ListAssigmentByCourseID(uint32(1)).Return(listAssigment, nil)
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
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	res, err := AUC.listAssigment(uint32(1))
	assert.NilError(t, err)
	assert.Equal(t, len(listAssigment), len(res))
}

func TestDispatchWorkerCreateReportAssigment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	// TODO Mock Canvas Assigment UseCase
	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
	listAssigment := []canvas.Assigment{
		{ID: 1, CourseID: 1, Name: "One"},
		{ID: 2, CourseID: 1, Name: "Two"},
		{ID: 3, CourseID: 1, Name: "Three"},
		{ID: 4, CourseID: 1, Name: "Four"},
		{ID: 5, CourseID: 1, Name: "Five"},
		{ID: 6, CourseID: 1, Name: "Six"},
	}
	// TODO Mock Canvas Course UseCase
	canvasCourseUC := canvas_course_uc.NewMockCourseUseCase(ctrl)
	// TODO Mock Canvas Discussion UseCase
	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
	// TODO Mock Canvas Enrollment UseCase
	canvasEnrollmentUC := canvas_enrollment_uc.NewMockEnrollmentUseCase(ctrl)

	// TODO Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	reportAssigmentCreate := []report.ReportAssigment{
		{CourseReportID: 1, Name: listAssigment[0].Name, AssigmentID: 1},
		{CourseReportID: 1, Name: listAssigment[1].Name, AssigmentID: 2},
		{CourseReportID: 1, Name: listAssigment[2].Name, AssigmentID: 3},
		{CourseReportID: 1, Name: listAssigment[3].Name, AssigmentID: 4},
		{CourseReportID: 1, Name: listAssigment[4].Name, AssigmentID: 5},
		{CourseReportID: 1, Name: listAssigment[5].Name, AssigmentID: 6},
	}
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[0], &reportAssigmentCreate[0]).AnyTimes()
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[1], &reportAssigmentCreate[1]).AnyTimes()
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[2], &reportAssigmentCreate[2]).AnyTimes()
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[3], &reportAssigmentCreate[3]).AnyTimes()
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[4], &reportAssigmentCreate[4]).AnyTimes()
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[5], &reportAssigmentCreate[5]).AnyTimes()
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	wg := new(sync.WaitGroup)
	in := make(chan canvas.Assigment)
	go AUC.dispatchWorkerCreateReportAssigment(ctx, wg, in, uint32(1))
	for _, assigment := range listAssigment {
		wg.Add(1)
		in <- assigment
	}
	wg.Wait()
}

func TestCreateReportAssigment(t *testing.T) {
	t.Parallel()
	ctrl, _ := gomock.WithContext(context.TODO(), t)
	defer ctrl.Finish()
	ctx := context.TODO()
	// TODO Mock Canvas Assigment UseCase
	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
	listAssigment := []canvas.Assigment{
		{ID: 1, CourseID: 1, Name: "One"},
		{ID: 2, CourseID: 1, Name: "Two"},
		{ID: 3, CourseID: 1, Name: "Three"},
	}
	canvasAssigmentUC.EXPECT().ListAssigmentByCourseID(uint32(1)).Return(listAssigment, nil)
	// TODO Mock Canvas Course UseCase
	canvasCourseUC := canvas_course_uc.NewMockCourseUseCase(ctrl)
	// TODO Mock Canvas Discussion UseCase
	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
	// TODO Mock Canvas Enrollment UseCase
	canvasEnrollmentUC := canvas_enrollment_uc.NewMockEnrollmentUseCase(ctrl)

	// TODO Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	reportAssigmentCreate := []report.ReportAssigment{
		{CourseReportID: 1, Name: listAssigment[0].Name, AssigmentID: 1},
		{CourseReportID: 1, Name: listAssigment[1].Name, AssigmentID: 2},
		{CourseReportID: 1, Name: listAssigment[2].Name, AssigmentID: 3},
	}
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[0], &reportAssigmentCreate[0]).AnyTimes()
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[1], &reportAssigmentCreate[1]).AnyTimes()
	reportAssigmentUC.EXPECT().CreateOrUpdateByFilter(ctx, reportAssigmentCreate[2], &reportAssigmentCreate[2]).AnyTimes()
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	wg := new(sync.WaitGroup)
	out := make(chan []canvas.Assigment)
	wg.Add(1)
	go AUC.createReportAssigment(wg, out, ctx, uint32(1), uint32(1))
	for i := 0; i < 1; i++ {
		select {
		case a := <-out:
			for key, each := range a {
				assert.Equal(t, each.ID, listAssigment[key].ID)
			}
		}
	}
	wg.Wait()
}

func TestListReportAssigment(t *testing.T) {
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
	filter := report.ReportAssigment{
		CourseReportID: 1,
	}
	listReportAssigment := []report.ReportAssigment{
		{ID: 1, CourseReportID: 1, Name: "Test"},
		{ID: 2, CourseReportID: 1, Name: "Test two"},
	}
	reportAssigmentUC.EXPECT().FindFilter(ctx, filter).Return(listReportAssigment, nil)
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	res, err := AUC.listReportAssigment(ctx, filter)
	assert.NilError(t, err)
	assert.Equal(t, len(listReportAssigment), len(res))
}
