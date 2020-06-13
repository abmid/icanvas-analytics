package repository

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock_assigment "github.com/abmid/icanvas-analytics/pkg/canvas/assigment/repository/mock/canvas"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestListAssigmentByCourseID(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	AssigmentRepo := NewRepositoryAPI(http.DefaultClient, srv.URL, "my-secret-token")
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
