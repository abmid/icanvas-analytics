package mock_enrollment

import (
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ListEnrollmentByCourseID() gin.HandlerFunc {
	return func(c *gin.Context) {
		EnrollmentGrade := entity.EnrollmentGrade{
			HtmlURL:      "html_url",
			CurrentGrade: 80.5,
			CurrentScore: 13,
			FinalScore:   12.2,
			FinalGrade:   11.2,
		}
		ListEnrollment := []entity.Enrollment{
			{ID: 1, Grades: EnrollmentGrade, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		c.JSON(http.StatusOK, ListEnrollment)
	}
}
