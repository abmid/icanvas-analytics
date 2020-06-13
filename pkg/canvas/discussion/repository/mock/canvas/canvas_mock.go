package mock_discussion

import (
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListDiscussionByCourseID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ListDiscussion := []entity.Discussion{
			{ID: 1},
		}
		c.JSON(http.StatusOK, ListDiscussion)
	}
}
