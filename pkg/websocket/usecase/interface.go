/*
 * File Created: Monday, 21st September 2020 12:41:13 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import "github.com/gorilla/websocket"

type WebsocketUseCase interface {
	Echo(conn *websocket.Conn)
	SendMessageToAll(message string)
}
