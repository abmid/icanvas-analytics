package entity

import "database/sql"

type ReportAssigment struct {
	ID             uint32       `json:"id"`
	CourseReportID uint32       `json:"course_report_id"`
	AssigmentID    uint32       `json:"assigment_id"`
	Name           string       `json:"name"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}
