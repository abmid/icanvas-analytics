package http

import (
	"database/sql"

	analytics_repo "github.com/abmid/icanvas-analytics/internal/analytics/repository"
	"github.com/abmid/icanvas-analytics/internal/analytics/usecase"
)

func SetupUseCase(db *sql.DB, canvasUrl, canvasAccessToken string) *usecase.AnalyticsUC {
	repoAnalytics := analytics_repo.NewRepositoryPG(db)
	UC := usecase.NewAnalyticsUseCase(repoAnalytics)
	return UC
}
