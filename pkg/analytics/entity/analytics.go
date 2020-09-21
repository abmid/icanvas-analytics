/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

import canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"

type AnalyticsCourse struct {
	ID                 uint32       `json:"id"`
	AccountID          uint32       `json:"account_id"`
	CourseID           uint32       `json:"course_id"`
	CourseName         string       `json:"course_name"`
	AssigmentCount     uint32       `json:"assigment_count"`
	DiscussionCount    uint32       `json:"discussion_count"`
	StudentCount       uint32       `json:"student_count"`
	FinishGradingCount uint32       `json:"finish_grading_count"`
	AverageGrading     float32      `json:"average_grading"` //Average Grading Teacher
	FinalScore         float64      `json:"final_score"`     // calculate (assigmentCount + Disscussion Count + Average Grade) / 3
	Teacher            *canvas.User `json:"teacher,omitempty"`
}
