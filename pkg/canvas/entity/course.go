package entity

import (
	"time"
)

type Course struct {
	ID            uint32    `json:"id"`
	Name          string    `json:"name"`
	AccountID     uint32    `json:"account_id"`
	CourseCode    string    `json:"course_code"`
	WorkflowStat  string    `json:"workflow_state"`
	RootAccountID uint32    `json:"root_account_id"`
	CreatedAt     time.Time `json:"created_at"`
	StartAt       time.Time `json:"start_at"`
	EndAt         time.Time `json:"end_at"`
}
