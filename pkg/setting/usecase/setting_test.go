package usecase

import (
	"testing"

	"github.com/abmid/icanvas-analytics/pkg/setting/entity"
	repo "github.com/abmid/icanvas-analytics/pkg/setting/repository/mock"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCreateOrUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("create", func(t *testing.T) {
		r := repo.NewMockSettingRepository(ctrl)
		// no result mock
		r.EXPECT().FindByFilter(entity.Setting{Category: "canvas"}).Return([]entity.Setting{}, nil)
		// create mock
		r.EXPECT().Create(gomock.Any()).DoAndReturn(func(setting *entity.Setting) error {
			setting.ID = 1
			return nil
		})

		uc := NewSettingUseCase(r)
		setting := entity.Setting{Category: "canvas"}
		err := uc.CreateOrUpdate(&setting)
		assert.NilError(t, err)
		assert.Equal(t, uint32(1), setting.ID)
	})
	t.Run("update", func(t *testing.T) {
		r := repo.NewMockSettingRepository(ctrl)
		setting := entity.Setting{
			Category: "canvas",
			Value:    "new",
		}
		// have result mock
		r.EXPECT().FindByFilter(entity.Setting{Category: "canvas"}).Return([]entity.Setting{{ID: 1, Category: "canvas"}}, nil)
		// Update mock
		r.EXPECT().Update(gomock.Any(), setting).Return(nil)
		uc := NewSettingUseCase(r)
		err := uc.CreateOrUpdate(&setting)
		assert.NilError(t, err)
	})
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	r := repo.NewMockSettingRepository(ctrl)
	r.EXPECT().Create(gomock.Any()).DoAndReturn(func(setting *entity.Setting) error {
		setting.ID = 2
		return nil
	})
	uc := NewSettingUseCase(r)
	setting := entity.Setting{
		Name: "test",
	}
	err := uc.Create(&setting)
	assert.NilError(t, err)
	assert.Equal(t, setting.ID, uint32(2))
}

func TestCreateAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	r := repo.NewMockSettingRepository(ctrl)
	// no result mock
	r.EXPECT().FindByFilter(entity.Setting{Category: "canvas"}).Return([]entity.Setting{}, nil)
	// create mock
	r.EXPECT().Create(gomock.Any()).DoAndReturn(func(setting *entity.Setting) error {
		setting.ID = 1
		return nil
	})
	uc := NewSettingUseCase(r)
	setting := []*entity.Setting{
		{Category: "canvas"},
	}
	err := uc.CreateAll(setting)
	assert.NilError(t, err)
	assert.Equal(t, uint32(1), setting[0].ID)
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	r := repo.NewMockSettingRepository(ctrl)
	r.EXPECT().Update(uint32(2), gomock.Any()).Return(nil)

	uc := NewSettingUseCase(r)
	setting := entity.Setting{}
	err := uc.Update(2, setting)
	assert.NilError(t, err)

}

func TestFindCanvasUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("exist", func(t *testing.T) {
		r := repo.NewMockSettingRepository(ctrl)
		filter := entity.Setting{
			Category: "canvas",
			Name:     "token",
		}
		exResult := []entity.Setting{
			{Value: "secret-token"},
		}
		r.EXPECT().FindByFilter(filter).Return(exResult, nil)

		uc := NewSettingUseCase(r)
		res, err := uc.FindCanvasToken()
		assert.NilError(t, err)
		assert.Equal(t, *res, exResult[0])
	})

	t.Run("not-exist", func(t *testing.T) {
		r := repo.NewMockSettingRepository(ctrl)
		filter := entity.Setting{
			Category: "canvas",
			Name:     "token",
		}
		exResult := []entity.Setting{}
		r.EXPECT().FindByFilter(filter).Return(exResult, nil)

		uc := NewSettingUseCase(r)
		res, err := uc.FindCanvasToken()
		assert.NilError(t, err)
		assert.Check(t, res == nil)
	})
}
