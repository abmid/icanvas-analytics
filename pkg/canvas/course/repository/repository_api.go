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
	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"github.com/abmid/icanvas-analytics/pkg/setting/usecase"
)

type APIRepository struct {
	Client  *http.Client
	Setting usecase.SettingUseCase
	Log     *logger.LoggerWrap
}

func NewRepositoryAPI(client *http.Client, settingUC usecase.SettingUseCase) *APIRepository {

	logger := logger.New()

	return &APIRepository{
		Client:  client,
		Setting: settingUC,
		Log:     logger,
	}
}

func (r *APIRepository) Courses(accountId, page uint32) (res []entity.Course, err error) {
	castPage := strconv.Itoa(int(page))
	castAccountID := strconv.Itoa(int(accountId))
	// Check User already setting Canvas configuration
	exists, url, token, err := r.Setting.ExistsCanvasConfig()
	if err != nil {
		return nil, err
	}
	// if not exist canvas configuration in db
	if !exists {
		return nil, inerr.ErrNoCanvasConfig
	}
	req, err := http.NewRequest("GET", url+"/api/v1/accounts/"+castAccountID+"/courses?per_page=50&page="+castPage, nil)
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	req.Header.Add(
		"Authorization", "Bearer "+token,
	)
	resp, err := r.Client.Do(req)
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Log.Error(err)
		return res, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}

	return res, nil
}

/**
* https://lms.umm.ac.id/doc/api/courses.html#method.courses.users
 */
func (r *APIRepository) ListUserInCourse(courseID uint32, enrollmentRole string) (res []entity.User, err error) {
	castCourseID := strconv.Itoa(int(courseID))
	// Check User already setting Canvas configuration
	exists, url, token, err := r.Setting.ExistsCanvasConfig()
	if err != nil {
		return nil, err
	}
	// if not exist canvas configuration in db
	if !exists {
		return nil, inerr.ErrNoCanvasConfig
	}
	req, err := http.NewRequest("GET", url+"/api/v1/courses/"+castCourseID+"/users?enrollment_role="+enrollmentRole, nil)
	if err != nil {
		r.Log.Error(err)
		return res, err
	}
	req.Header.Add(
		"Authorization", "Bearer "+token,
	)
	resp, err := r.Client.Do(req)
	if err != nil {
		r.Log.Error(err)
		return res, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Log.Error(err)
		return res, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		r.Log.Error(err)
		return res, err
	}
	return res, nil
}

// http://www.inanzzz.com/index.php/post/fb0m/mocking-and-testing-http-clients-in-golang
// https://www.thegreatcodeadventure.com/mocking-http-requests-in-golang/
