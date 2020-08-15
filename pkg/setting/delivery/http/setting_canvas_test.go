package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/validation"
	"github.com/abmid/icanvas-analytics/pkg/auth"
	uc "github.com/abmid/icanvas-analytics/pkg/setting/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestCreateOrUpdateCanvas(t *testing.T) {
	ctrl := gomock.NewController(t)
	uc := uc.NewMockSettingUseCase(ctrl)

	uc.EXPECT().CreateAll(gomock.Any()).Return(nil)

	r := echo.New()
	validation.AlphaValidation(r)
	g := r.Group("/v1")
	JWTToken := auth.GenerateTokenDummy(1, "super-secret")
	NewHandler("/setting", g, "super-secret", uc)

	w := httptest.NewRecorder()
	formData := make(url.Values)
	formData.Set("canvas_url", "url")
	formData.Set("canvas_token", "token")
	req, _ := http.NewRequest("POST", "/v1/setting/canvas", strings.NewReader(formData.Encode()))
	req.Header.Set("content-type", echo.MIMEApplicationForm)
	cookieToken := http.Cookie{
		Name:     "icanvas_token",
		Value:    JWTToken,
		HttpOnly: true,
	}
	req.AddCookie(&cookieToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestExistsCanvasConfiguraton(t *testing.T) {
	ctrl := gomock.NewController(t)
	uc := uc.NewMockSettingUseCase(ctrl)
	uc.EXPECT().ExistsCanvasConfig().Return(true, "url", "token", nil)

	r := echo.New()
	validation.AlphaValidation(r)
	g := r.Group("/v1")
	JWTToken := auth.GenerateTokenDummy(1, "super-secret")
	NewHandler("/setting", g, "super-secret", uc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/setting/canvas", nil)
	cookieToken := http.Cookie{
		Name:     "icanvas_token",
		Value:    JWTToken,
		HttpOnly: true,
	}
	req.AddCookie(&cookieToken)
	r.ServeHTTP(w, req)

	var res map[string]interface{}
	json.NewDecoder(w.Body).Decode(&res)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res["token"].(string), "token")
	assert.Equal(t, res["url"].(string), "url")
}
