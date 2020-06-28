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
	DBTABLE = "report_assigments"
	POOL    = 2
)

func NewAssigmentPG(con *sql.DB) *pgRepository {
	return &pgRepository{
		con: con,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgRepository) Create(ctx context.Context, assigment *entity.ReportAssigment) error {
	query := r.sq.Insert(DBTABLE).Columns(
		"course_report_id",
		"assigment_id",
		"name",
		"created_at",
		"updated_at",
	).Values(
		assigment.CourseReportID,
		assigment.AssigmentID,
		assigment.Name,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING \"id\"").RunWith(r.con)
	err := query.QueryRow().Scan(&assigment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Read(ctx context.Context) ([]entity.ReportAssigment, error) {
	results := []entity.ReportAssigment{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"assigment_id",
		"name",
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
	for rows.Next() {
		result := entity.ReportAssigment{}
		err = rows.Scan(
			&result.ID,
			&result.CourseReportID,
			&result.AssigmentID,
			&result.Name,
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

func (r *pgRepository) Update(ctx context.Context, assigment *entity.ReportAssigment) error {
	query := r.sq.Update(DBTABLE).
		Set("course_report_id", assigment.CourseReportID).
		Set("assigment_id", assigment.AssigmentID).
		Set("name", assigment.Name).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": assigment.ID}).RunWith(r.con)
	_, err := query.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Delete(ctx context.Context, reportAssigmentID uint32) error {
	query := r.sq.Delete(DBTABLE).Where(sq.Eq{"id": reportAssigmentID}).RunWith(r.con)
	_, err := query.Exec()
	if err != nil {
		return nil
	}
	return nil
}

func (r *pgRepository) FindFilter(ctx context.Context, filter entity.ReportAssigment) (res []entity.ReportAssigment, err error) {
	nullValue := entity.ReportAssigment{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"assigment_id",
		"name",
		"created_at",
		"updated_at").From(DBTABLE)

	// TODO : Filter
	if filter.ID != nullValue.ID {
		query = query.Where(sq.Eq{"id": filter.ID})
	}
	if filter.AssigmentID != nullValue.AssigmentID {
		query = query.Where(sq.Eq{"assigment_id": filter.AssigmentID})
	}
	if filter.CourseReportID != nullValue.CourseReportID {
		query = query.Where(sq.Eq{"course_report_id": filter.CourseReportID})
	}
	if filter.Name != nullValue.Name {
		query = query.Where(sq.Eq{"name": filter.Name})
	}
	if filter.CreatedAt.Time != nullValue.CreatedAt.Time {
		query = query.Where("date(created_at) = ? ", filter.CreatedAt.Time.Format("2006-01-02"))
	}
	if filter.UpdatedAt.Time != nullValue.UpdatedAt.Time {
		query = query.Where("date(updated_at) = ? ", filter.UpdatedAt.Time.Format("2006-01-02"))
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
		result := entity.ReportAssigment{}
		err = rows.Scan(
			&result.ID,
			&result.CourseReportID,
			&result.AssigmentID,
			&result.Name,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, result)
	}

	return res, nil
}

func (r *pgRepository) FindFirstByFilter(ctx context.Context, filter entity.ReportAssigment) (res entity.ReportAssigment, err error) {
	nullValue := entity.ReportAssigment{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"assigment_id",
		"name",
		"created_at",
		"updated_at").From(DBTABLE)

	// TODO : Filter
	if filter.ID != nullValue.ID {
		query = query.Where(sq.Eq{"id": filter.ID})
	}
	if filter.AssigmentID != nullValue.AssigmentID {
		query = query.Where(sq.Eq{"assigment_id": filter.AssigmentID})
	}
	if filter.CourseReportID != nullValue.CourseReportID {
		query = query.Where(sq.Eq{"course_report_id": filter.CourseReportID})
	}
	if filter.Name != nullValue.Name {
		query = query.Where(sq.Eq{"name": filter.Name})
	}
	// RUN find First
	query = query.Limit(1).RunWith(r.con)

	err = query.QueryRow().Scan(
		&res.ID,
		&res.CourseReportID,
		&res.AssigmentID,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		return res, err
	}

	return res, nil
}

func (r *pgRepository) CreateOrUpdateByCourseReportID(ctx context.Context, assigment *entity.ReportAssigment) error {
	var findCourseAssigmentID uint32
	query := r.sq.Select("id").
		From(DBTABLE).
		Where(sq.Eq{"course_report_id": assigment.CourseReportID}).
		Limit(1).
		RunWith(r.con)

	err := query.QueryRow().Scan(&findCourseAssigmentID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	assigment.ID = findCourseAssigmentID
	if findCourseAssigmentID == 0 {
		err = r.Create(ctx, assigment)
		if err != nil {
			return err
		}
	}
	return nil
}
