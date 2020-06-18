package entity

import "time"

// This filter for handle request user
type FilterAnalytics struct {
	AccountID        uint32    `json:"account_id" form:"account_id" query:"account_id"`                      // To Filter By AccountID
	AnalyticsTeacher bool      `json:"analytics_teacher" form:"analytics_teacher" query:"analytics_teacher"` // To Get Analytics With Teacher
	Date             time.Time `json:"date" form:"date" query:"date" time_format:"2006-01-02"`               // Filter By Date
	Page             uint64    `json:"page" form:"page" query:"page"`                                        // Paginate
	Limit            uint64    `json:"limit" form:"limit" query:"limit"`                                     // Limit
	OrderBy          string    `json:"order_by" form:"order_by" query:"order_by"`                            // Order final_score desc or asc
}
