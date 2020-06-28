package auth

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
