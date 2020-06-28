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

	sq "github.com/Masterminds/squirrel"
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

func (r *pgRepository) FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) (res []entity.AnalyticsCourse, err error) {
	var query sq.SelectBuilder

	nullVall := entity.FilterAnalytics{}
	if filter.AnalyticsTeacher == true {
		query = r.sq.Select(
			"report_courses.id",
			"report_courses.account_id",
			"report_courses.course_id",
			"report_courses.course_name",
			"report_course_results.assigment_count",
			"report_course_results.discussion_count",
			"report_course_results.student_count",
			"report_course_results.finish_grading_count",
			"report_course_results.final_score",
			"report_enrollments.full_name",
			"report_enrollments.login_id",
			"report_enrollments.user_id",
			"report_enrollments.role",
		).
			From("report_courses").
			Join("report_course_results on report_course_results.report_course_id = report_courses.id").
			Join("report_enrollments on report_enrollments.course_report_id = report_courses.id").Where(sq.Eq{"report_enrollments.role": "TeacherEnrollment"})
	} else {
		query = r.sq.Select(
			"report_courses.id",
			"report_courses.account_id",
			"report_courses.course_id",
			"report_courses.course_name",
			"report_course_results.assigment_count",
			"report_course_results.discussion_count",
			"report_course_results.student_count",
			"report_course_results.finish_grading_count",
			"report_course_results.final_score").
			From("report_courses").
			Join("report_course_results on report_course_results.report_course_id = report_courses.id")
	}

	if filter.AccountID != nullVall.AccountID {
		query = query.Where(sq.Eq{"report_courses.account_id": filter.AccountID})
	}
	if filter.Page != nullVall.Page {
		query = query.Offset(filter.Page)
	}
	if filter.Date != nullVall.Date {
		query = query.Where("date(report_courses.created_at) = ?", filter.Date.Format("2006-01-02"))
	}
	if filter.Limit != nullVall.Limit {
		query = query.Limit(filter.Limit)
	} else {
		query = query.Limit(10)
	}

	if filter.OrderBy != nullVall.OrderBy {
		query = query.OrderBy("report_course_results.final_score " + filter.OrderBy)
	} else {
		query = query.OrderBy("report_course_results.final_score desc")
	}

	query = query.RunWith(r.con)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return res, nil
		}
		return nil, err
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
				return nil, err
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
				return nil, err
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
	return res, nil
}
