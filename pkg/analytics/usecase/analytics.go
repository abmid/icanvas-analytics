/*
 * File Created: Thursday, 11th June 2020 1:25:01 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */
package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/pagination"
	"github.com/abmid/icanvas-analytics/pkg/analytics/entity"
	"github.com/abmid/icanvas-analytics/pkg/analytics/repository"
)

type AnalyticsUC struct {
	repoAnalytics repository.AnalyticsRepository
}

func NewAnalyticsUseCase(repoAnalytics repository.AnalyticsRepository) *AnalyticsUC {
	return &AnalyticsUC{
		repoAnalytics: repoAnalytics,
	}
}

// FindBestCourseByFilter a function to get all course or by filter from usecase, and will be use in layer 3 (Controller, etc)
func (aUC *AnalyticsUC) FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) (res []entity.AnalyticsCourse, pag pagination.Pagination, err error) {
	res, pag, err = aUC.repoAnalytics.FindBestCourseByFilter(ctx, filter)
	if err != nil {
		return nil, pag, err
	}
	return res, pag, nil
}
