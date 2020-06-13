package mock_canvas

import (
	"net/http"

	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/gin-gonic/gin"
)

func Courses() gin.HandlerFunc {
	return func(c *gin.Context) {
		listCourse := []entity.Course{
			{ID: 1},
		}
		c.JSON(http.StatusOK, listCourse)
	}
}

func ListUserInCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		teacher := c.Query("enrollment_role")
		ListUserInCourse := []entity.User{}
		if teacher == "TeacherEnrollment" {
			ListUserInCourse = []entity.User{
				{ID: 1, Name: "TestingTeacher User", Enrollments: "TeacherEnrollment"},
			}
		} else {
			ListUserInCourse = []entity.User{
				{ID: 2, Name: "Student User", Enrollments: "StudentEnrollment"},
			}
		}
		c.JSON(http.StatusOK, ListUserInCourse)
	}
}
