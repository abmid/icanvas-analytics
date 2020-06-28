/*
 * File Created: Thursday, 4th June 2020 3:01:02 pm
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
	DBTABLE = "report_course_results"
	POOL    = 2
)

func NewResultPG(con *sql.DB) *pgRepository {
	return &pgRepository{
		con: con,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgRepository) Create(ctx context.Context, result *entity.ReportResult) error {
	query := r.sq.Insert(DBTABLE).Columns(
		"report_course_id",
		"assigment_count",
		"discussion_count",
		"student_count",
		"finish_grading_count",
		"final_score",
		"created_at",
		"updated_at",
	).Values(
		result.ReportCourseID,
		result.AssigmentCount,
		result.DiscussionCount,
		result.StudentCount,
		result.FinishGradingCount,
		result.FinalScore,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING \"id\"").RunWith(r.con)
	err := query.QueryRow().Scan(&result.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Update(ctx context.Context, result *entity.ReportResult) error {
	query := r.sq.Update(DBTABLE).
		Set("report_course_id", result.ReportCourseID).
		Set("assigment_count", result.AssigmentCount).
		Set("discussion_count", result.DiscussionCount).
		Set("student_count", result.StudentCount).
		Set("finish_grading_count", result.FinishGradingCount).
		Set("final_score", result.FinalScore).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": result.ID}).
		Suffix("RETURNING \"id\"").RunWith(r.con)

	err := query.QueryRow().Scan(&result.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) CreateOrUpdateByCourseReportID(ctx context.Context, result *entity.ReportResult) error {
	var findCourseAssigmentID uint32
	query := r.sq.Select("id").
		From(DBTABLE).
		Where(sq.Eq{"report_course_id": result.ReportCourseID}).
		Limit(1).
		RunWith(r.con)

	err := query.QueryRow().Scan(&findCourseAssigmentID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	result.ID = findCourseAssigmentID
	if findCourseAssigmentID == 0 {
		err = r.Create(ctx, result)
		if err != nil {
			return err
		}
		return nil
	}
	err = r.Update(ctx, result)
	if err != nil {
		return err
	}
	return nil
}
