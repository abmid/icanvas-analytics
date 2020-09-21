package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/validation"
	mock_login_uc "github.com/abmid/icanvas-analytics/pkg/auth/login/usecase/mock"
	user "github.com/abmid/icanvas-analytics/pkg/user/entity"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("unauthorized", func(t *testing.T) {

		mockLoginUC := mock_login_uc.NewMockLoginUseCase(ctrl)
		mockLoginUC.EXPECT().Login("test@test.com", "pass").Return(nil, http.StatusUnauthorized, errors.New("Test"))

		e := echo.New()
		validation.AlphaValidation(e)
		v1 := e.Group("/v1")

		NewHandler("/auth", v1, "super-secret", mockLoginUC)

		w := httptest.NewRecorder()
		url := make(url.Values)
		url.Set("email", "test@test.com")
		url.Set("password", "pass")
		req, _ := http.NewRequest(echo.POST, "/v1/auth/login", strings.NewReader(url.Encode()))
		req.Header.Set("Content-Type", echo.MIMEApplicationForm)
		e.ServeHTTP(w, req)

		var result map[string]interface{}
		json.NewDecoder(w.Body).Decode(&result)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("success", func(t *testing.T) {

		user := user.User{
			ID:    1,
			Email: "test@test.com",
		}
		mockLoginUC := mock_login_uc.NewMockLoginUseCase(ctrl)
		mockLoginUC.EXPECT().Login("test@test.com", "pass").Return(&user, http.StatusOK, nil)

		e := echo.New()
		validation.AlphaValidation(e)
		v1 := e.Group("/v1")

		NewHandler("/auth", v1, "super-secret", mockLoginUC)

		w := httptest.NewRecorder()
		url := make(url.Values)
		url.Set("email", "test@test.com")
		url.Set("password", "pass")
		req, _ := http.NewRequest(echo.POST, "/v1/auth/login", strings.NewReader(url.Encode()))
		req.Header.Set("Content-Type", echo.MIMEApplicationForm)
		e.ServeHTTP(w, req)

		var result map[string]interface{}
		json.NewDecoder(w.Body).Decode(&result)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
