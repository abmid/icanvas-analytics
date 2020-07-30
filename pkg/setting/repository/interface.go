package repository

import (
	"github.com/abmid/icanvas-analytics/pkg/setting/entity"
)

type SettingRepository interface {
	Create(setting *entity.Setting) error
	Update(id uint32, setting entity.Setting) error
	FindByFilter(filter entity.Setting) ([]entity.Setting, error)
}
