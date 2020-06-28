/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

type ScoreEnrollment struct {
	CourseReportID uint32
	FinishGrading  uint32  // how many student are finish count
	StudentCount   uint32  // how many student
	AverageGrading float32 // GradedCount / StudentCount * 100
}
