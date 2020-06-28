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

func TestListEnrollment(t *testing.T) {
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
	EnrollmentGrade := canvas.EnrollmentGrade{
		HtmlURL:      "html_url",
		CurrentGrade: 32,
		CurrentScore: 13,
		FinalScore:   12.2,
		FinalGrade:   11.2,
	}
	ListEnrollment := []canvas.Enrollment{
		{ID: 1, Grades: EnrollmentGrade, Role: "StudentEnrollment"},
		{ID: 2, Role: "StudentEnrollment"},
	}
	canvasEnrollmentUC.EXPECT().ListEnrollmentByCourseID(uint32(1)).Return(ListEnrollment, nil)

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
	res, err := AUC.listEnrollment(uint32(1))
	assert.NilError(t, err)
	assert.Equal(t, len(res), len(ListEnrollment))
}

// func TestCheckScoreGrade(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	// TODO Mock Canvas Assigment UseCase
// 	canvasAssigmentUC := canvas_assigment_uc.NewMockAssigmentUseCase(ctrl)
// 	// TODO Mock Canvas Course UseCase
// 	canvasCourseUC := canvas_course_uc.NewMockCourseUseCase(ctrl)
// 	// TODO Mock Canvas Discussion UseCase
// 	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
// 	// TODO Mock Canvas Enrollment UseCase
// 	canvasEnrollmentUC := canvas_enrollment_uc.NewMockEnrollmentUseCase(ctrl)

// 	// TODO Mock Report Assigment Usecase
// 	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
// 	// TODO Mock Report Course UseCase
// 	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
// 	// TODO Mock Report Discussion UseCase
// 	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
// 	// TODO Mock Report Enrollment UseCase
// 	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
// 	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC)

// 	listReportEnrollment := []report.ReportEnrollment{
// 		{ID: 1, CourseReportID: 1, UserID: 1, CurrentGrade: 32, CurrentScore: 13, FinalGrade: 11.2, FinalScore: 12.2, Role: "StudentEnrollment"},
// 		{ID: 1, CourseReportID: 1, UserID: 2, Role: "StudentEnrollment"},
// 	}
// 	studentCount, finishGrading, averageGrading := AUC.CheckScoreGrade(listReportEnrollment)
// 	assert.Equal(t, studentCount, uint32(len(listReportEnrollment)), "Count Student Not Same")
// 	assert.Equal(t, finishGrading, uint32(1), "Count grade not same")
// 	assert.Equal(t, averageGrading, float32(50), "Final Score not same")
// }

// func BenchmarkCheckScoreGrade(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		TestCheckScoreGrade(&testing.T{})
// 	}
// }

func TestDispatchWorkerCreateReportEnrollment(t *testing.T) {
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
	Users := []canvas.User{
		{ID: 1, Name: "1"},
		{ID: 2, Name: "2"},
	}
	EnrollmentGrade := canvas.EnrollmentGrade{
		HtmlURL:      "html_url",
		CurrentGrade: 32,
		CurrentScore: 13,
		FinalScore:   12.2,
		FinalGrade:   11.2,
	}
	ListEnrollment := []canvas.Enrollment{
		{ID: 1, Grades: EnrollmentGrade, Role: "StudentEnrollment", User: Users[0]},
		{ID: 2, Role: "StudentEnrollment", User: Users[1]},
	}

	// TODO Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	reportEnroll := []report.ReportEnrollment{
		{
			CourseReportID: 1,
			EnrollmentID:   ListEnrollment[0].ID,
			CurrentGrade:   ListEnrollment[0].Grades.CurrentGrade,
			CurrentScore:   ListEnrollment[0].Grades.CurrentScore,
			FinalGrade:     ListEnrollment[0].Grades.FinalGrade,
			FinalScore:     ListEnrollment[0].Grades.FinalScore,
			Role:           ListEnrollment[0].Role,
			RoleID:         ListEnrollment[0].RoleID,
			UserID:         ListEnrollment[0].UserID,
			LoginID:        ListEnrollment[0].User.LoginID,
			FullName:       ListEnrollment[0].User.Name,
		},
		{
			CourseReportID: 1,
			EnrollmentID:   ListEnrollment[1].ID,
			CurrentGrade:   ListEnrollment[1].Grades.CurrentGrade,
			CurrentScore:   ListEnrollment[1].Grades.CurrentScore,
			FinalGrade:     ListEnrollment[1].Grades.FinalGrade,
			FinalScore:     ListEnrollment[1].Grades.FinalScore,
			Role:           ListEnrollment[1].Role,
			RoleID:         ListEnrollment[1].RoleID,
			UserID:         ListEnrollment[1].UserID,
			LoginID:        ListEnrollment[1].User.LoginID,
			FullName:       ListEnrollment[1].User.Name,
		},
	}

	filter := []report.ReportEnrollment{
		{CourseReportID: 1, EnrollmentID: reportEnroll[0].EnrollmentID},
		{CourseReportID: 1, EnrollmentID: reportEnroll[1].EnrollmentID},
	}
	reportEnrollmentUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), filter[0], &reportEnroll[0]).Return(nil).AnyTimes()
	reportEnrollmentUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), filter[1], &reportEnroll[1]).Return(nil).AnyTimes()
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)

	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	wg := new(sync.WaitGroup)
	ch := make(chan canvas.Enrollment)
	go AUC.dispatchWorkerCreateReportEnrollment(context.TODO(), wg, ch, uint32(1))
	for _, enrollment := range ListEnrollment {
		wg.Add(1)
		ch <- enrollment
	}
	wg.Wait()
}

func TestCreateReportEnrollment(t *testing.T) {
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
	Users := []canvas.User{
		{ID: 1, Name: "1"},
		{ID: 2, Name: "2"},
	}
	EnrollmentGrade := canvas.EnrollmentGrade{
		HtmlURL:      "html_url",
		CurrentGrade: 32,
		CurrentScore: 13,
		FinalScore:   12.2,
		FinalGrade:   11.2,
	}
	ListEnrollment := []canvas.Enrollment{
		{ID: 1, Grades: EnrollmentGrade, Role: "StudentEnrollment", User: Users[0]},
		{ID: 2, Role: "StudentEnrollment", User: Users[1]},
	}
	canvasEnrollmentUC.EXPECT().ListEnrollmentByCourseID(uint32(1)).Return(ListEnrollment, nil)

	// TODO Mock Report Assigment Usecase
	reportAssigmentUC := report_assigment_uc.NewMockReportAssigmentUseCase(ctrl)
	// TODO Mock Report Course UseCase
	reportCourseUC := report_course_uc.NewMockReportCourseUseCase(ctrl)
	// TODO Mock Report Discussion UseCase
	reportDiscussionUC := report_discussion_uc.NewMockReportDiscussionUseCase(ctrl)
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	reportEnroll := []report.ReportEnrollment{
		{
			CourseReportID: 1,
			EnrollmentID:   ListEnrollment[0].ID,
			CurrentGrade:   ListEnrollment[0].Grades.CurrentGrade,
			CurrentScore:   ListEnrollment[0].Grades.CurrentScore,
			FinalGrade:     ListEnrollment[0].Grades.FinalGrade,
			FinalScore:     ListEnrollment[0].Grades.FinalScore,
			Role:           ListEnrollment[0].Role,
			RoleID:         ListEnrollment[0].RoleID,
			UserID:         ListEnrollment[0].UserID,
			LoginID:        ListEnrollment[0].User.LoginID,
			FullName:       ListEnrollment[0].User.Name,
		},
		{
			CourseReportID: 1,
			EnrollmentID:   ListEnrollment[1].ID,
			CurrentGrade:   ListEnrollment[1].Grades.CurrentGrade,
			CurrentScore:   ListEnrollment[1].Grades.CurrentScore,
			FinalGrade:     ListEnrollment[1].Grades.FinalGrade,
			FinalScore:     ListEnrollment[1].Grades.FinalScore,
			Role:           ListEnrollment[1].Role,
			RoleID:         ListEnrollment[1].RoleID,
			UserID:         ListEnrollment[1].UserID,
			LoginID:        ListEnrollment[1].User.LoginID,
			FullName:       ListEnrollment[1].User.Name,
		},
	}

	filter := []report.ReportEnrollment{
		{CourseReportID: 1, EnrollmentID: reportEnroll[0].EnrollmentID},
		{CourseReportID: 1, EnrollmentID: reportEnroll[1].EnrollmentID},
	}
	reportEnrollmentUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), filter[0], &reportEnroll[0]).Return(nil).AnyTimes()
	reportEnrollmentUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), filter[1], &reportEnroll[1]).Return(nil).AnyTimes()
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	wg := new(sync.WaitGroup)
	out := make(chan []canvas.Enrollment)
	wg.Add(1)
	go AUC.createReportEnrollment(wg, out, context.TODO(), uint32(1), uint32(1))
	for i := 0; i < 1; i++ {
		select {
		case e := <-out:
			for key, each := range e {
				assert.Equal(t, each.ID, ListEnrollment[key].ID)
			}
		}
	}
	wg.Wait()
}

func TestListReportEnrollment(t *testing.T) {
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
	// TODO Mock Report Enrollment UseCase
	reportEnrollmentUC := report_enrollment_uc.NewMockReportEnrollmentUseCase(ctrl)
	filter := report.ReportEnrollment{
		CourseReportID: 1,
	}
	listReportEnrollment := []report.ReportEnrollment{
		{ID: 1, CourseReportID: 1},
		{ID: 2, CourseReportID: 1},
	}
	reportEnrollmentUC.EXPECT().FindFilter(ctx, filter).Return(listReportEnrollment, nil)
	// TODO Mock Report Result
	reportResultUC := report_result_uc.NewMockReportResultUseCase(ctrl)

	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	res, err := AUC.listReportEnrollment(ctx, filter)
	assert.NilError(t, err)
	assert.Equal(t, len(listReportEnrollment), len(res))
}
