/*
 * File Created: Thursday, 16th July 2020 3:16:32 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package repository

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/abmid/icanvas-analytics/pkg/setting/entity"
)

type pgRepository struct {
	con *sql.DB
	sq  sq.StatementBuilderType
}

var (
	DBNAME = "settings"
)

func NewRepositoryPG(db *sql.DB) *pgRepository {
	return &pgRepository{
		con: db,
		sq:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *pgRepository) Create(setting *entity.Setting) error {
	query := r.sq.Insert(DBNAME).Columns("name", "category", "value", "created_at", "updated_at").Values(
		setting.Name,
		setting.Category,
		setting.Value,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING \"id\"").RunWith(r.con)

	err := query.QueryRow().Scan(&setting.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) Update(id uint32, setting entity.Setting) error {
	query := r.sq.Update(DBNAME).
		Set("name", setting.Name).
		Set("category", setting.Category).
		Set("value", setting.Value).
		Set("updated_at", time.Now()).Where(sq.Eq{"id": id}).RunWith(r.con)

	_, err := query.Exec()
	if err != nil {
		return err
	}

	return nil

}

func (r *pgRepository) FindByFilter(filter entity.Setting) (res []entity.Setting, err error) {
	nullFilter := entity.Setting{}
	query := r.sq.Select("id", "name", "category", "value", "created_at", "updated_at").From(DBNAME)

	// Check if id not null
	if filter.ID != nullFilter.ID {
		query = query.Where(sq.Eq{"id": filter.ID})
	}

	// Check if name not null
	if filter.Name != nullFilter.Name {
		query = query.Where(sq.Eq{"name": filter.Name})
	}

	// Check if category not null
	if filter.Category != nullFilter.Category {
		query = query.Where(sq.Eq{"category": filter.Category})
	}

	query = query.RunWith(r.con)

	rows, err := query.Query()
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		setting := entity.Setting{}
		err = rows.Scan(
			&setting.ID,
			&setting.Name,
			&setting.Category,
			&setting.Value,
			&setting.CreatedAt,
			&setting.UpdatedAt,
		)

		res = append(res, setting)
	}

	return res, nil
}
