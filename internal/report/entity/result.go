package entity

import "database/sql"

type ReportResult struct {
	ID                 uint32
	ReportCourseID     uint32
	AssigmentCount     uint32
	DiscussionCount    uint32
	StudentCount       uint32
	FinishGradingCount uint32
	FinalScore         float32
	CreatedAt          sql.NullTime
	UpdatedAt          sql.NullTime
}

// id serial PRIMARY KEY,
// report_course_id integer references report_courses(id),
// assigment_count integer,
// discussion_count integer,
// student_count integer,
// finish_grading_count integer,
// final_score float,
// created_at timestamp,
// deleted_at timestamp
