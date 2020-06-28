package http

import (
	"database/sql"

	register_uc "github.com/abmid/icanvas-analytics/pkg/auth/register/usecase"
	user_repo "github.com/abmid/icanvas-analytics/pkg/user/repository"
	user_uc "github.com/abmid/icanvas-analytics/pkg/user/usecase"
)

func SetupUseCase(db *sql.DB) register_uc.RegisterUseCase {
	userRepo := user_repo.NewPG(db)
	userUC := user_uc.New(userRepo)

	registerUC := register_uc.New(userUC)

	return registerUC
}
