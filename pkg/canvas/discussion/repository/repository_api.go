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
	"log"
	"net/http"
	"strconv"

	"github.com/abmid/icanvas-analytics/internal/inerr"
	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"github.com/abmid/icanvas-analytics/pkg/setting/usecase"
)

type Repository struct {
	Client  *http.Client
	Setting usecase.SettingUseCase
	Log     *logger.LoggerWrap
}

func NewRepositoryAPI(client *http.Client, settingUC usecase.SettingUseCase) *Repository {
	logger := logger.New()

	return &Repository{
		Client:  client,
		Setting: settingUC,
		Log:     logger,
	}
}

/**
* https://lms.umm.ac.id/doc/api/discussion_topics.html#method.discussion_topics.index
 */
func (r *Repository) ListDiscussionByCourseID(courseID uint32) (res []entity.Discussion, err error) {
	castCourseID := strconv.Itoa(int(courseID))
	// Check User already setting Canvas configuration
	exists, url, token, err := r.Setting.ExistsCanvasConfig()
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	// if not exist canvas configuration in db
	if !exists {
		return nil, inerr.ErrNoCanvasConfig
	}
	req, err := http.NewRequest("GET", url+"/api/v1/courses/"+castCourseID+"/discussion_topics", nil)
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	req.Header.Set(
		"Authorization", "Bearer "+token,
	)
	resp, err := r.Client.Do(req)
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check if response not success
	if resp.StatusCode != 200 {
		r.Log.Error(err)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	return res, nil
}
