package mock_assigment

import (
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAssigmentByCourseID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ListAssigment []entity.Assigment
		castCourseID, err := strconv.Atoi(c.Param("courseID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, "CourseID not found")
		}
		if castCourseID == 1 {
			ListAssigment = []entity.Assigment{
				{ID: 1},
			}
		}
		c.JSON(http.StatusOK, ListAssigment)
	}
}
