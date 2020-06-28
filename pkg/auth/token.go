package auth

import (
	"encoding/json"
	"log"

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
