package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_canvas "github.com/abmid/icanvas-analytics/pkg/canvas/course/repository/mock/canvas"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestCourses(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	CourseRepo := NewRepositoryAPI(http.DefaultClient, srv.URL, "my-secret-token")
	res, err := CourseRepo.Courses(1, 1)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, res[0].ID, uint32(1))
	assert.NilError(t, err, "Courses Error")
}

func TestListUserInCourse(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	CourseRepo := NewRepositoryAPI(http.DefaultClient, srv.URL, "my-secret-token")
	res, err := CourseRepo.ListUserInCourse(1, "TeacherEnrollment")
	t.Log(res)
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

func TestMock(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	resp, _ := http.Get(srv.URL + "/api/v1/course/1/users")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := []entity.User{}
	_ = json.Unmarshal(body, &res)
	t.Log(res)
}

// func TestCourses2(t *testing.T) {
// 	MyClient := RealHTTP()
// 	CourseRepo := NewRepositoryAPI(MyClient, "https://lms.umm.ac.id", "2Q8LJIJs7gCo8XsftFOtq53UT3cUlBIHsTQS7WAi6Le0TTjT2sL7bNtkm5ERT7cb")
// 	courses, err := CourseRepo.Courses(39, 1)
// 	if err != nil {
// 		t.Logf("Error Get Course %s", err)
// 	}
// 	for _, course := range courses {
// 		t.Log(course)
// 	}
// 	t.Fatalf("Pause")
// }
