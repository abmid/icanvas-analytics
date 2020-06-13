package repository

import (
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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

/**
* https://lms.umm.ac.id/doc/api/discussion_topics.html#method.discussion_topics.index
 */
func (r *Repository) ListDiscussionByCourseID(courseID uint32) (res []entity.Discussion, err error) {
	castCourseID := strconv.Itoa(int(courseID))
	req, err := http.NewRequest("GET", r.BaseUrl+"/api/v1/courses/"+castCourseID+"/discussion_topics", nil)
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
