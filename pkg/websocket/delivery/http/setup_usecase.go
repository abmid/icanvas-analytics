/*
 * File Created: Monday, 21st September 2020 12:59:40 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package http

import "github.com/abmid/icanvas-analytics/pkg/websocket/usecase"

func SetupUseCase() usecase.WebsocketUseCase {
	UC := usecase.New()
	return UC
}
