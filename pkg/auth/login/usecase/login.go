/*
 * File Created: Saturday, 27th June 2020 11:33:49 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/user/entity"
	"github.com/abmid/icanvas-analytics/pkg/user/usecase"
	"golang.org/x/crypto/bcrypt"
)

type loginUseCase struct {
	userUC usecase.UserUseCase
	Log    *logger.LoggerWrap
}

func New(userUC usecase.UserUseCase) *loginUseCase {
	logger := logger.New()

	return &loginUseCase{
		userUC: userUC,
		Log:    logger,
	}
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice

	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	logger := logger.New()
	if err != nil {
		logger.Error(err)
		return false
	}

	return true
}

func (UC *loginUseCase) Login(email, password string) (res *entity.User, httpStatus int, err error) {

	user, err := UC.userUC.Find(email)
	if err != nil {
		UC.Log.Error(err)
		if err == sql.ErrNoRows {
			return nil, http.StatusUnauthorized, errors.New("User not found")
		}
		return nil, http.StatusUnauthorized, errors.New("Failed get resource")
	}

	if user == nil {
		return nil, http.StatusUnauthorized, errors.New("User not found")
	}

	checkPassword := comparePasswords(user.Password, password)
	if checkPassword {
		return user, http.StatusOK, nil
	}
	return nil, http.StatusUnauthorized, errors.New("Wrong password !")
}
