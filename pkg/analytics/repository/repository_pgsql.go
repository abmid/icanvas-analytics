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
	"math"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/abmid/icanvas-analytics/internal/pagination"
	"github.com/abmid/icanvas-analytics/pkg/analytics/entity"
	canvas "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type pgRepository struct {
	con *sql.DB
	sq  sq.StatementBuilderType
}

var (
	DBNAME = "report_course_results"
)

func NewRepositoryPG(db *sql.DB) *pgRepository {
	return &pgRepository{
		con: db,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgRepository) buildPaginationInfo(query sq.SelectBuilder, filter entity.FilterAnalytics) (res pagination.Pagination, err error) {
	// total = get sum querySQL
	// per_page = get queryURL ? default
	// Current page = get page ? 1
	// Last page = total / per_page (ceil)
	// next_page_url = current_page + 1
	// prev_page_url = current_page != 1 ? current_page - 1 : null
	// from :
	//  -> if current_page == 1 ? 1
	//  -> if current_page == 2 ? per_page + 1
	//  -> if current_page > 2 ? per_page * (current_page - 1) + 1
	// to :
	//  -> if current_page == 1 ? per_page
	//  -? if current_page > 1 ? per_page * current_page
	var total uint32
	var nullFilter entity.FilterAnalytics
	pagQuery := query
	pagQuery = query.Columns("count(*) as total_count")
	pagQuery = pagQuery.RemoveLimit()
	pagQuery = pagQuery.RemoveOffset()
	pagQuery = pagQuery.RunWith(r.con)
	err = pagQuery.QueryRow().Scan(&total)
	if err != nil {
		return res, err
	}
	// Set Total
	res.Total = total
	// Set PerPage
	// if filter.limit is null per_page default is 10
	if res.PerPage = 10; filter.Limit != nullFilter.Limit {
		res.PerPage = uint32(filter.Limit)
	}

	// Set Current Page
	// if filter more than 1 current page same with filter.page
	if res.CurrentPage = 1; filter.Page > 1 {
		res.CurrentPage = uint32(filter.Page)
	}

	// Set LastPage
	calcLastPage := float64(total) / float64(res.PerPage)
	// Ceil returns the least integer value greater than or equal to x.
	res.LastPage = uint32(math.Ceil(calcLastPage))

	// Set NextPage URL
	res.NextPageUrl = strconv.Itoa(int(res.CurrentPage + uint32(1)))

	// Set PrevPage URL
	if res.PrevPageUrl = "1"; res.CurrentPage > 1 {
		sum := res.CurrentPage - 1
		res.PrevPageUrl = strconv.Itoa(int(sum))
	}

	// Set From
	// If current page == 2 just add 1
	if res.From = 1; res.CurrentPage == 2 {
		res.From = res.PerPage + 1
	}
	// if current page > 2, limit * currentpage - 1 and + 1 for initial first number
	if res.CurrentPage > 2 {
		res.From = res.PerPage*(res.CurrentPage-1) + 1
	}

	// Set To
	// ex. Total 17
	// Last page is 2
	// Page 1 == 10
	// Page 2 == 7
	// if general
	if res.To = res.PerPage; res.CurrentPage > 1 {
		res.To = res.PerPage * res.CurrentPage
	}
	// for get difference if currentPage == lastPage -> total - (limit * (lastPage - 1))
	if res.CurrentPage == res.LastPage {
		difference := res.Total - (res.PerPage * (res.LastPage - 1))
		// and then sum (limit * lastPage - 1) + difference
		res.To = (res.PerPage * (res.LastPage - 1)) + difference
	}
	return res, nil
}

func (r *pgRepository) FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) (res []entity.AnalyticsCourse, pag pagination.Pagination, err error) {
	var query sq.SelectBuilder

	nullVall := entity.FilterAnalytics{}
	// Set Join
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

	// Get query first for paginate
	queryPaginate := query
	pag, err = r.buildPaginationInfo(queryPaginate, filter)
	if err != nil {
		return nil, pag, err
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
			return res, pag, nil
		}
		return nil, pag, err
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
				return nil, pag, err
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
				return nil, pag, err
			}
		}

		var averageGrading float32
		averageGrading = 0
		if result.StudentCount != 0 {
			averageGrading = (float32(result.FinishGradingCount) / float32(result.StudentCount)) * 100
		}
		result.AverageGrading = averageGrading

		res = append(res, result)
	}
	return res, pag, nil
}
