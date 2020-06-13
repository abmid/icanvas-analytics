package entity

import "database/sql"

type ReportDiscussion struct {
	ID             uint32       `json:"id"`
	CourseReportID uint32       `json:"course_report_id"`
	DiscussionID   uint32       `json:"discussion_id"`
	Title          string       `json:"title"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}
