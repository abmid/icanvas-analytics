/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package http

import (
	"database/sql"

	"github.com/abmid/icanvas-analytics/internal/pagination"
	analytics_repo "github.com/abmid/icanvas-analytics/pkg/analytics/repository"
	"github.com/abmid/icanvas-analytics/pkg/analytics/usecase"
)

// SetupUseCase a function to completely requirement use case. And the return usecase ready to use.
func SetupUseCase(db *sql.DB) usecase.AnalyticsUseCase {
	pagination := pagination.New(db)
	repoAnalytics := analytics_repo.NewRepositoryPG(db, pagination)
	UC := usecase.NewAnalyticsUseCase(repoAnalytics)
	return UC
}
