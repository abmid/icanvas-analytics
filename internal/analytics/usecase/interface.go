/*
 * File Created: Thursday, 11th June 2020 1:24:56 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/analytics/entity"
)

type AnalyticsUseCase interface {
	FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) ([]entity.AnalyticsCourse, error)
}
