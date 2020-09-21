package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abmid/icanvas-analytics/pkg/auth"
	"github.com/abmid/icanvas-analytics/pkg/setting/entity"
	uc "github.com/abmid/icanvas-analytics/pkg/setting/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	uc := uc.NewMockSettingUseCase(ctrl)

	settings := []entity.Setting{
		{ID: 1},
	}
	uc.EXPECT().FindByFilter(entity.Setting{}).Return(settings, nil)

	r := echo.New()
	g := r.Group("/v1")
	JWTToken := auth.GenerateTokenDummy(1, "super-secret")
	NewHandler("/setting", g, "super-secret", uc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/setting", nil)
	req.Header.Set("content-type", echo.MIMEApplicationJSON)
	cookieToken := http.Cookie{
		Name:     "icanvas_token",
		Value:    JWTToken,
		HttpOnly: true,
	}
	req.AddCookie(&cookieToken)
	r.ServeHTTP(w, req)

	var response []entity.Setting
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, uint32(1), response[0].ID)
}
