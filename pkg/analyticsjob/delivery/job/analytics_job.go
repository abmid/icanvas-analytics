/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package job

import (
	"database/sql"
	"fmt"
	"net/http"

	analyticsjob "github.com/abmid/icanvas-analytics/pkg/analyticsjob/usecase"
	canvas_repo_assigment "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/repository"
	canvas_uc_assigment "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/usecase"
	canvas_repo_course "github.com/abmid/icanvas-analytics/pkg/canvas/course/repository"
	canvas_uc_course "github.com/abmid/icanvas-analytics/pkg/canvas/course/usecase"
	canvas_repo_discussion "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/repository"
	canvas_uc_discussion "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/usecase"
	canvas_repo_enroll "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/repository"
	canvas_uc_enroll "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/usecase"
	report_repo_assigment "github.com/abmid/icanvas-analytics/pkg/report/assigment/repository"
	report_uc_assigment "github.com/abmid/icanvas-analytics/pkg/report/assigment/usecase"
	report_repo_course "github.com/abmid/icanvas-analytics/pkg/report/course/repository"
	report_uc_course "github.com/abmid/icanvas-analytics/pkg/report/course/usecase"
	report_repo_discussion "github.com/abmid/icanvas-analytics/pkg/report/discussion/repository"
	report_uc_discussion "github.com/abmid/icanvas-analytics/pkg/report/discussion/usecase"
	report_repo_enroll "github.com/abmid/icanvas-analytics/pkg/report/enrollment/repository"
	report_uc_enroll "github.com/abmid/icanvas-analytics/pkg/report/enrollment/usecase"
	report_repo_result "github.com/abmid/icanvas-analytics/pkg/report/result/repository"
	report_uc_result "github.com/abmid/icanvas-analytics/pkg/report/result/usecase"
	setting_uc "github.com/abmid/icanvas-analytics/pkg/setting/usecase"

	"github.com/robfig/cron/v3"
)

// RunScheduling
func RunScheduling(c *cron.Cron, db *sql.DB, settingUC setting_uc.SettingUseCase) {
	// useCase := JobUseCase(db, settingUC)
	status := true
	// useCase.RunJob(uint32(1))
	c.AddFunc("@every 0h1m1s", func() {
		if status {
			fmt.Println("Running Run Job")
			// useCase.RunJob(uint32(39))
		}

	})
}

// JobUseCase a function to make simple completely requirement analytics job for layer 3 in clean architecture
func JobUseCase(db *sql.DB, settingUC setting_uc.SettingUseCase) *analyticsjob.AnalyticJobUseCase {
	client := http.DefaultClient
	canvasRepoAssigment := canvas_repo_assigment.NewRepositoryAPI(client, settingUC)
	canvasRepoCourse := canvas_repo_course.NewRepositoryAPI(client, settingUC)
	canvasRepoDiscussion := canvas_repo_discussion.NewRepositoryAPI(client, settingUC)
	canvasRepoEnroll := canvas_repo_enroll.NewRepositoryAPI(client, settingUC)

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
