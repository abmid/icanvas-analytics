package http

import (
	"database/sql"

	login_uc "github.com/abmid/icanvas-analytics/pkg/auth/login/usecase"
	user_repo "github.com/abmid/icanvas-analytics/pkg/user/repository"
	user_uc "github.com/abmid/icanvas-analytics/pkg/user/usecase"
)

func SetupUseCase(db *sql.DB) login_uc.LoginUseCase {

	userRepo := user_repo.NewPG(db)
	userUC := user_uc.New(userRepo)

	loginUC := login_uc.New(userUC)

	return loginUC
}
