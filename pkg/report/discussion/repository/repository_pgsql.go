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
	DBTABLE = "report_discussions"
	POOL    = 2
)

func NewDiscussionPG(con *sql.DB) *pgRepository {
	return &pgRepository{
		con: con,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgRepository) Create(ctx context.Context, reportDiss *entity.ReportDiscussion) error {
	query := r.sq.Insert(DBTABLE).Columns(
		"course_report_id",
		"discussion_id",
		"title",
		"created_at",
		"updated_at",
	).Values(
		reportDiss.CourseReportID,
		reportDiss.DiscussionID,
		reportDiss.Title,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING \"id\"").RunWith(r.con)
	err := query.QueryRow().Scan(&reportDiss.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Read(ctx context.Context) ([]entity.ReportDiscussion, error) {
	results := []entity.ReportDiscussion{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"discussion_id",
		"title",
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
		result := entity.ReportDiscussion{}
		err = rows.Scan(
			&result.ID,
			&result.CourseReportID,
			&result.DiscussionID,
			&result.Title,
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

func (r *pgRepository) Update(ctx context.Context, reportDiss *entity.ReportDiscussion) error {
	query := r.sq.Update(DBTABLE).
		Set("course_report_id", reportDiss.CourseReportID).
		Set("discussion_id", reportDiss.DiscussionID).
		Set("title", reportDiss.Title).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": reportDiss.ID}).Suffix("RETURNING \"id\"").RunWith(r.con)
	err := query.QueryRow().Scan(&reportDiss.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Delete(ctx context.Context, reportDissID uint32) error {
	query := r.sq.Delete(DBTABLE).Where(sq.Eq{"id": reportDissID}).RunWith(r.con)
	_, err := query.Exec()
	if err != nil {
		return nil
	}
	return nil
}

func (r *pgRepository) FindFirstByFilter(ctx context.Context, filter entity.ReportDiscussion) (res entity.ReportDiscussion, err error) {
	nullValue := entity.ReportDiscussion{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"discussion_id",
		"title",
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
	if filter.DiscussionID != nullValue.DiscussionID {
		query = query.Where(sq.Eq{"discussion_id": filter.DiscussionID})
	}
	if filter.Title != nullValue.Title {
		query = query.Where(sq.Eq{"title": filter.Title})
	}
	// RUN find First
	query = query.Limit(1).RunWith(r.con)

	err = query.QueryRow().Scan(
		&res.ID,
		&res.CourseReportID,
		&res.DiscussionID,
		&res.Title,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		return res, err
	}

	return res, nil
}

func (r *pgRepository) FindFilter(ctx context.Context, filter entity.ReportDiscussion) ([]entity.ReportDiscussion, error) {
	results := []entity.ReportDiscussion{}
	query := r.sq.Select(
		"id",
		"course_report_id",
		"discussion_id",
		"title",
		"created_at",
		"updated_at",
	).From(DBTABLE)

	nullValue := entity.ReportDiscussion{}
	if filter.ID != nullValue.ID {
		query = query.Where(sq.Eq{"id": filter.ID})
	}
	if filter.CourseReportID != nullValue.CourseReportID {
		query = query.Where(sq.Eq{"course_report_id": filter.CourseReportID})
	}
	if filter.DiscussionID != nullValue.DiscussionID {
		query = query.Where(sq.Eq{"discussion_id": filter.DiscussionID})
	}
	if filter.Title != nullValue.Title {
		query = query.Where("title like ?", fmt.Sprint("%", filter.Title, "%"))
	}
	if filter.CreatedAt.Time != nullValue.CreatedAt.Time {
		query = query.Where("date(created_at) = ?", filter.CreatedAt.Time.Format("2006-01-02"))
	}
	if filter.UpdatedAt.Time != nullValue.UpdatedAt.Time {
		query = query.Where("date(updated_at)", filter.UpdatedAt.Time.Format("2006-01-02"))
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
		result := entity.ReportDiscussion{}
		err = rows.Scan(
			&result.ID,
			&result.CourseReportID,
			&result.DiscussionID,
			&result.Title,
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
