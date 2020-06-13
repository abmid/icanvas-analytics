package entity

import "time"

type Assigment struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CourseID  uint32    `json:"course_id"`
}
