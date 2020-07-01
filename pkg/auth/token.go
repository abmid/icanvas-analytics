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

func WriteTokenCookie(c echo.Context, token string) error {
	cookie := new(http.Cookie)
	cookie.Name = "icanvas_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(8760 * time.Hour)
	cookie.HttpOnly = true
	cookie.Path = "/"
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
