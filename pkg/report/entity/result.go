/*
 * File Created: Thursday, 4th June 2020 3:07:34 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

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
