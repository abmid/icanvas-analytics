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
	"time"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"

	sq "github.com/Masterminds/squirrel"
)

type pgRepository struct {
	con *sql.DB
	sq  sq.StatementBuilderType
}

var (
	DBTABLE = "report_courses"
	POOL    = 2
)

func NewCoursePG(con *sql.DB) *pgRepository {
	return &pgRepository{
		con: con,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgRepository) Create(ctx context.Context, reportCourse *entity.ReportCourse) error {
	query := r.sq.Insert(DBTABLE).Columns(
		"course_id",
		"course_name",
		"account_id",
		"created_at",
		"updated_at",
	).Values(
		reportCourse.CourseID,
		reportCourse.CourseName,
		reportCourse.AccountID,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING \"id\"").RunWith(r.con)
	err := query.QueryRow().Scan(&reportCourse.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Read(ctx context.Context) ([]entity.ReportCourse, error) {
	results := []entity.ReportCourse{}
	query := r.sq.Select(
		"id",
		"course_id",
		"course_name",
		"account_id",
		"created_at",
		"updated_at",
		"deleted_at",
	).From(DBTABLE).RunWith(r.con)

	rows, err := query.Query()
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, err
	}
	for rows.Next() {
		result := entity.ReportCourse{}
		err = rows.Scan(
			&result.ID,
			&result.CourseID,
			&result.CourseName,
			&result.AccountID,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *pgRepository) FindFilter(ctx context.Context, filter entity.ReportCourse) ([]entity.ReportCourse, error) {
	results := []entity.ReportCourse{}
	query := r.sq.Select(
		"id",
		"course_id",
		"course_name",
		"account_id",
		"created_at",
		"updated_at",
		"deleted_at",
	).From(DBTABLE)

	nullValue := entity.ReportCourse{}
	if filter.ID != nullValue.ID {
		query = query.Where(sq.Eq{"id": filter.ID})
	}
	if filter.AccountID != nullValue.AccountID {
		query = query.Where(sq.Eq{"account_id": filter.AccountID})
	}
	if filter.CourseID != nullValue.CourseID {
		query = query.Where(sq.Eq{"course_id": filter.CourseID})
	}
	if filter.CourseName != nullValue.CourseName {
		query = query.Where("course_name like ?", fmt.Sprint("%", filter.CourseName, "%"))
	}
	if filter.CreatedAt.Time != nullValue.CreatedAt.Time {
		query = query.Where("date(created_at) = ? ", filter.CreatedAt.Time.Format("2006-01-02"))
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
	for rows.Next() {
		result := entity.ReportCourse{}
		err = rows.Scan(
			&result.ID,
			&result.CourseID,
			&result.CourseName,
			&result.AccountID,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *pgRepository) Update(ctx context.Context, reportCourse *entity.ReportCourse) error {
	query := r.sq.Update(DBTABLE).
		Set("course_id", reportCourse.CourseID).
		Set("course_name", reportCourse.CourseName).
		Set("account_id", reportCourse.AccountID).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": reportCourse.ID}).Suffix("RETURNING \"id\"").RunWith(r.con)
	err := query.QueryRow().Scan(&reportCourse.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Delete(ctx context.Context, reportCourseID uint32) error {
	query := r.sq.Delete(DBTABLE).Where(sq.Eq{"id": reportCourseID}).RunWith(r.con)
	_, err := query.Exec()
	if err != nil {
		return nil
	}
	return nil
}

func (r *pgRepository) FindByID(ctx context.Context, id uint32) (res entity.ReportCourse, err error) {
	query := r.sq.Select(
		"id",
		"course_id",
		"course_name",
		"account_id",
		"created_at",
		"updated_at",
		"deleted_at",
	).From(DBTABLE).Where(sq.Eq{"id": id}).RunWith(r.con)
	err = query.QueryRow().Scan(
		&res.ID,
		&res.CourseID,
		&res.CourseName,
		&res.AccountID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *pgRepository) FindByCourseIDDateNow(ctx context.Context, courseID uint32) (res entity.ReportCourse, err error) {
	query := r.sq.Select(
		"id",
		"course_id",
		"course_name",
		"account_id",
		"created_at",
		"updated_at",
		"deleted_at",
	).From(DBTABLE).
		Where(sq.Eq{"course_id": courseID}).
		Where("date(created_at) = current_date").
		Limit(1).
		RunWith(r.con)

	err = query.QueryRow().Scan(
		&res.ID,
		&res.CourseID,
		&res.CourseName,
		&res.AccountID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return res, err
	}

	return res, nil
}
