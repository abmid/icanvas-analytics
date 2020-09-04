/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/abmid/icanvas-analytics/internal/inerr"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"github.com/abmid/icanvas-analytics/pkg/setting/usecase"
)

type Repository struct {
	Client  *http.Client
	Setting usecase.SettingUseCase
}

func NewRepositoryAPI(client *http.Client, settingUC usecase.SettingUseCase) *Repository {
	return &Repository{
		Client:  client,
		Setting: settingUC,
	}
}

func (r *Repository) ListAssigmentByCourseID(CourseID uint32) (res []entity.Assigment, err error) {
	caseCourseID := strconv.Itoa(int(CourseID))
	// Check User already setting Canvas configuration
	exists, url, token, err := r.Setting.ExistsCanvasConfig()
	if err != nil {
		return nil, err
	}
	// if not exist canvas configuration in db
	if !exists {
		return nil, inerr.ErrNoCanvasConfig
	}
	req, err := http.NewRequest("GET", url+"/api/v1/courses/"+caseCourseID+"/assignments", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"Authorization", "Bearer "+token,
	)
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
