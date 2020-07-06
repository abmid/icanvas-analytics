package usecase

import (
	"testing"

	account_repo_mock "github.com/abmid/icanvas-analytics/pkg/canvas/account/repository/mock"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestListAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := account_repo_mock.NewMockAccountRepository(ctrl)
	ListAccount := []entity.Account{
		{ID: 1},
	}
	repo.EXPECT().ListAccount(uint32(1)).Return(ListAccount, nil)

	UC := NewUseCase(repo)
	res, err := UC.ListAccount(uint32(1))
	assert.NilError(t, err)
	assert.Equal(t, len(res), len(ListAccount))
}
