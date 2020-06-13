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
