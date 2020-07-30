package usecase

import "github.com/abmid/icanvas-analytics/pkg/setting/entity"

type SettingUseCase interface {
	CreateOrUpdate(setting *entity.Setting) error
	Create(setting *entity.Setting) error
	CreateAll(settings []*entity.Setting) error
	Update(id uint32, setting entity.Setting) error
	FindByFilter(filter entity.Setting) ([]entity.Setting, error)
	FindCanvasURL() (*entity.Setting, error)
	FindCanvasToken() (*entity.Setting, error)
}
