/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/pagination"
	"github.com/abmid/icanvas-analytics/pkg/analytics/entity"
)

type AnalyticsRepository interface {
	FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) ([]entity.AnalyticsCourse, pagination.Pagination, error)
}
