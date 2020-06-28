package http

import (
	"net/http"

	"github.com/abmid/icanvas-analytics/pkg/auth"
	usecase "github.com/abmid/icanvas-analytics/pkg/auth/login/usecase"
	"github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
)

type AuthHandler struct {
	loginUC usecase.LoginUseCase
	JWTKey  string
}

type FormLogin struct {
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (AH *AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var form FormLogin
		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, ResponseError{Message: "Failed request"})
		}
		err = c.Validate(form)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
		}

		user, httpStatus, err := AH.loginUC.Login(form.Email, form.Password)
		if err != nil {
			return c.JSON(httpStatus, ResponseError{Message: err.Error()})
		}

		claims := auth.JwtCustomClaims{
			int(user.ID),
			jwt.StandardClaims{
				Issuer:    "icanvas-analytics",
				ExpiresAt: 0,
			},
		}

		// Create token with claims
		createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

		// Generate encoded token and send it as response.
		token, err := createToken.SignedString([]byte(AH.JWTKey))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, ResponseSuccess{
			Email: user.Email,
			Name:  user.Name,
			Token: token,
		})
	}
}

func NewHandler(path string, g *echo.Group, jwtKey string, loginUC usecase.LoginUseCase) {
	handler := AuthHandler{
		loginUC: loginUC,
		JWTKey:  jwtKey,
	}
	r := g.Group(path)
	r.POST("/login", handler.Login())
}
