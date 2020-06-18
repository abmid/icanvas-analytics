/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package http

import (
	"database/sql"

	analytics_repo "github.com/abmid/icanvas-analytics/internal/analytics/repository"
	"github.com/abmid/icanvas-analytics/internal/analytics/usecase"
)

func SetupUseCase(db *sql.DB, canvasUrl, canvasAccessToken string) usecase.AnalyticsUseCase {
	repoAnalytics := analytics_repo.NewRepositoryPG(db)
	UC := usecase.NewAnalyticsUseCase(repoAnalytics)
	return UC
}
