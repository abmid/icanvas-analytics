/*
 * File Created: Monday, 21st September 2020 12:40:48 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"fmt"
	"time"

	"github.com/abmid/icanvas-analytics/internal/logger"
	entity "github.com/abmid/icanvas-analytics/pkg/websocket/entity"
	"github.com/gorilla/websocket"
)

type msg struct {
	Message string
}

type websocketUseCase struct {
	Log *logger.LoggerWrap
}

type WebSocketConnection struct {
	*websocket.Conn
}

var connections = make([]*WebSocketConnection, 0)

func New() *websocketUseCase {
	logger := logger.New()

	return &websocketUseCase{
		Log: logger,
	}
}

func broadcastMessage(currentConn *WebSocketConnection, message entity.Notification) {
	for _, connection := range connections {
		// Except for connection sender
		if connection == currentConn {
			continue
		}
		connection.WriteJSON(message)
	}
}

func (useCase *websocketUseCase) SendMessageToAll(message string) {
	data := entity.Notification{
		Message:   message,
		CreatedAt: time.Now(),
	}
	// test := entity.Notifi
	for _, connection := range connections {
		connection.WriteJSON(data)
	}
}

func (useCase *websocketUseCase) Echo(conn *websocket.Conn) {

	currentConn := WebSocketConnection{Conn: conn}
	connections = append(connections, &currentConn)

	for {
		m := entity.Notification{}
		err := conn.ReadJSON(&m)
		if err != nil {
			useCase.Log.Error(err)
			continue
		}

		broadcastMessage(&currentConn, m)
		fmt.Printf("%s\n", m)
	}
}
