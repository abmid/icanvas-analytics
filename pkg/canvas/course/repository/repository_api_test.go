package repository

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock_canvas "github.com/abmid/icanvas-analytics/pkg/canvas/course/repository/mock/canvas"
	mock_setting "github.com/abmid/icanvas-analytics/pkg/setting/usecase/mock"
	"github.com/golang/mock/gomock"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestCourses(t *testing.T) {
	srv := serverMock()
	ctrl := gomock.NewController(t)
	defer srv.Close()
	// Mock Setting UC
	settingUC := mock_setting.NewMockSettingUseCase(ctrl)
	settingUC.EXPECT().ExistsCanvasConfig().Return(true, srv.URL, "my-secret-token", nil)

	CourseRepo := NewRepositoryAPI(http.DefaultClient, settingUC)
	res, err := CourseRepo.Courses(1, 1)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, res[0].ID, uint32(1))
	assert.NilError(t, err, "Courses Error")
}

func TestListUserInCourse(t *testing.T) {
	srv := serverMock()
	ctrl := gomock.NewController(t)
	defer srv.Close()
	// Mock Setting UC
	settingUC := mock_setting.NewMockSettingUseCase(ctrl)
	settingUC.EXPECT().ExistsCanvasConfig().Return(true, srv.URL, "my-secret-token", nil)

	CourseRepo := NewRepositoryAPI(http.DefaultClient, settingUC)
	res, err := CourseRepo.ListUserInCourse(1, "TeacherEnrollment")
	assert.NilError(t, err, "Have Error Get List User")
	assert.Equal(t, res[0].ID, uint32(1))
	assert.Equal(t, res[0].Enrollments, "TeacherEnrollment")
}

func serverMock() *httptest.Server {
	myGin := ginMock()
	srv := httptest.NewServer(myGin)

	return srv
}

func ginMock() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/api/v1/accounts/1/courses", mock_canvas.Courses())
	r.GET("/api/v1/courses/1/users", mock_canvas.ListUserInCourse())

	return r
}
