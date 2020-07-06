package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/validation"
	"github.com/abmid/icanvas-analytics/pkg/auth"
	account_uc "github.com/abmid/icanvas-analytics/pkg/canvas/account/usecase/mock"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestListAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	// Mock Account
	listAccount := []entity.Account{
		{ID: 1},
	}
	accountUC := account_uc.NewMockAccountUseCase(ctrl)
	accountUC.EXPECT().ListAccount(uint32(1)).Return(listAccount, nil)
	// Init Echo
	g := echo.New()
	validation.AlphaValidation(g)
	// Set Routing and Handler
	gr := g.Group("/v1")
	NewHandler("/canvas", gr, "super-secret", accountUC)
	// Generate Token
	token := auth.GenerateTokenDummy(1, "super-secret")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/canvas/accounts/1/sub_accounts", nil)
	// req.Header.Add(echo.HeaderAuthorization, "Bearer "+token)
	cookieToken := http.Cookie{
		Name:     "icanvas_token",
		Value:    token,
		HttpOnly: true,
	}
	req.AddCookie(&cookieToken)
	g.ServeHTTP(w, req)

	var result []entity.Account
	json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(result), len(listAccount))
}
