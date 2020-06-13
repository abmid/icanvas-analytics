package usecase

import (
	"context"
	"sync"

	"github.com/abmid/icanvas-analytics/internal/analyticsjob/entity"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/sirupsen/logrus"
)

var (
	MAX_RETRY       = 3
	WORKER_DATABASE = 10
)

/*
This method to get List Assingment From Repository Canvas
@param courseID
@return []entity.Assigment
@return err
*/
func (AUC *AnalyticJobUseCase) listAssigment(courseID uint32) (res []canvas.Assigment, err error) {
	countTry := 0
	for {
		ass, err := AUC.Assigment.ListAssigmentByCourseID(courseID)
		if err != nil {
			if countTry > MAX_RETRY {
				break
			}
			countTry++
		} else {
			res = ass
			break
		}
	}
	return res, nil
}

// TODO : CREATE WORKER FOR Insert Database
// TODO : 1. Buat Channel Untuk Nampung Assigment
// TODO : 2. Init Worker Sebelum perulangan
// TODO : 3. Didalam Worker kirim data assigment ke channel
// ? Worker akan bekerja setelah mendapat data kiriman dari looping
func (AUC *AnalyticJobUseCase) dispatchWorkerCreateReportAssigment(ctx context.Context, wg *sync.WaitGroup, in <-chan canvas.Assigment, reportCourseID uint32) {
	for i := 0; i < WORKER_DATABASE; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup, in <-chan canvas.Assigment, reportCourseID uint32) {
			for assigment := range in {
				reportAss := report.ReportAssigment{
					AssigmentID:    assigment.ID,
					CourseReportID: reportCourseID,
					Name:           assigment.Name,
				}
				filter := reportAss
				countTry := 0
				for {
					err := AUC.ReportAssigment.CreateOrUpdateByFilter(ctx, filter, &reportAss)
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

/*
This method for store list assigment into database
*/
func (AUC *AnalyticJobUseCase) createReportAssigment(wg *sync.WaitGroup, out chan<- []canvas.Assigment, ctx context.Context, reportCourseID, courseID uint32) {
	assigments, err := AUC.listAssigment(courseID)
	if err != nil {
		logrus.Error(err)
	}
	// TODO : Send assigments to out channel
	out <- assigments
	ch := make(chan canvas.Assigment)
	// Fan-in, fan-out pattern
	go AUC.dispatchWorkerCreateReportAssigment(ctx, wg, ch, reportCourseID)
	for _, assigment := range assigments {
		wg.Add(1)
		ch <- assigment
	}
	close(ch)
	wg.Done()
}

/*
Get Report Assigment with handle Retry, this method a part of AnalyzeReportAssigment
*/
func (AUC *AnalyticJobUseCase) listReportAssigment(ctx context.Context, filter report.ReportAssigment) (res []report.ReportAssigment, err error) {
	countTry := 0
	for {
		ass, err := AUC.ReportAssigment.FindFilter(ctx, filter)
		if err != nil {
			if countTry > MAX_RETRY {
				return res, err
			}
			countTry++
		} else {
			res = ass
			break
		}
	}
	return res, nil
}

func (AUC *AnalyticJobUseCase) AnalyzeReportAssigment(ctx context.Context, filter report.ReportAssigment, ch chan<- entity.ScoreAssigment) {

	reportAssigments, err := AUC.listReportAssigment(ctx, filter)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	score := entity.ScoreAssigment{
		CourseReportID: filter.CourseReportID,
		AssigmentCount: uint32(len(reportAssigments)),
	}
	ch <- score
}
