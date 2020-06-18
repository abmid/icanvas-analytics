/*
 * File Created: Saturday, 6th June 2020 10:49:09 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

import "database/sql"

type ReportUser struct {
	ID             uint32
	ReportCourseID uint32
	UserID         uint32
	LoginID        string
	FullName       string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}
