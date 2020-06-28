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

func TestDoAnalyzeReportCourseJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.TODO()
	defer ctrl.Finish()
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
	discussions := []canvas.Discussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
	}
	canvasDiscussionUC.EXPECT().ListDiscussionByCourseID(uint32(1)).Return(discussions, nil)
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
	// ? =============
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
	reportDiss := []report.ReportDiscussion{
		{CourseReportID: 1, Title: discussions[0].Title, DiscussionID: discussions[0].ID},
		{CourseReportID: 1, Title: discussions[1].Title, DiscussionID: discussions[1].ID},
	}
	reportDiscussionUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), reportDiss[0], &reportDiss[0]).Return(nil).AnyTimes()
	reportDiscussionUC.EXPECT().CreateOrUpdateByFilter(context.TODO(), reportDiss[1], &reportDiss[1]).Return(nil).AnyTimes()
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
	assigmentCount := len(listAssigment)
	discussionCount := len(discussions)
	averageGrading := float32(1) / float32(2) * 100
	reportResult := report.ReportResult{
		AssigmentCount:     uint32(assigmentCount),
		DiscussionCount:    uint32(discussionCount),
		StudentCount:       uint32(len(ListEnrollment)),
		FinishGradingCount: 1,
		FinalScore:         (float32(assigmentCount) + float32(discussionCount) + averageGrading) / 3,
		ReportCourseID:     1,
	}
	reportResultUC.EXPECT().CreateOrUpdateByCourseReportID(ctx, &reportResult).Return(nil).AnyTimes()
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	err := AUC.doAnalyzeReportCourseJob(uint32(1), uint32(1))
	assert.NilError(t, err)
}

func TestDispatchWorkerJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.TODO()
	defer ctrl.Finish()
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
	discussions := []canvas.Discussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
	}
	canvasDiscussionUC.EXPECT().ListDiscussionByCourseID(uint32(1)).Return(discussions, nil)
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
	// ? =============
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
	reportCourse := report.ReportCourse{
		CourseID: 1, CourseName: "Test", AccountID: 1,
	}
	reportCourseUC.EXPECT().CreateOrUpdateCourseID(context.TODO(), &reportCourse).DoAndReturn(func(ctx context.Context, m *report.ReportCourse) error {
		m.ID = 1 // Set ID
		m.CourseID = reportCourse.CourseID
		m.CourseName = reportCourse.CourseName
		m.AccountID = reportCourse.AccountID
		return nil
	})
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
	assigmentCount := len(listAssigment)
	discussionCount := len(discussions)
	averageGrading := float32(1) / float32(2) * 100
	reportResult := report.ReportResult{
		AssigmentCount:     uint32(assigmentCount),
		DiscussionCount:    uint32(discussionCount),
		StudentCount:       uint32(len(ListEnrollment)),
		FinishGradingCount: 1,
		FinalScore:         (float32(assigmentCount) + float32(discussionCount) + averageGrading) / 3,
		ReportCourseID:     1,
	}
	reportResultUC.EXPECT().CreateOrUpdateByCourseReportID(ctx, &reportResult).Return(nil).AnyTimes()
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	wg := new(sync.WaitGroup)
	ch := make(chan []canvas.Course)
	courses := []canvas.Course{
		{ID: 1, Name: "Test", AccountID: 1},
	}
	go AUC.dispatchWorkerJob(ch, wg)
	for i := 0; i < 1; i++ {
		wg.Add(1)
		ch <- courses
	}
	wg.Wait()
}

func TestRunJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.TODO()
	defer ctrl.Finish()
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
	canvasCourseUC.EXPECT().GoAllCourse(uint32(1), gomock.Any(), gomock.Any()).DoAndReturn(func(accountID uint32, ch chan<- []canvas.Course, wg *sync.WaitGroup) {
		wg.Add(1)
		ch <- []canvas.Course{
			{ID: 1, Name: "Test", AccountID: 1},
		}
	})

	// TODO Mock Canvas Discussion UseCase
	canvasDiscussionUC := canvas_discussion_uc.NewMockDiscussionUseCase(ctrl)
	discussions := []canvas.Discussion{
		{ID: 1, Title: "One"},
		{ID: 2, Title: "Two"},
	}
	canvasDiscussionUC.EXPECT().ListDiscussionByCourseID(uint32(1)).Return(discussions, nil)
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
	// ? =============
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
	reportCourse := report.ReportCourse{
		CourseID: 1, CourseName: "Test", AccountID: 1,
	}
	reportCourseUC.EXPECT().CreateOrUpdateCourseID(context.TODO(), &reportCourse).DoAndReturn(func(ctx context.Context, m *report.ReportCourse) error {
		m.ID = 1 // Set ID
		m.CourseID = reportCourse.CourseID
		m.CourseName = reportCourse.CourseName
		m.AccountID = reportCourse.AccountID
		return nil
	})
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
	assigmentCount := len(listAssigment)
	discussionCount := len(discussions)
	averageGrading := float32(1) / float32(2) * 100
	reportResult := report.ReportResult{
		AssigmentCount:     uint32(assigmentCount),
		DiscussionCount:    uint32(discussionCount),
		StudentCount:       uint32(len(ListEnrollment)),
		FinishGradingCount: 1,
		FinalScore:         (float32(assigmentCount) + float32(discussionCount) + averageGrading) / 3,
		ReportCourseID:     1,
	}
	reportResultUC.EXPECT().CreateOrUpdateByCourseReportID(ctx, &reportResult).Return(nil).AnyTimes()
	AUC := NewAnalyticJobUseCase(canvasCourseUC, canvasAssigmentUC, canvasEnrollmentUC, canvasDiscussionUC, reportAssigmentUC, reportCourseUC, reportDiscussionUC, reportEnrollmentUC, reportResultUC)
	AUC.RunJob(uint32(1))
}
