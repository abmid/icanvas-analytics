package http

import (
	"net/http"

	account_repo "github.com/abmid/icanvas-analytics/pkg/canvas/account/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/account/usecase"
	account_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/account/usecase"
)

func SetupUseCase(canvasUrl, canvasAccessToken string) usecase.AccountUseCase {
	client := http.DefaultClient
	accountRepo := account_repo.NewRepositoryAPI(client, canvasUrl, canvasAccessToken)
	accountUC := account_usecase.NewUseCase(accountRepo)

	return accountUC
}
