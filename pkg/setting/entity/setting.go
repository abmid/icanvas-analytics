/*
 * File Created: Thursday, 16th July 2020 3:16:53 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

import "database/sql"

type Setting struct {
	ID        uint32 `query:"id"`
	Name      string `query:"name"`
	Category  string `query:"category"`
	Value     string `query:"value"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
