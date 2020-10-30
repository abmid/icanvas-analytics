package http

import (
	"net/http"
	"sync"

	analyticsjob "github.com/abmid/icanvas-analytics/pkg/analyticsjob/usecase"
	"github.com/abmid/icanvas-analytics/pkg/auth"
	wsUC "github.com/abmid/icanvas-analytics/pkg/websocket/usecase"
	echo "github.com/labstack/echo/v4"
)

type AnalyticsJobHandler struct {
	AJobUseCase *analyticsjob.AnalyticJobUseCase
	WSUseCase   wsUC.WebsocketUseCase
}

func (handler *AnalyticsJobHandler) GenerateNow() echo.HandlerFunc {
	return func(c echo.Context) error {
		handler.WSUseCase.SendMessageToAll("Please wait to generate analytics. If already finish, you will behave a notification in here")
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			finish := handler.AJobUseCase.RunJob(1)
			if <-finish {
				handler.WSUseCase.SendMessageToAll("Generate data already finished, please check new data in analytics")
			}
		}()

		return c.JSON(http.StatusOK, "OK")
	}
}

func NewHandler(path string, g *echo.Group, JWTKey string, AJobUseCase *analyticsjob.AnalyticJobUseCase, wsUC wsUC.WebsocketUseCase) {
	handler := AnalyticsJobHandler{
		AJobUseCase: AJobUseCase,
		WSUseCase:   wsUC,
	}

	r := g.Group(path)
	r.Use(auth.MiddlewareAuthJWT(JWTKey))

	r.POST("/generate-now", handler.GenerateNow())
}
