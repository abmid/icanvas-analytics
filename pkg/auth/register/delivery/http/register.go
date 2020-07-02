/*
 * File Created: Thursday, 18th June 2020 5:31:14 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */
package http

import (
	"net/http"

	"github.com/abmid/icanvas-analytics/pkg/user/entity"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type DataForm struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	ID uint32 `json:"id"`
}

type ResponseRegisterCheck struct {
	Status bool `json:"status"`
}

func hashAndSalt(password string) (string, error) {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *RegisterHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		dataForm := new(DataForm)
		/**
		* Bind Data
		 */
		if err := c.Bind(dataForm); err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}

		/**
		* Validate Form
		 */
		if err := c.Validate(dataForm); err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}

		/**
		* Encrypt password
		 */
		encryptPassword, err := hashAndSalt(dataForm.Password)
		if err != nil {
			return c.JSON(http.StatusConflict, ResponseError{Message: "Failed encrypt password"})
		}

		/**
		* Store user to database
		 */
		user := entity.User{
			Email:    dataForm.Email,
			Name:     dataForm.Name,
			Password: encryptPassword,
		}
		err = h.registerUC.Register(&user)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
		}

		return c.JSON(http.StatusCreated, ResponseSuccess{ID: user.ID})
	}
}

func (RH *RegisterHandler) RegisterCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		isRegister, err := RH.registerUC.RegisterCheck()
		if err != nil {
			return c.JSON(http.StatusConflict, ResponseError{Message: err.Error()})
		}
		if isRegister {
			return c.JSON(http.StatusOK, ResponseRegisterCheck{Status: true})
		}
		return c.JSON(http.StatusOK, ResponseRegisterCheck{Status: false})
	}
}
