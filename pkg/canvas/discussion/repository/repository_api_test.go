package repository

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock_discussion "github.com/abmid/icanvas-analytics/pkg/canvas/discussion/repository/mock/canvas"
	mock_setting "github.com/abmid/icanvas-analytics/pkg/setting/usecase/mock"
	"github.com/golang/mock/gomock"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestListDiscussionByCourseID(t *testing.T) {
	srv := serverMock()
	ctrl := gomock.NewController(t)
	defer srv.Close()
	// Mock Setting UC
	settingUC := mock_setting.NewMockSettingUseCase(ctrl)
	settingUC.EXPECT().ExistsCanvasConfig().Return(true, srv.URL, "my-secret-token", nil)

	DiscussionRepo := NewRepositoryAPI(http.DefaultClient, settingUC)
	res, err := DiscussionRepo.ListDiscussionByCourseID(1)
	assert.NilError(t, err, "Error List Discussion")
	assert.Equal(t, res[0].ID, uint32(1))
	assert.Equal(t, len(res), 1)
}

func serverMock() *httptest.Server {
	myGin := ginMock()
	srv := httptest.NewServer(myGin)

	return srv
}

func ginMock() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/courses/1/discussion_topics", mock_discussion.ListDiscussionByCourseID())

	return r
}
