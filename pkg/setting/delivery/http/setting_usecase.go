package http

import (
	"database/sql"

	"github.com/abmid/icanvas-analytics/pkg/setting/repository"
	"github.com/abmid/icanvas-analytics/pkg/setting/usecase"
)

func SetupUseCase(db *sql.DB) usecase.SettingUseCase {
	repo := repository.NewRepositoryPG(db)
	uc := usecase.NewSettingUseCase(repo)

	return uc
}
