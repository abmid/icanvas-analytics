/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"
	"fmt"
	"sync"

	"github.com/abmid/icanvas-analytics/pkg/analyticsjob/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	report "github.com/abmid/icanvas-analytics/pkg/report/entity"
)

/*
This method for get List Enrollment By CourseID
*/
func (AUC *AnalyticJobUseCase) listEnrollment(courseID uint32) (res []canvas.Enrollment, err error) {
	countTry := 0
	for {
		enroll, err := AUC.Enrollment.ListEnrollmentByCourseID(courseID)
		if err != nil {
			if countTry > MAX_RETRY {
				break
			}
			countTry++
		} else {
			res = enroll
			break
		}
	}
	return res, nil
}

/*
Method for check how many student already give score
@param courseID int
@return studentcount, countGrade int, finalGrade float32
*/
func (AUC *AnalyticJobUseCase) CheckScoreGrade(reportEnrollments []canvas.Enrollment) (studentcount, finishGrading uint32, averadeGrading float32) {
	for _, enroll := range reportEnrollments {
		if enroll.Role == "StudentEnrollment" {
			if enroll.Grades.CurrentScore != 0 {
				finishGrading++
			}
			studentcount++
		}
	}
	if finishGrading != 0 {
		averadeGrading = float32(finishGrading) / float32(studentcount) * 100
	}
	return studentcount, finishGrading, averadeGrading
}

// dispatchWorkerCreateReportEnrollment a function running go routine such as many worker database. This function wait value from upstream via inbound channel
func (AUC *AnalyticJobUseCase) dispatchWorkerCreateReportEnrollment(ctx context.Context, wg *sync.WaitGroup, in <-chan canvas.Enrollment, reportCourseID uint32) {
	for i := 0; i < WORKER_DATABASE; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup, in <-chan canvas.Enrollment, reportCourseID uint32) {
			for enrollment := range in {
				reportEnroll := report.ReportEnrollment{
					CourseReportID: reportCourseID,
					EnrollmentID:   enrollment.ID,
					CurrentGrade:   enrollment.Grades.CurrentGrade,
					CurrentScore:   enrollment.Grades.CurrentScore,
					FinalGrade:     enrollment.Grades.FinalGrade,
					FinalScore:     enrollment.Grades.FinalScore,
					Role:           enrollment.Role,
					RoleID:         enrollment.RoleID,
					UserID:         enrollment.UserID,
					LoginID:        enrollment.User.LoginID,
					FullName:       enrollment.User.Name,
				}
				filter := report.ReportEnrollment{
					CourseReportID: reportCourseID,
					EnrollmentID:   enrollment.ID,
				}
				countTry := 0
				for {
					err := AUC.ReportEnrollment.CreateOrUpdateByFilter(ctx, filter, &reportEnroll)
					if err != nil {
						if countTry > MAX_RETRY {
							break
						}
						countTry++
					} else {
						break
					}
				}
				wg.Done()
			}
		}(ctx, wg, in, reportCourseID)
	}
}

// createReportEnrollment a function to create report enrollment, in this function will have 2 operation after get list diss.
// 1. This function will be send value list enrollment to outbound channel immediately
// 2. And then process store report enrollment
func (AUC *AnalyticJobUseCase) createReportEnrollment(wg *sync.WaitGroup, out chan<- []canvas.Enrollment, ctx context.Context, reportCourseID, courseID uint32) {

	enrollments, err := AUC.listEnrollment(courseID)
	if err != nil {
		AUC.Log.Error(err)
	}
	// Send out channel
	out <- enrollments
	ch := make(chan canvas.Enrollment)
	// Wait inbound channel
	go AUC.dispatchWorkerCreateReportEnrollment(ctx, wg, ch, reportCourseID)
	// send to channel
	for _, enrollment := range enrollments {
		fmt.Println(enrollment)
		wg.Add(1)
		ch <- enrollment
	}
	wg.Done()
}

// listReportEnrollment This method for get list Report Enrollment and a part of AnalyzeEnrollment
func (AUC *AnalyticJobUseCase) listReportEnrollment(ctx context.Context, filter report.ReportEnrollment) (res []report.ReportEnrollment, err error) {
	countTry := 0
	for {
		enroll, err := AUC.ReportEnrollment.FindFilter(ctx, filter)
		if err != nil {
			if countTry > MAX_RETRY {
				return res, err
			}
			countTry++
		} else {
			res = enroll
			break
		}
	}
	return res, nil
}

// ! DEPRECATED
func (AUC *AnalyticJobUseCase) AnalyzeReportEnrollment(ctx context.Context, filter report.ReportEnrollment, ch chan<- entity.ScoreEnrollment) {
	// TODO : Fix This
	_, err := AUC.listReportEnrollment(ctx, filter)
	if err != nil {
		AUC.Log.Error(err)
		panic(err)
	}
	// TODO : FIx This
	// stundentCount, finishGrading, averageGrading := AUC.CheckScoreGrade(reportEnroll)
	stundentCount, finishGrading := uint32(0), uint32(0)
	averageGrading := float32(0)
	// <<<< END HEAD
	score := entity.ScoreEnrollment{
		CourseReportID: filter.CourseReportID,
		StudentCount:   stundentCount,
		AverageGrading: averageGrading,
		FinishGrading:  finishGrading,
	}
	ch <- score
}
