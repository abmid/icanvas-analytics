/*
 * File Created: Thursday, 18th June 2020 4:59:24 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import (
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/abmid/icanvas-analytics/internal/logger"
	"github.com/abmid/icanvas-analytics/pkg/user/entity"
)

type repositoryPG struct {
	DB  *sql.DB
	sq  squirrel.StatementBuilderType
	Log *logger.LoggerWrap
}

const DB_NAME = "users"

func NewPG(db *sql.DB) *repositoryPG {

	logger := logger.New()

	return &repositoryPG{
		DB:  db,
		sq:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		Log: logger,
	}
}

func (r *repositoryPG) Create(user *entity.User) error {

	query := r.sq.Insert(DB_NAME).Columns(
		"name",
		"email",
		"password",
		"created_at",
		"updated_at",
	).Values(
		user.Name,
		user.Email,
		user.Password,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING \"id\"").RunWith(r.DB)

	err := query.QueryRow().Scan(&user.ID)
	if err != nil {
		r.Log.Error(err)
		return err
	}

	return nil
}

func (r *repositoryPG) All() (res []entity.User, err error) {

	query := r.sq.Select("id", "name", "email", "password", "created_at", "updated_at").From(DB_NAME).RunWith(r.DB)

	rows, err := query.Query()
	if err != nil {
		r.Log.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			r.Log.Error(err)
			return nil, err
		}
		res = append(res, user)
	}

	return res, nil
}

func (r *repositoryPG) Find(email string) (res *entity.User, err error) {

	query := r.sq.Select("id", "name", "email", "password", "created_at", "updated_at").
		From(DB_NAME).
		Where(squirrel.Eq{"email": email}).
		Limit(1).
		RunWith(r.DB)

	user := entity.User{}
	err = query.QueryRow().Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.Log.Error(err)
		return nil, err
	}
	res = &user

	return res, nil
}
