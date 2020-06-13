package entity

import "time"

type User struct {
	ID            uint32    `json:"id"`
	Name          string    `json:"name"`
	ShortableName string    `json:"sortable_name"`
	SISUserID     string    `json:"sis_user_id"`
	SISImportID   string    `json:"sis_import_id"`
	IntegrationID string    `json:"integration_id"`
	LoginID       string    `json:"login_id"`
	AvatarURL     string    `json:"avatar_url"`
	Enrollments   string    `json:"enrollments"`
	Email         string    `json:"email"`
	Locale        string    `json:"locale"`
	LastLogin     time.Time `json:"last_login"`
	TimeZone      string    `json:"time_zone"`
	Bio           string    `json:"bio"`
}
