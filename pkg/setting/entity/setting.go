/*
 * File Created: Thursday, 16th July 2020 3:16:53 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

import "database/sql"

type Setting struct {
	ID        uint32       `query:"id" json:"id"`
	Name      string       `query:"name" json:"name"`
	Category  string       `query:"category" json:"category"`
	Value     string       `query:"value" json:"value"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
