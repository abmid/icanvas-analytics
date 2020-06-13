package usecase

import (
	"context"
	"fmt"
	"sync"

	report "github.com/abmid/icanvas-analytics/internal/report/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

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

func (AUC *AnalyticJobUseCase) createReportCourse(course canvas.Course) (ReportCourseID uint32, err error) {
	reportCourse := report.ReportCourse{
		CourseID:   course.ID,
		AccountID:  course.AccountID,
		CourseName: course.Name,
	}
	countTry := 0
	for {
		err = AUC.ReportCourse.CreateOrUpdateCourseID(context.TODO(), &reportCourse)
		fmt.Println("ERROR CREATE REPORT COURSE ", err)
		if err != nil {
			if countTry > 2 {
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
