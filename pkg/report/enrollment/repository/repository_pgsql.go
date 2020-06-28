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
	"time"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"

	sq "github.com/Masterminds/squirrel"
)

type pgRepository struct {
	con *sql.DB
	sq  sq.StatementBuilderType
}

var (
	DBTABLE = "report_enrollments"
	POOL    = 2
)

func NewEnrollmentPG(con *sql.DB) *pgRepository {
	return &pgRepository{
		con: con,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgRepository) Create(ctx context.Context, reportEnroll *entity.ReportEnrollment) error {
	query := r.sq.Insert(DBTABLE).Columns(
		"course_report_id",
		"enrollment_id",
		"user_id",
		"login_id",
		"full_name",
		"role_id",
		"role",
		"current_score",
		"current_grade",
		"final_score",
		"final_grade",
		"created_at",
		"updated_at",
	).Values(
		reportEnroll.CourseReportID,
		reportEnroll.EnrollmentID,
		reportEnroll.UserID,
		reportEnroll.LoginID,
		reportEnroll.FullName,
		reportEnroll.RoleID,
		reportEnroll.Role,
		reportEnroll.CurrentScore,
		reportEnroll.CurrentGrade,
		reportEnroll.FinalScore,
		reportEnroll.FinalGrade,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING \"id\"").RunWith(r.con)
	err := query.QueryRow().Scan(&reportEnroll.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Read(ctx context.Context) ([]entity.ReportEnrollment, error) {
	results := []entity.ReportEnrollment{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"enrollment_id",
		"user_id",
		"login_id",
		"full_name",
		"role_id",
		"role",
		"current_score",
		"current_grade",
		"final_score",
		"final_grade",
		"created_at",
		"updated_at",
	).From(DBTABLE).RunWith(r.con)

	rows, err := query.Query()
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		result := entity.ReportEnrollment{}
		err = rows.Scan(
			&result.ID,
			&result.CourseReportID,
			&result.EnrollmentID,
			&result.UserID,
			&result.LoginID,
			&result.FullName,
			&result.RoleID,
			&result.Role,
			&result.CurrentScore,
			&result.CurrentGrade,
			&result.FinalScore,
			&result.FinalGrade,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *pgRepository) Update(ctx context.Context, reportEnroll *entity.ReportEnrollment) error {
	query := r.sq.Update(DBTABLE).
		Set("course_report_id", reportEnroll.CourseReportID).
		Set("enrollment_id", reportEnroll.EnrollmentID).
		Set("user_id", reportEnroll.UserID).
		Set("login_id", reportEnroll.LoginID).
		Set("full_name", reportEnroll.FullName).
		Set("role_id", reportEnroll.RoleID).
		Set("role", reportEnroll.Role).
		Set("current_score", reportEnroll.CurrentScore).
		Set("current_grade", reportEnroll.CurrentGrade).
		Set("final_score", reportEnroll.FinalScore).
		Set("final_grade", reportEnroll.FinalGrade).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": reportEnroll.ID}).Suffix("RETURNING \"id\"").RunWith(r.con)
	err := query.QueryRow().Scan(&reportEnroll.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Delete(ctx context.Context, reportEnrollID uint32) error {
	query := r.sq.Delete(DBTABLE).Where(sq.Eq{"id": reportEnrollID}).RunWith(r.con)
	_, err := query.Exec()
	if err != nil {
		return nil
	}
	return nil
}

func (r *pgRepository) FindFirstByFilter(ctx context.Context, filter entity.ReportEnrollment) (res entity.ReportEnrollment, err error) {
	nullValue := entity.ReportEnrollment{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"enrollment_id",
		"user_id",
		"login_id",
		"full_name",
		"role_id",
		"role",
		"current_score",
		"current_grade",
		"final_score",
		"final_grade",
		"created_at",
		"updated_at",
	).From(DBTABLE)

	// TODO : Filter
	if filter.ID != nullValue.ID {
		query = query.Where(sq.Eq{"id": filter.ID})
	}
	if filter.CourseReportID != nullValue.CourseReportID {
		query = query.Where(sq.Eq{"course_report_id": filter.CourseReportID})
	}
	if filter.EnrollmentID != nullValue.EnrollmentID {
		query = query.Where(sq.Eq{"enrollment_id": filter.EnrollmentID})
	}
	// RUN find First
	query = query.Limit(1).RunWith(r.con)

	err = query.QueryRow().Scan(
		&res.ID,
		&res.CourseReportID,
		&res.EnrollmentID,
		&res.UserID,
		&res.LoginID,
		&res.FullName,
		&res.RoleID,
		&res.Role,
		&res.CurrentScore,
		&res.CurrentGrade,
		&res.FinalScore,
		&res.FinalGrade,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		return res, err
	}

	return res, nil
}

func (r *pgRepository) FindFilter(ctx context.Context, filter entity.ReportEnrollment) ([]entity.ReportEnrollment, error) {
	results := []entity.ReportEnrollment{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"enrollment_id",
		"user_id",
		"login_id",
		"full_name",
		"role_id",
		"role",
		"current_score",
		"current_grade",
		"final_score",
		"final_grade",
		"created_at",
		"updated_at",
	).From(DBTABLE)

	nullValue := entity.ReportEnrollment{}
	if filter.ID != nullValue.ID {
		query = query.Where(sq.Eq{"id": filter.ID})
	}
	if filter.CourseReportID != nullValue.CourseReportID {
		query = query.Where(sq.Eq{"course_report_id": filter.CourseReportID})
	}
	if filter.EnrollmentID != nullValue.EnrollmentID {
		query = query.Where(sq.Eq{"enrollment_id": filter.EnrollmentID})
	}
	if filter.CreatedAt.Time != nullValue.CreatedAt.Time {
		query = query.Where("date(created_at) = ?", filter.CreatedAt.Time.Format("2006-01-02"))
	}
	if filter.UpdatedAt.Time != nullValue.UpdatedAt.Time {
		query = query.Where("date(updated_at) = ?", filter.UpdatedAt.Time.Format("2006-01-02"))
	}

	query = query.RunWith(r.con)

	rows, err := query.Query()
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		result := entity.ReportEnrollment{}
		err = rows.Scan(
			&result.ID,
			&result.CourseReportID,
			&result.EnrollmentID,
			&result.UserID,
			&result.LoginID,
			&result.FullName,
			&result.RoleID,
			&result.Role,
			&result.CurrentScore,
			&result.CurrentGrade,
			&result.FinalScore,
			&result.FinalGrade,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
