package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type Repository struct {
	Client      *http.Client
	BaseUrl     string
	AccessToken string
}

func NewRepositoryAPI(client *http.Client, baseUrl, accessToken string) *Repository {
	return &Repository{
		Client:      client,
		BaseUrl:     baseUrl,
		AccessToken: accessToken,
	}
}

func (r *Repository) ListAssigmentByCourseID(CourseID uint32) (res []entity.Assigment, err error) {
	caseCourseID := strconv.Itoa(int(CourseID))
	req, err := http.NewRequest("GET", r.BaseUrl+"/api/v1/courses/"+caseCourseID+"/assignments", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"Authorization", "Bearer "+r.AccessToken,
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
