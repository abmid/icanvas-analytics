package http

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	register_usecase "github.com/abmid/icanvas-analytics/internal/auth/register/usecase"
	"github.com/abmid/icanvas-analytics/internal/user/entity"
	user_uc_mock "github.com/abmid/icanvas-analytics/internal/user/usecase/mock"
	"github.com/abmid/icanvas-analytics/internal/validation"
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

	New("/auth", gr, registerUC)
	w := httptest.NewRecorder()
	f := make(url.Values)
	f.Set("name", "test")
	f.Set("email", "test@test.com")
	f.Set("password", "pass")
	req, _ := http.NewRequest("POST", "/v1/auth/register", strings.NewReader(f.Encode()))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationForm)
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
