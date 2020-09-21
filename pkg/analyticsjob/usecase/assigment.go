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

	"github.com/abmid/icanvas-analytics/pkg/analyticsjob/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	report "github.com/abmid/icanvas-analytics/pkg/report/entity"
)

var (
	MAX_RETRY       = 3
	WORKER_DATABASE = 10
)

// listAssigment This method to get List Assingment From Repository Canvas
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
// dispatchWorkerCreateReportAssigment a function running go routine such as many worker database. This function wait value from upstream via inbound channel
func (AUC *AnalyticJobUseCase) dispatchWorkerCreateReportAssigment(ctx context.Context, wg *sync.WaitGroup, in <-chan canvas.Assigment, reportCourseID uint32) {
	for i := 0; i < WORKER_DATABASE; i++ {
		// Run go routine of WORKER_DATABASE, and all goroute wait value from upstream via inbound channel
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
					// Store report assigment
					err := AUC.ReportAssigment.CreateOrUpdateByFilter(ctx, filter, &reportAss)
					if err != nil {
						// will be try 3 times if failed
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

// createReportAssigment a function to create report assigment, in this function will have 2 operation after get list assigment.
// 1. This function will be send value listAssigment to outbound channel immediately
// 2. And then process store report assigment
func (AUC *AnalyticJobUseCase) createReportAssigment(wg *sync.WaitGroup, out chan<- []canvas.Assigment, ctx context.Context, reportCourseID, courseID uint32) {
	assigments, err := AUC.listAssigment(courseID)
	if err != nil {
		AUC.Log.Error(err)
	}
	// TODO : Send assigments to out channel
	// Operation 1
	out <- assigments
	ch := make(chan canvas.Assigment)
	// Running worker and wait value from upstream inbound channel
	go AUC.dispatchWorkerCreateReportAssigment(ctx, wg, ch, reportCourseID)
	// Send assigment to channel
	for _, assigment := range assigments {
		wg.Add(1)
		ch <- assigment
	}
	close(ch)
	wg.Done()
}

// listReportAssigment get list report assigment with try 3 time if failed
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
		AUC.Log.Error(err)
	}
	score := entity.ScoreAssigment{
		CourseReportID: filter.CourseReportID,
		AssigmentCount: uint32(len(reportAssigments)),
	}
	ch <- score
}
