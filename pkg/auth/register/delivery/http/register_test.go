package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/validation"
	register_usecase "github.com/abmid/icanvas-analytics/pkg/auth/register/usecase"
	"github.com/abmid/icanvas-analytics/pkg/user/entity"
	user_uc_mock "github.com/abmid/icanvas-analytics/pkg/user/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	userUC := user_uc_mock.NewMockUserUseCase(ctrl)
	userUC.EXPECT().Create(gomock.Any()).DoAndReturn(func(r *entity.User) error {
		r.ID = 1
		return nil
	})

	registerUC := register_usecase.New(userUC)

	e := echo.New()
	validation.AlphaValidation(e)
	gr := e.Group("/v1")

	NewHandler("/auth", gr, registerUC)
	w := httptest.NewRecorder()
	f := make(url.Values)
	f.Set("name", "test")
	f.Set("email", "test@test.com")
	f.Set("password", "pass")
	req, _ := http.NewRequest("POST", "/v1/auth/register", strings.NewReader(f.Encode()))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationForm)
	e.ServeHTTP(w, req)

	var result ResponseSuccess
	json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, result.ID, uint32(1))
}
