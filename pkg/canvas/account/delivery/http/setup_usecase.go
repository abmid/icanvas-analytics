package http

import (
	"net/http"

	account_repo "github.com/abmid/icanvas-analytics/pkg/canvas/account/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/account/usecase"
	account_usecase "github.com/abmid/icanvas-analytics/pkg/canvas/account/usecase"
	setting_usecase "github.com/abmid/icanvas-analytics/pkg/setting/usecase"
)

func SetupUseCase(settingUC setting_usecase.SettingUseCase) usecase.AccountUseCase {
	client := http.DefaultClient
	accountRepo := account_repo.NewRepositoryAPI(client, settingUC)
	accountUC := account_usecase.NewUseCase(accountRepo)

	return accountUC
}
