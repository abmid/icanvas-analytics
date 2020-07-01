package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/validation"
	echo "github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestLogout(t *testing.T) {
	e := echo.New()
	validation.AlphaValidation(e)
	v1 := e.Group("/v1")

	NewHandler("/auth", v1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(echo.POST, "/v1/auth/logout", nil)
	req.Header.Set("Content-Type", echo.MIMEApplicationForm)
	e.ServeHTTP(w, req)

	var result map[string]interface{}
	setCookie := w.Header().Get(echo.HeaderSetCookie)
	splitSetCookie := strings.Split(setCookie, ";")
	json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, result["status"].(string), "OK")
	assert.Equal(t, strings.Split(splitSetCookie[0], "=")[1], "logout")
}
