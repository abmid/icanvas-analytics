/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"
	"sync"

	logger "github.com/abmid/icanvas-analytics/internal/logger"
	assigment_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/usecase"
	course_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/course/usecase"
	discussion_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/usecase"
	enrollment_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/usecase"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	report_assigment_usecase "github.com/abmid/icanvas-analytics/pkg/report/assigment/usecase"
	report_course_usecase "github.com/abmid/icanvas-analytics/pkg/report/course/usecase"
	report_discussion_usecase "github.com/abmid/icanvas-analytics/pkg/report/discussion/usecase"
	report_enrollment_usecase "github.com/abmid/icanvas-analytics/pkg/report/enrollment/usecase"
	report "github.com/abmid/icanvas-analytics/pkg/report/entity"
	report_result_usecase "github.com/abmid/icanvas-analytics/pkg/report/result/usecase"
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
	Log              *logger.LoggerWrap
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

	logger := logger.New()
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
		Log:              logger,
	}
}

const (
	totalWorker    = 150
	totalWorkerJob = 220
)

// RunJob
func (AUC *AnalyticJobUseCase) RunJob(accountID uint32) <-chan bool {
	wg := new(sync.WaitGroup)
	finish := make(chan bool)
	// Pipeline
	chCourses := make(chan []canvas.Course)
	// Running Worker and wait value upstream via inbound channel (chCourses)
	go AUC.dispatchWorkerJob(chCourses, wg)
	// Get all course and send value downstream via outbound channel (chCourse)
	AUC.Course.GoAllCourse(accountID, chCourses, wg)

	go func() {
		wg.Wait()
		finish <- true
	}()

	return finish
}

// dispatchWorkerJob a function go routine that runs as many totalWorkerJob and wait value from upstream via inbound channel
func (AUC *AnalyticJobUseCase) dispatchWorkerJob(chCourses <-chan []canvas.Course, wg *sync.WaitGroup) {
	// Init running worker same like var totalWorkerJob
	for i := 0; i < totalWorkerJob; i++ {
		// Create concurrent
		go func(in <-chan []canvas.Course, wg *sync.WaitGroup) {
			// Waiting Receive data from channel
			for courses := range in {

				for _, course := range courses {
					// Create report course.
					reportCourseID, err := AUC.createReportCourse(course)
					if err != nil {
						AUC.Log.Error(err)
						continue
						// The future will be save into database for better information
					}
					// Analyze Assigment, Discussion, Enrollment and store to result report course
					err = AUC.doAnalyzeReportCourseJob(reportCourseID, course.ID)
					if err != nil {
						AUC.Log.Error(err)
						continue
						// The future will be save into database for better information
					}
				}
				// After receive data from channel and process in above wg.Done
				wg.Done()
			}
		}(chCourses, wg)
	}
}

// doAnalyzeReportCourseJob analyze how many assigment, discussion, finish enrollment and save to DB report_course (result)
func (AUC *AnalyticJobUseCase) doAnalyzeReportCourseJob(reportCourseID, courseID uint32) error {

	wg := new(sync.WaitGroup)

	chListAssigment := make(chan []canvas.Assigment)
	chListDiscussion := make(chan []canvas.Discussion)
	chListEnrollment := make(chan []canvas.Enrollment)

	wg.Add(3)
	// Running Worker Create Report and send result to channel
	go AUC.createReportAssigment(wg, chListAssigment, context.TODO(), reportCourseID, courseID)
	go AUC.createReportDiscussion(wg, chListDiscussion, context.TODO(), reportCourseID, courseID)
	go AUC.createReportEnrollment(wg, chListEnrollment, context.TODO(), reportCourseID, courseID)

	var listAssigment []canvas.Assigment
	var listDiscussion []canvas.Discussion
	var listEnrollment []canvas.Enrollment
	// Wait channel send data and save to list
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

	// Get Count Assigment
	assigmentCount := len(listAssigment)
	// Get Count Discussion
	discussionCount := len(listDiscussion)
	// Check Score Grade
	studentCount, finishGrading, averageGrading := AUC.CheckScoreGrade(listEnrollment)
	// Fill report result
	reportResult := report.ReportResult{
		AssigmentCount:     uint32(assigmentCount),
		DiscussionCount:    uint32(discussionCount),
		StudentCount:       studentCount,
		FinishGradingCount: finishGrading,
		FinalScore:         (float32(assigmentCount) + float32(discussionCount) + averageGrading) / 3,
		ReportCourseID:     reportCourseID,
	}

	// Svae to DB Report Course (Result Report).
	err := AUC.ReportResult.CreateOrUpdateByCourseReportID(context.TODO(), &reportResult)
	if err != nil {
		AUC.Log.Error(err)
		return err
	}

	wg.Wait()
	return nil
}
