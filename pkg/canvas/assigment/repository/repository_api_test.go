package repository

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock_assigment "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/repository/mock/canvas"
	mock_setting "github.com/abmid/icanvas-analytics/pkg/setting/usecase/mock"
	"github.com/golang/mock/gomock"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestListAssigmentByCourseID(t *testing.T) {
	srv := serverMock()
	ctrl := gomock.NewController(t)
	defer srv.Close()
	// Mock Setting UC
	settingUC := mock_setting.NewMockSettingUseCase(ctrl)
	settingUC.EXPECT().ExistsCanvasConfig().Return(true, srv.URL, "my-secret-token", nil)

	AssigmentRepo := NewRepositoryAPI(http.DefaultClient, settingUC)
	res, err := AssigmentRepo.ListAssigmentByCourseID(1)
	if err != nil {
		t.Error(err)
	}
	assert.NilError(t, err, "Error List Assigment By Course ID")
	assert.Equal(t, len(res), 1, "Result not match")
}

func serverMock() *httptest.Server {
	myGin := ginMock()
	srv := httptest.NewServer(myGin)

	return srv
}

func ginMock() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/api/v1/courses/:courseID/assignments", mock_assigment.ListAssigmentByCourseID())

	return r
}
