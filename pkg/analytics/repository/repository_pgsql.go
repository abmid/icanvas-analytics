/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"

	"github.com/abmid/icanvas-analytics/internal/pagination"
	"github.com/abmid/icanvas-analytics/pkg/analytics/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type pgRepository struct {
	con *sql.DB
	sq  sq.StatementBuilderType
	pag pagination.PaginationInterface
}

var (
	DBNAME = "report_course_results"
)

func NewRepositoryPG(db *sql.DB, pag pagination.PaginationInterface) *pgRepository {
	return &pgRepository{
		con: db,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		pag: pag,
	}
}

// FindBestCourseByFile a function to get all course or by filter
func (r *pgRepository) FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) (res []entity.AnalyticsCourse, pag pagination.Pagination, err error) {
	var query sq.SelectBuilder

	nullVall := entity.FilterAnalytics{}
	// Filter require teacher or not, for require teacher must have another join
	if filter.AnalyticsTeacher == true {
		query = r.sq.Select().From("report_courses").
			Join("report_course_results on report_course_results.report_course_id = report_courses.id").
			Join("report_enrollments on report_enrollments.course_report_id = report_courses.id").Where(sq.Eq{"report_enrollments.role": "TeacherEnrollment"})
	} else {
		query = r.sq.Select().
			From("report_courses").
			Join("report_course_results on report_course_results.report_course_id = report_courses.id")
	}

	// Filter Account ID
	if filter.AccountID != nullVall.AccountID {
		query = query.Where(sq.Eq{"report_courses.account_id": filter.AccountID})
	}

	// Filter search query
	if filter.Q != nullVall.Q {
		query = query.Where(sq.Or{
			// course name
			sq.Like{"lower(report_courses.course_name)": fmt.Sprint("%", strings.ToLower(filter.Q), "%")},
			// teacher full name
			sq.Like{"lower(report_enrollments.full_name)": fmt.Sprint("%", strings.ToLower(filter.Q), "%")},
			// teacher email
			sq.Like{"lower(report_enrollments.login_id)": fmt.Sprint("%", strings.ToLower(filter.Q), "%")},
		})
	}

	// Filter page or offset
	if filter.Page != nullVall.Page {
		//
		var skip, limit uint64
		if limit = 10; filter.Limit != nullVall.Limit {
			limit = filter.Limit
		}
		if skip = 0; filter.Page > 1 {
			skip = (filter.Page - 1) * limit
		}
		query = query.Offset(skip)
	}

	// Filter Date
	if filter.Date != nullVall.Date {
		query = query.Where("date(report_courses.created_at) = ?", filter.Date.Format("2006-01-02"))
	}

	// Build Pagination
	pag, err = r.pag.BuildPagination(query, filter.Limit, filter.Page)
	if err != nil {
		return nil, pagination.Pagination{}, err
	}

	// Filter limit
	if filter.Limit != nullVall.Limit {
		query = query.Limit(filter.Limit)
	} else {
		query = query.Limit(10)
	}

	// Default order by final_score desc
	if filter.OrderBy != nullVall.OrderBy {
		query = query.OrderBy("report_course_results.final_score " + filter.OrderBy)
	} else {
		query = query.OrderBy("report_course_results.final_score desc")
	}

	// Fix persistance result
	// small number course_id = teacher fast to create course
	query = query.OrderBy("report_courses.course_id asc")

	// Set Select
	if filter.AnalyticsTeacher == true {
		// Fix persistance data for teacher
		// small user_id thats mean the teacher fast to create acount
		query = query.OrderBy("report_enrollments.user_id asc")
		// SELECT
		query = query.Columns("report_courses.id").
			Columns("report_courses.account_id").
			Columns("report_courses.course_id").
			Columns("report_courses.course_name").
			Columns("report_course_results.assigment_count").
			Columns("report_course_results.discussion_count").
			Columns("report_course_results.student_count").
			Columns("report_course_results.finish_grading_count").
			Columns("report_course_results.final_score").
			Columns("report_enrollments.full_name").
			Columns("report_enrollments.login_id").
			Columns("report_enrollments.user_id").
			Columns("report_enrollments.role")
	} else {
		query = query.Columns("report_courses.id").
			Columns("report_courses.account_id").
			Columns("report_courses.course_id").
			Columns("report_courses.course_name").
			Columns("report_course_results.assigment_count").
			Columns("report_course_results.discussion_count").
			Columns("report_course_results.student_count").
			Columns("report_course_results.finish_grading_count").
			Columns("report_course_results.final_score")
	}
	query = query.RunWith(r.con)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return res, pagination.Pagination{}, nil
		}
		return nil, pagination.Pagination{}, err
	}
	defer rows.Close()
	for rows.Next() {
		result := entity.AnalyticsCourse{}
		if filter.AnalyticsTeacher == true {
			teacher := canvas.User{}
			err = rows.Scan(
				&result.ID,
				&result.AccountID,
				&result.CourseID,
				&result.CourseName,
				&result.AssigmentCount,
				&result.DiscussionCount,
				&result.StudentCount,
				&result.FinishGradingCount,
				&result.FinalScore,
				&teacher.Name,
				&teacher.LoginID,
				&teacher.ID,
				&teacher.Enrollments,
			)
			result.Teacher = &teacher
			if err != nil {
				return nil, pagination.Pagination{}, err
			}
		} else {
			err = rows.Scan(
				&result.ID,
				&result.AccountID,
				&result.CourseID,
				&result.CourseName,
				&result.AssigmentCount,
				&result.DiscussionCount,
				&result.StudentCount,
				&result.FinishGradingCount,
				&result.FinalScore,
			)
			if err != nil {
				return nil, pagination.Pagination{}, err
			}
		}

		var averageGrading float32
		averageGrading = 0
		if result.StudentCount != 0 {
			// Get Average
			averageGrading = (float32(result.FinishGradingCount) / float32(result.StudentCount)) * 100
		}
		result.AverageGrading = averageGrading

		res = append(res, result)
	}
	return res, pag, nil
}
