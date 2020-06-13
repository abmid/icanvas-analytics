package entity

import "database/sql"

type ReportEnrollment struct {
	ID             uint32       `json:"id"`
	CourseReportID uint32       `json:"course_report_id"`
	EnrollmentID   uint32       `json:"enrollment_id"`
	UserID         uint32       `json:"user_id"`
	LoginID        string       `json:"login_id"`
	FullName       string       `json:"full_name"`
	RoleID         uint32       `json:"role_id"`
	Role           string       `json:"role"`
	CurrentScore   float32      `json:"current_score"`
	CurrentGrade   float32      `json:"current_grade"`
	FinalScore     float32      `json:"final_score"`
	FinalGrade     float32      `json:"final_grade"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}
