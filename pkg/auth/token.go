/*
 * File Created: Saturday, 27th June 2020 12:46:08 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func HandleToken(c echo.Context) (JwtCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)
	claims := JwtCustomClaims{}
	tmp, err := json.Marshal(user.Claims)
	if err != nil {
		return claims, err
	}
	err = json.Unmarshal(tmp, &claims)
	if err != nil {
		return claims, err
	}

	return claims, nil
}

// WriteTokenCookie is function to write cookie flag httpOnly from server
// When it use for login, the fill for payload is token and when use for logout the fill for payload is logout
func WriteTokenCookie(c echo.Context, payload string) error {
	cookie := new(http.Cookie)
	if payload == "logout" {
		cookie.Name = "icanvas_token"
		cookie.Value = "logout"
		cookie.Expires = time.Unix(0, 0)
		cookie.MaxAge = -1
		cookie.HttpOnly = true
		cookie.Path = "/"
	} else {
		cookie.Name = "icanvas_token"
		cookie.Value = payload
		cookie.Expires = time.Now().Add(8760 * time.Hour)
		cookie.HttpOnly = true
		cookie.Path = "/"
	}
	c.SetCookie(cookie)
	return nil
}

func GenerateTokenDummy(userID int, jwtKey string) string {
	claims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    "icanvas-analytics",
			ExpiresAt: 0,
		},
	}

	// Create token with claims
	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	// Generate encoded token and send it as response.
	token, err := createToken.SignedString([]byte(jwtKey))
	if err != nil {
		log.Fatal(err)
	}
	return token
}
