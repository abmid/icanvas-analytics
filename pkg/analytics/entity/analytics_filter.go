/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

import (
	"encoding/json"
	"strings"
	"time"
)

type JsonDate time.Time

// This filter for handle request user
type FilterAnalytics struct {
	AccountID        uint32   `json:"account_id" form:"account_id" query:"account_id"`                      // To Filter By AccountID
	AnalyticsTeacher bool     `json:"analytics_teacher" form:"analytics_teacher" query:"analytics_teacher"` // To Get Analytics With Teacher
	Date             JsonDate `json:"date" form:"date" query:"date"`                                        // Filter By Date
	Page             uint64   `json:"page" form:"page" query:"page"`                                        // Paginate
	Limit            uint64   `json:"limit" form:"limit" query:"limit"`                                     // Limit
	Q                string   `json:"q" form:"q" query:"q"`                                                 // Filter search
	OrderBy          string   `json:"order_by" form:"order_by" query:"order_by"`                            // Order final_score desc or asc
}

// imeplement Marshaler und Unmarshalere interface
func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Maybe a Format function for printing your date
func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j *JsonDate) UnmarshalParam(src string) error {
	s := strings.Trim(string(src), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return err
}
