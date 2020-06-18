/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/course/repository"
	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type reportCourseUseCase struct {
	RepoReportCourse repository.CourseRepositoryPG
}

func NewReportCourseUseCase(repoReportCourse repository.CourseRepositoryPG) *reportCourseUseCase {
	return &reportCourseUseCase{
		RepoReportCourse: repoReportCourse,
	}
}

func (useCase *reportCourseUseCase) Create(ctx context.Context, reportCourse *entity.ReportCourse) error {
	err := useCase.RepoReportCourse.Create(ctx, reportCourse)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportCourseUseCase) Read(ctx context.Context) ([]entity.ReportCourse, error) {
	results, err := useCase.RepoReportCourse.Read(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (useCase *reportCourseUseCase) FindFilter(ctx context.Context, filter entity.ReportCourse) ([]entity.ReportCourse, error) {
	results, err := useCase.RepoReportCourse.FindFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (useCase *reportCourseUseCase) Update(ctx context.Context, reportCourse *entity.ReportCourse) error {
	err := useCase.RepoReportCourse.Update(ctx, reportCourse)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportCourseUseCase) Delete(ctx context.Context, reportCourseID uint32) error {
	err := useCase.RepoReportCourse.Delete(ctx, reportCourseID)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportCourseUseCase) CreateOrUpdateCourseID(ctx context.Context, reportCourse *entity.ReportCourse) error {
	// Find report course By Course ID and Date Now are exist
	resReport, err := useCase.RepoReportCourse.FindByCourseIDDateNow(ctx, reportCourse.CourseID)
	if err != nil {
		return err
	}
	// If Not Exist
	if resReport.ID == 0 {
		err := useCase.RepoReportCourse.Create(ctx, reportCourse)
		if err != nil {
			return err
		}
		return nil
	}
	reportCourse.ID = resReport.ID
	return nil
}
