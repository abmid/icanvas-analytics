package repository

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock_account "github.com/abmid/icanvas-analytics/pkg/canvas/account/repository/mock/canvas"
	mock_setting "github.com/abmid/icanvas-analytics/pkg/setting/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestListAccount(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	ctrl := gomock.NewController(t)
	// Mock Setting UC
	settingUC := mock_setting.NewMockSettingUseCase(ctrl)
	settingUC.EXPECT().ExistsCanvasConfig().Return(true, srv.URL, "my-secret-token", nil)

	AccountRepo := NewRepositoryAPI(http.DefaultClient, settingUC)
	res, err := AccountRepo.ListAccount(uint32(1))
	assert.NilError(t, err, "#Error List Account")
	assert.Equal(t, len(res), 2, "Result not match")
}

func serverMock() *httptest.Server {
	myEcho := echoMock()
	srv := httptest.NewServer(myEcho)

	return srv
}

func echoMock() *echo.Echo {
	r := echo.New()
	r.GET("/api/v1/accounts/:id/sub_accounts", mock_account.ListAccount())

	return r
}
