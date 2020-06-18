package usecase

import (
	"testing"

	"github.com/abmid/icanvas-analytics/internal/user/entity"
	repo "github.com/abmid/icanvas-analytics/internal/user/repository/mock"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T)  {
	ctrl := gomock.NewController(t)
	mockRepo := repo.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func (r *entity.User) error {
		r.ID = 1
		return nil
	})

	uc := New(mockRepo)
	user := entity.User{
		Name: "test",
	}
	err := uc.Create(&user)
	assert.NilError(t, err)
	assert.Equal(t, user.ID, uint32(1))
}