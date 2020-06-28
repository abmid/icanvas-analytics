/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

import "database/sql"

type ReportCourse struct {
	ID         uint32       `json:"id"`
	CourseID   uint32       `json:"course_id"`
	CourseName string       `json:"course_name"`
	AccountID  uint32       `json:"account_id"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
}

type ReportCourseFilter struct {
	ReportCourse ReportCourse
	Limit        uint32
}
