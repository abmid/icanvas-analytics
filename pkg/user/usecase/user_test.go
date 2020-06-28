package usecase

import (
	"testing"

	"github.com/abmid/icanvas-analytics/pkg/user/entity"
	repo "github.com/abmid/icanvas-analytics/pkg/user/repository/mock"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repo.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(r *entity.User) error {
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

func TestFind(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("exist", func(t *testing.T) {

		mockRepo := repo.NewMockUserRepository(ctrl)
		user := entity.User{
			Email: "test@test.com",
		}
		mockRepo.EXPECT().Find("test@test.com").Return(&user, nil)

		uc := New(mockRepo)
		res, err := uc.Find("test@test.com")
		assert.NilError(t, err)
		assert.Equal(t, res.Email, "test@test.com")
	})
	t.Run("not-exists", func(t *testing.T) {
		mockRepo := repo.NewMockUserRepository(ctrl)
		mockRepo.EXPECT().Find("test@test.com").Return(nil, nil)

		uc := New(mockRepo)
		res, err := uc.Find("test@test.com")
		assert.NilError(t, err)
		assert.Equal(t, res == nil, true)
	})
}
