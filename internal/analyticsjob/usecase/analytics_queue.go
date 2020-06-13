package usecase

import (
	"context"
	"sync"

	report_assigment_usecase "github.com/abmid/icanvas-analytics/internal/report/assigment/usecase"
	report_course_usecase "github.com/abmid/icanvas-analytics/internal/report/course/usecase"
	report_discussion_usecase "github.com/abmid/icanvas-analytics/internal/report/discussion/usecase"
	report_enrollment_usecase "github.com/abmid/icanvas-analytics/internal/report/enrollment/usecase"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"
	report_result_usecase "github.com/abmid/icanvas-analytics/internal/report/result/usecase"
	assigment_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/usecase"
	course_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/course/usecase"
	discussion_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/usecase"
	enrollment_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/usecase"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/sirupsen/logrus"
)

type AnalyticJobUseCase struct {
	Course           course_usecase.CourseUseCase
	Assigment        assigment_usecase.AssigmentUseCase
	Enrollment       enrollment_usecase.EnrollmentUseCase
	Discussion       discussion_usecase.DiscussionUseCase
	ReportAssigment  report_assigment_usecase.ReportAssigmentUseCase
	ReportCourse     report_course_usecase.ReportCourseUseCase
	ReportDiscussion report_discussion_usecase.ReportDiscussionUseCase
	ReportEnrollment report_enrollment_usecase.ReportEnrollmentUseCase
	ReportResult     report_result_usecase.ReportResultUseCase
}

func NewAnalyticJobUseCase(
	course course_usecase.CourseUseCase,
	assigment assigment_usecase.AssigmentUseCase,
	enroll enrollment_usecase.EnrollmentUseCase,
	discuss discussion_usecase.DiscussionUseCase,
	reportAss report_assigment_usecase.ReportAssigmentUseCase,
	reportCourse report_course_usecase.ReportCourseUseCase,
	reportDiss report_discussion_usecase.ReportDiscussionUseCase,
	reportEnroll report_enrollment_usecase.ReportEnrollmentUseCase,
	reportResult report_result_usecase.ReportResultUseCase) *AnalyticJobUseCase {
	return &AnalyticJobUseCase{
		Course:           course,
		Assigment:        assigment,
		Enrollment:       enroll,
		Discussion:       discuss,
		ReportAssigment:  reportAss,
		ReportCourse:     reportCourse,
		ReportDiscussion: reportDiss,
		ReportEnrollment: reportEnroll,
		ReportResult:     reportResult,
	}
}

var (
	totalWorker    = 150
	totalWorkerJob = 220
)

func (AUC *AnalyticJobUseCase) RunJob(accountID uint32) {
	wg := new(sync.WaitGroup)
	chCourses := make(chan []canvas.Course)
	go AUC.dispatchWorkerJob(chCourses, wg)
	AUC.Course.GoAllCourse(accountID, chCourses, wg)
	wg.Wait()
}

/**
* This method for init running worker to analyze course from RunJob
 */
func (AUC *AnalyticJobUseCase) dispatchWorkerJob(chCourses <-chan []canvas.Course, wg *sync.WaitGroup) {
	// Init running worker same like var totalWorkerJob
	for i := 0; i < totalWorkerJob; i++ {
		// Create concurrent
		go func(in <-chan []canvas.Course, wg *sync.WaitGroup) {
			// Waiting Receive data from channel
			for courses := range in {
				// Create to pgsql in report_course to get id
				for _, course := range courses {
					reportCourseID, err := AUC.createReportCourse(course)
					if err != nil {
						panic(err)
					}
					err = AUC.doAnalyzeReportCourseJob(reportCourseID, course.ID)
					if err != nil {
						logrus.Error(err)
						continue
					}
				}
				// After receive data from channel and process in above wg.Done
				wg.Done()
			}
		}(chCourses, wg)
	}
}

func (AUC *AnalyticJobUseCase) doAnalyzeReportCourseJob(reportCourseID, courseID uint32) error {
	wg := new(sync.WaitGroup)
	chListAssigment := make(chan []canvas.Assigment)
	chListDiscussion := make(chan []canvas.Discussion)
	chListEnrollment := make(chan []canvas.Enrollment)
	wg.Add(3)
	go AUC.createReportAssigment(wg, chListAssigment, context.TODO(), reportCourseID, courseID)
	go AUC.createReportDiscussion(wg, chListDiscussion, context.TODO(), reportCourseID, courseID)
	go AUC.createReportEnrollment(wg, chListEnrollment, context.TODO(), reportCourseID, courseID)
	var listAssigment []canvas.Assigment
	var listDiscussion []canvas.Discussion
	var listEnrollment []canvas.Enrollment
	for i := 0; i < 3; i++ {
		select {
		case ass := <-chListAssigment:
			listAssigment = ass
		case diss := <-chListDiscussion:
			listDiscussion = diss
		case enroll := <-chListEnrollment:
			listEnrollment = enroll
		}
	}
	assigmentCount := len(listAssigment)
	discussionCount := len(listDiscussion)
	studentCount, finishGrading, averageGrading := AUC.CheckScoreGrade(listEnrollment)
	reportResult := report.ReportResult{
		AssigmentCount:     uint32(assigmentCount),
		DiscussionCount:    uint32(discussionCount),
		StudentCount:       studentCount,
		FinishGradingCount: finishGrading,
		FinalScore:         (float32(assigmentCount) + float32(discussionCount) + averageGrading) / 3,
		ReportCourseID:     reportCourseID,
	}

	err := AUC.ReportResult.CreateOrUpdateByCourseReportID(context.TODO(), &reportResult)
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

// TODO : 1 GET ALL COURSE 50
// TODO : 2. Run Worker to get Course 50 before
// TODO : 3. doAnalyzeReportCourseJob
// TODO : 3.1 Inside doAnalyzeReportCourseJob from go channel createReportAssigment send back channel list assigment
// TODO : 3.2 createReportDiscussion send back channel List discussion
// TODO : 3.3 createReportEnrollment send back channel list enrollment and then check score
