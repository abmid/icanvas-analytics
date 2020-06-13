package usecase

import (
	"context"
	"sync"

	"github.com/abmid/icanvas-analytics/internal/analyticsjob/entity"
	report "github.com/abmid/icanvas-analytics/internal/report/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/sirupsen/logrus"
)

func (AUC *AnalyticJobUseCase) listDiscussion(courseID uint32) (res []canvas.Discussion, err error) {
	countTry := 0
	for {
		diss, err := AUC.Discussion.ListDiscussionByCourseID(courseID)
		if err != nil {
			if countTry > MAX_RETRY {
				break
			}
			countTry++
		} else {
			res = diss
			break
		}

	}
	return res, nil
}

func (AUC *AnalyticJobUseCase) dispatchWorkerCreateReportDiscussion(ctx context.Context, wg *sync.WaitGroup, in <-chan canvas.Discussion, reportCourseID uint32) {
	for i := 0; i < WORKER_DATABASE; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup, in <-chan canvas.Discussion, reportCourseID uint32) {
			for discussion := range in {
				reportDiss := report.ReportDiscussion{
					CourseReportID: reportCourseID,
					Title:          discussion.Title,
					DiscussionID:   discussion.ID,
				}
				filter := reportDiss
				countTry := 0
				for {
					err := AUC.ReportDiscussion.CreateOrUpdateByFilter(ctx, filter, &reportDiss)
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

func (AUC *AnalyticJobUseCase) createReportDiscussion(wg *sync.WaitGroup, out chan<- []canvas.Discussion, ctx context.Context, reportCourseID, courseID uint32) {
	discussions, err := AUC.Discussion.ListDiscussionByCourseID(courseID)
	if err != nil {
		logrus.Error(err)
	}
	out <- discussions
	ch := make(chan canvas.Discussion)
	go AUC.dispatchWorkerCreateReportDiscussion(ctx, wg, ch, reportCourseID)
	for _, discussion := range discussions {
		wg.Add(1)
		ch <- discussion
	}
	close(ch)
	wg.Done()
}

/*
This method for get list Report Discussion and a part of AnalyzeDiscussion
*/
func (AUC *AnalyticJobUseCase) listReportDiscussion(ctx context.Context, filter report.ReportDiscussion) (res []report.ReportDiscussion, err error) {
	countTry := 0
	for {
		ass, err := AUC.ReportDiscussion.FindFilter(ctx, filter)
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

// ! DEPRECATED
func (AUC *AnalyticJobUseCase) AnalyzeReportDiscussion(ctx context.Context, filter report.ReportDiscussion, ch chan<- entity.ScoreDiscussion) {

	reportDiscussion, err := AUC.listReportDiscussion(ctx, filter)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	score := entity.ScoreDiscussion{
		CourseReportID:  filter.CourseReportID,
		DiscussionCount: uint32(len(reportDiscussion)),
	}
	ch <- score
}
