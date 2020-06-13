package job

import (
	"database/sql"
	"fmt"
	"net/http"

	analyticsjob "github.com/abmid/icanvas-analytics/internal/analyticsjob/usecase"
	report_repo_assigment "github.com/abmid/icanvas-analytics/internal/report/assigment/repository"
	report_uc_assigment "github.com/abmid/icanvas-analytics/internal/report/assigment/usecase"
	report_repo_course "github.com/abmid/icanvas-analytics/internal/report/course/repository"
	report_uc_course "github.com/abmid/icanvas-analytics/internal/report/course/usecase"
	report_repo_discussion "github.com/abmid/icanvas-analytics/internal/report/discussion/repository"
	report_uc_discussion "github.com/abmid/icanvas-analytics/internal/report/discussion/usecase"
	report_repo_enroll "github.com/abmid/icanvas-analytics/internal/report/enrollment/repository"
	report_uc_enroll "github.com/abmid/icanvas-analytics/internal/report/enrollment/usecase"
	report_repo_result "github.com/abmid/icanvas-analytics/internal/report/result/repository"
	report_uc_result "github.com/abmid/icanvas-analytics/internal/report/result/usecase"
	canvas_repo_assigment "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/repository"
	canvas_uc_assigment "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/usecase"
	canvas_repo_course "github.com/abmid/icanvas-analytics/pkg/canvas/course/repository"
	canvas_uc_course "github.com/abmid/icanvas-analytics/pkg/canvas/course/usecase"
	canvas_repo_discussion "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/repository"
	canvas_uc_discussion "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/usecase"
	canvas_repo_enroll "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/repository"
	canvas_uc_enroll "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/usecase"

	"github.com/robfig/cron/v3"
)

func RunScheduling(c *cron.Cron, db *sql.DB, canvasUrl, canvasAccessToken string) {
	_ = JobUseCase(db, canvasUrl, canvasAccessToken)
	status := true
	// useCase.RunJob(uint32(1))
	c.AddFunc("@every 0h0m1s", func() {
		if status {
			fmt.Println("Running Run Job")
			// useCase.RunJob(uint32(39))
			status = false
		}

	})
}

func JobUseCase(db *sql.DB, canvasUrl, canvasAccessToken string) *analyticsjob.AnalyticJobUseCase {
	client := http.DefaultClient
	canvasRepoAssigment := canvas_repo_assigment.NewRepositoryAPI(client, canvasUrl, canvasAccessToken)
	canvasRepoCourse := canvas_repo_course.NewRepositoryAPI(client, canvasUrl, canvasAccessToken)
	canvasRepoDiscussion := canvas_repo_discussion.NewRepositoryAPI(client, canvasUrl, canvasAccessToken)
	canvasRepoEnroll := canvas_repo_enroll.NewRepositoryAPI(client, canvasUrl, canvasAccessToken)

	canvasUCAssigment := canvas_uc_assigment.NewAssigmentUseCase(canvasRepoAssigment)
	canvasUCCourse := canvas_uc_course.NewCourseUseCase(canvasRepoCourse)
	canvasUCDiscussion := canvas_uc_discussion.NewDiscussUseCase(canvasRepoDiscussion)
	canvasUCEnroll := canvas_uc_enroll.NewEnrollUseCase(canvasRepoEnroll)

	reportRepoAssigment := report_repo_assigment.NewAssigmentPG(db)
	reportRepoCourse := report_repo_course.NewCoursePG(db)
	reportRepoDiscussion := report_repo_discussion.NewDiscussionPG(db)
	reportRepoEnrollment := report_repo_enroll.NewEnrollmentPG(db)
	reportRepoResult := report_repo_result.NewResultPG(db)

	reportUCAssigment := report_uc_assigment.NewReportAssigmentUseCase(reportRepoAssigment)
	reportUCCourse := report_uc_course.NewReportCourseUseCase(reportRepoCourse)
	reportUCDiscussion := report_uc_discussion.NewReportDiscussionUseCase(reportRepoDiscussion)
	reportUCEnrollment := report_uc_enroll.NewReportEnrollUseCase(reportRepoEnrollment)
	reportUCResult := report_uc_result.NewReportResultUseCase(reportRepoResult)

	UC := analyticsjob.NewAnalyticJobUseCase(
		canvasUCCourse,
		canvasUCAssigment,
		canvasUCEnroll,
		canvasUCDiscussion,
		reportUCAssigment,
		reportUCCourse,
		reportUCDiscussion,
		reportUCEnrollment,
		reportUCResult)

	return UC
}
