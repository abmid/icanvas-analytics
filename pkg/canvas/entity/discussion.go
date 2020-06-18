/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package entity

type Discussion struct {
	ID        uint32 `json:"id"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
}
