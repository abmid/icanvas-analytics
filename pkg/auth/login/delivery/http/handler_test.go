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
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
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
	t.Log(result)
	t.Fatal("")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, result["status"].(string), "success")
}
