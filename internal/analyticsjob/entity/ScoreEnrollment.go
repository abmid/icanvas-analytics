package entity

type ScoreEnrollment struct {
	CourseReportID uint32
	FinishGrading  uint32  // how many student are finish count
	StudentCount   uint32  // how many student
	AverageGrading float32 // GradedCount / StudentCount * 100
}
