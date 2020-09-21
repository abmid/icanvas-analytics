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
	"time"

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

func (r *Repository) ListEnrollmentByCourseID(courseID uint32) (res []entity.Enrollment, err error) {
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
	req, err := http.NewRequest("GET", url+"/api/v1/courses/"+castCourseID+"/enrollments", nil)
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok && jsonErr.Value == "string" {
			res, err = fixErrorUnmarshalStringJSON(body)
			if err != nil {
				r.Log.Error(err)
				return nil, err
			}
			return
		}
		return nil, err
	}
	return res, nil
}

func fixErrorUnmarshalStringJSON(body []byte) (res []entity.Enrollment, err error) {
	var enrollments []map[string]interface{}

	err = json.Unmarshal(body, &enrollments)
	if err != nil {
		return nil, err
	}
	for _, enrollment := range enrollments {
		enGrade := enrollment["grades"].(map[string]interface{})
		tmpEnroll := entity.Enrollment{
			ID:        safeGetUint(enrollment["id"]),
			CourseID:  safeGetUint(enrollment["course_id"]),
			UserID:    safeGetUint(enrollment["user_id"]),
			RoleID:    safeGetUint(enrollment["role_id"]),
			Role:      enrollment["role"].(string),
			Type:      enrollment["type"].(string),
			CreatedAt: safeGetTime(enrollment["created_at"]),
			UpdatedAt: safeGetTime(enrollment["updated_at"]),
			Grades: entity.EnrollmentGrade{
				HtmlURL:      enGrade["html_url"].(string),
				CurrentGrade: 0,
				CurrentScore: safeGetFloat32(enGrade["current_score"]),
				FinalGrade:   safeGetFloat32(enGrade["final_grade"]),
				FinalScore:   safeGetFloat32(enGrade["final_score"]),
			},
		}
		res = append(res, tmpEnroll)
	}
	return res, nil
}

func safeGetUint(i interface{}) uint32 {
	switch i.(type) {
	case int:
		return uint32(i.(int))
	case int16:
		return uint32(i.(int16))
	case int32:
		return uint32(i.(int32))
	case int64:
		return uint32(i.(int64))
	case float32:
		return uint32(i.(float32))
	case float64:
		return uint32(i.(float64))
	default:
		return uint32(0)
	}
}

func safeGetFloat32(i interface{}) float32 {
	switch i.(type) {
	case float32:
		return float32(i.(float32))
	case float64:
		return float32(i.(float64))
	default:
		return float32(0)
	}
}

func safeGetTime(i interface{}) time.Time {
	switch i.(type) {
	case string:
		if i.(string) == "" {
			value := "2000-01-13T00:00:00+00:00"
			tm, _ := time.Parse(time.RFC3339, value)
			return tm
		}
		parseTime, _ := time.Parse(time.RFC3339, i.(string))
		return parseTime
	default:
		value := "2000-01-13T00:00:00+00:00"
		tm, _ := time.Parse(time.RFC3339, value)
		return tm
	}
}
