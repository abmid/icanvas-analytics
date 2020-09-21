/*
 * File Created: Saturday, 6th June 2020 10:53:06 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Your Company
 */

package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/report/entity"

	sq "github.com/Masterminds/squirrel"
)

var (
	DBTABLE = "report_users"
	POOL    = 2
)

type pgRepository struct {
	con *sql.DB
	sq  sq.StatementBuilderType
	Log *logger.LoggerWrap
}

func NewUserPG(con *sql.DB) *pgRepository {

	logger := logger.New()
	return &pgRepository{
		con: con,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		Log: logger,
	}
}

func (r *pgRepository) Create(ctx context.Context, reportUser *entity.ReportUser) error {

	query := r.sq.Insert(DBTABLE).
		Columns(
			"report_course_id",
			"user_id",
			"login_id",
			"full_name",
			"created_at",
			"updated_at",
		).Values(
		reportUser.ReportCourseID,
		reportUser.UserID,
		reportUser.LoginID,
		reportUser.FullName,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING \"id\"").RunWith(r.con)

	err := query.QueryRow().Scan(&reportUser.ID)

	if err != nil {
		r.Log.Error(err)
		return err
	}

	return nil
}

func (r *pgRepository) Update(ctx context.Context, reportUser *entity.ReportUser) error {
	query := r.sq.Update(DBTABLE).
		Set("report_course_id", reportUser.ReportCourseID).
		Set("user_id", reportUser.UserID).
		Set("login_id", reportUser.LoginID).
		Set("full_name", reportUser.FullName).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": reportUser.ID}).Suffix("RETURNING \"id\"").RunWith(r.con)

	err := query.QueryRow().Scan(&reportUser.ID)
	if err != nil {
		r.Log.Error(err)
		return err
	}
	return nil
}

func (r *pgRepository) CreateOrUpdateByCourseReportID(ctx context.Context, reportUser *entity.ReportUser) error {
	query := r.sq.Select("id").From(DBTABLE).Where(sq.Eq{"report_course_id": reportUser.ReportCourseID}).RunWith(r.con)

	err := query.QueryRow().Scan(&reportUser.ID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	// IF Exist
	if reportUser.ID != 0 {
		err := r.Update(ctx, reportUser)
		if err != nil {
			r.Log.Error(err)
			return err
		}
		return nil
	}
	// If Not exist
	err = r.Create(ctx, reportUser)
	if err != nil {
		r.Log.Error(err)
		return err
	}
	return nil
}
