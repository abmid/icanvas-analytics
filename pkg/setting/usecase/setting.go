/*
 * File Created: Thursday, 16th July 2020 4:09:15 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"github.com/abmid/icanvas-analytics/internal/inerr"
	"github.com/abmid/icanvas-analytics/pkg/setting/entity"
	"github.com/abmid/icanvas-analytics/pkg/setting/repository"
)

type settingUseCase struct {
	SettingRepository repository.SettingRepository
}

func NewSettingUseCase(sr repository.SettingRepository) *settingUseCase {
	return &settingUseCase{
		SettingRepository: sr,
	}
}

// CreateOrUpdate is function to create if data not exists and update if data exists
func (UC *settingUseCase) CreateOrUpdate(setting *entity.Setting) error {
	// Check data
	filter := entity.Setting{
		Name:     setting.Name,
		Category: setting.Category,
	}
	exists, err := UC.SettingRepository.FindByFilter(filter)
	if err != nil {
		return err
	}
	if len(exists) == 0 {
		// Data Not Exist, must create
		err := UC.SettingRepository.Create(setting)
		if err != nil {
			return err
		}
		return nil
	}
	// Data exist must update
	err = UC.SettingRepository.Update(exists[0].ID, *setting)
	if err != nil {
		return err
	}
	return nil
}

func (UC *settingUseCase) Create(setting *entity.Setting) error {
	err := UC.SettingRepository.Create(setting)
	if err != nil {
		return err
	}

	return nil
}

func (UC *settingUseCase) CreateAll(settings []*entity.Setting) error {
	for _, setting := range settings {
		// Create or Update
		err := UC.CreateOrUpdate(setting)
		if err != nil {
			return err
		}
	}
	return nil
}

func (UC *settingUseCase) Update(id uint32, setting entity.Setting) error {
	err := UC.SettingRepository.Update(id, setting)
	if err != nil {
		return err
	}

	return nil
}

// FindByFilter is function to get all value of settings or get value by criteria
func (UC *settingUseCase) FindByFilter(filter entity.Setting) (res []entity.Setting, err error) {
	res, err = UC.SettingRepository.FindByFilter(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FindCanvasURL is function to get Quickly value of canvas url
func (UC *settingUseCase) FindCanvasURL() (res *entity.Setting, err error) {
	filter := entity.Setting{
		Category: "canvas",
		Name:     "url",
	}
	settings, err := UC.SettingRepository.FindByFilter(filter)
	if err != nil {
		return nil, err
	}

	if len(settings) == 0 {
		return nil, nil
	}

	// if data exist but empty value
	if settings[0].Value == "" {
		return nil, nil
	}

	return &settings[0], nil
}

// FindCanvasToken is function to get quickly value of canvas token
func (UC *settingUseCase) FindCanvasToken() (res *entity.Setting, err error) {
	filter := entity.Setting{
		Category: "canvas",
		Name:     "token",
	}
	settings, err := UC.SettingRepository.FindByFilter(filter)
	if err != nil {
		return nil, err
	}

	if len(settings) == 0 {
		return nil, nil
	}

	// if data exist but empty value
	if settings[0].Value == "" {
		return nil, nil
	}

	return &settings[0], nil
}

func (UC *settingUseCase) ExistsCanvasConfig() (exists bool, url, token string, err error) {
	resURL, err := UC.FindCanvasURL()
	if err != nil || resURL == nil {
		return false, "", "", inerr.ErrNoCanvasConfig
	}

	resToken, err := UC.FindCanvasToken()
	if err != nil || resToken == nil {
		return false, "", "", inerr.ErrNoCanvasConfig
	}

	return true, resURL.Value, resToken.Value, nil
}
