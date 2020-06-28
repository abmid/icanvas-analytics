/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/pkg/report/discussion/repository"
	"github.com/abmid/icanvas-analytics/pkg/report/entity"
	report "github.com/abmid/icanvas-analytics/pkg/report/entity"
)

type reportDiscussionUseCase struct {
	RepoReportDiscussion repository.DisscussionRepositoryPG
}

func NewReportDiscussionUseCase(repoReportDiscussion repository.DisscussionRepositoryPG) *reportDiscussionUseCase {
	return &reportDiscussionUseCase{
		RepoReportDiscussion: repoReportDiscussion,
	}
}

func (useCase *reportDiscussionUseCase) Create(ctx context.Context, reportDiss *report.ReportDiscussion) error {
	err := useCase.RepoReportDiscussion.Create(ctx, reportDiss)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *reportDiscussionUseCase) Read(ctx context.Context) ([]report.ReportDiscussion, error) {
	results, err := useCase.RepoReportDiscussion.Read(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (useCase *reportDiscussionUseCase) Update(ctx context.Context, reportDiss *report.ReportDiscussion) error {
	err := useCase.RepoReportDiscussion.Update(ctx, reportDiss)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportDiscussionUseCase) Delete(ctx context.Context, reportDissID uint32) error {
	err := useCase.RepoReportDiscussion.Delete(ctx, reportDissID)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *reportDiscussionUseCase) CreateOrUpdateByFilter(ctx context.Context, filter report.ReportDiscussion, reportDiss *report.ReportDiscussion) error {
	findReportDiss, err := useCase.RepoReportDiscussion.FindFirstByFilter(ctx, filter)
	if err != nil {
		return err
	}
	reportDiss.ID = findReportDiss.ID
	if findReportDiss.ID == 0 {
		err := useCase.RepoReportDiscussion.Create(ctx, reportDiss)
		if err != nil {
			return err
		}
	} else {
		err := useCase.RepoReportDiscussion.Update(ctx, reportDiss)
		if err != nil {
			return err
		}
	}
	return nil
}

func (useCase *reportDiscussionUseCase) FindFilter(ctx context.Context, filter entity.ReportDiscussion) ([]entity.ReportDiscussion, error) {
	results, err := useCase.RepoReportDiscussion.FindFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return results, nil
}
