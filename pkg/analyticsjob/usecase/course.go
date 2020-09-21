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

	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	report "github.com/abmid/icanvas-analytics/pkg/report/entity"
)

// listReportCourse @Deprecated
func (AUC *AnalyticJobUseCase) listReportCourse(filter report.ReportCourse, out chan<- report.ReportCourse, wg *sync.WaitGroup) {
	reportCourses, err := AUC.ReportCourse.FindFilter(context.Background(), filter)
	if err != nil {
		out <- report.ReportCourse{}
	}
	for _, reportCourse := range reportCourses {
		wg.Add(1)
		out <- reportCourse
	}
	close(out)
}

// createReportCourse a function to store report course
func (AUC *AnalyticJobUseCase) createReportCourse(course canvas.Course) (ReportCourseID uint32, err error) {
	reportCourse := report.ReportCourse{
		CourseID:   course.ID,
		AccountID:  course.AccountID,
		CourseName: course.Name,
	}
	countTry := 0
	for {
		err = AUC.ReportCourse.CreateOrUpdateCourseID(context.TODO(), &reportCourse)
		if err != nil {
			if countTry > 2 {
				AUC.Log.Error(err)
				break
			}
			countTry++
		} else {
			ReportCourseID = reportCourse.ID
			break
		}
	}
	return ReportCourseID, nil
}
