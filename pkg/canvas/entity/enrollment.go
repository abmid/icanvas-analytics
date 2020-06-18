/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

import "time"

type EnrollmentGrade struct {
	HtmlURL      string  `json:"html_url"`
	CurrentScore float32 `json:"current_score"`
	CurrentGrade float32 `json:"current_grade"` //fix json: cannot unmarshal string into Go struct type float32
	FinalScore   float32 `json:"final_score"`
	FinalGrade   float32 `json:"final_grade"`
}

type Enrollment struct {
	ID        uint32          `json:"id"`
	CourseID  uint32          `json:"course_id"`
	Type      string          `json:"type"`
	UserID    uint32          `json:"user_id"`
	Role      string          `json:"role"`
	RoleID    uint32          `json:"role_id"`
	Grades    EnrollmentGrade `json:"grades"`
	User      User            `json:"user"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
