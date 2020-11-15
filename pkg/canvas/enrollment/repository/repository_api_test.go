package repository

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock_enrollment "github.com/abmid/icanvas-analytics/pkg/canvas/enrollment/repository/mock/canvas"
	mock_setting "github.com/abmid/icanvas-analytics/pkg/setting/usecase/mock"
	"github.com/golang/mock/gomock"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestListEnrollmentByCourseID(t *testing.T) {
	srv := serverMock()
	ctrl := gomock.NewController(t)
	defer srv.Close()
	// Mock Setting UC
	settingUC := mock_setting.NewMockSettingUseCase(ctrl)
	settingUC.EXPECT().ExistsCanvasConfig().Return(true, srv.URL, "my-secret-token", nil)

	EnrollmentRepo := NewRepositoryAPI(http.DefaultClient, settingUC)
	res, err := EnrollmentRepo.ListEnrollmentByCourseID(1)
	t.Log(res)
	assert.NilError(t, err, "Error List Enrollment")
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
	r.GET("/api/v1/courses/1/enrollments", mock_enrollment.ListEnrollmentByCourseID())
	return r
}

func TestFixErrorUnmarshalStringJSON(t *testing.T) {
	jsonn := `[
		{
			"id": 1,
			"course_id": 1,
			"user_id": 1,
			"role_id": 1,
			"role": "role",
			"type": "type",
			"created_at": "2019-09-23T12:50:28+07:00",
			"updated_at": "2019-09-23T12:50:28+07:00",
			"grades": {
				"html_url": "string",
				"current_grade": "",
				"current_score": 32.3,
				"final_grade": 32.3,
				"final_score": 33.3
			}
		}
	]`
	rr, err := fixErrorUnmarshalStringJSON([]byte(jsonn))
	t.Log(rr)
	assert.NilError(t, err, "Error Safe Get Int")
	assert.Equal(t, rr[0].ID, uint32(1))
}

func TestSafeGetUint(t *testing.T) {
	var TestExpectation uint32
	TestExpectation = 10
	assert.Equal(t, safeGetUint(int32(10)), TestExpectation)
	assert.Equal(t, safeGetUint(int64(10)), TestExpectation)
	assert.Equal(t, safeGetUint(int16(10)), TestExpectation)
}

func TestSafeGetFloat32(t *testing.T) {
	var TestExpectation float32
	TestExpectation = 33.3
	assert.Equal(t, safeGetFloat32(float64(33.3)), TestExpectation)
	assert.Equal(t, safeGetFloat32(float32(33.3)), TestExpectation)
}
func TestSafeGetTime(t *testing.T) {
	value := "2019-09-23T12:50:28+07:00"
	exceptation, _ := time.Parse(time.RFC3339, value)
	assert.Equal(t, safeGetTime(value).String(), exceptation.String())
}
