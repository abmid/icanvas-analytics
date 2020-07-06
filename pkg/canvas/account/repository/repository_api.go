/*
 * File Created: Monday, 6th July 2020 2:28:33 pm
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

func (r *Repository) ListAccount(accountID uint32) (res []entity.Account, err error) {
	castStringID := strconv.Itoa(int(accountID))
	req, err := http.NewRequest("GET", r.BaseUrl+"/api/v1/accounts/"+castStringID+"/sub_accounts?per_page=250", nil)
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
