/*
 * File Created: Thursday, 11th June 2020 1:25:01 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */
package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
	"github.com/abmid/icanvas-analytics/internal/analytics/repository"
)

type AnalyticsUC struct {
	repoAnalytics repository.AnalyticsRepository
}

func NewAnalyticsUseCase(repoAnalytics repository.AnalyticsRepository) *AnalyticsUC {
	return &AnalyticsUC{
		repoAnalytics: repoAnalytics,
	}
}

func (aUC *AnalyticsUC) FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) (res []entity.AnalyticsCourse, err error) {
	res, err = aUC.repoAnalytics.FindBestCourseByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
