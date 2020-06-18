/*
 * File Created: Thursday, 18th June 2020 5:32:30 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */
package http

import (
	"database/sql"

	register_usecase "github.com/abmid/icanvas-analytics/internal/auth/register/usecase"
	user_repository "github.com/abmid/icanvas-analytics/internal/user/repository"
	user_usecase "github.com/abmid/icanvas-analytics/internal/user/usecase"
)

func SetupUseCase(db *sql.DB) register_usecase.RegisterUseCase {
	userRepo := user_repository.NewPG(db)
	userUC := user_usecase.New(userRepo)

	registerUC := register_usecase.New(userUC)

	return registerUC
}
