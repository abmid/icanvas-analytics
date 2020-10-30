/*
 * File Created: Monday, 28th September 2020 3:58:24 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

import "time"

type Notification struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
