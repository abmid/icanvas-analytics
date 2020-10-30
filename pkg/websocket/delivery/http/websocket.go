/*
 * File Created: Monday, 21st September 2020 12:36:12 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package http

import (
	"net/http"

	"github.com/abmid/icanvas-analytics/pkg/websocket/usecase"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WSHandler struct {
	WSUseCase usecase.WebsocketUseCase
}

type WebSocketConnection struct {
	*websocket.Conn
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func (handler *WSHandler) WebSocketServer() echo.HandlerFunc {
	return func(c echo.Context) error {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
		if err != nil {
			return err
		}
		defer ws.Close()

		handler.WSUseCase.Echo(ws)

		return nil
	}
}

func NewHandler(path string, g *echo.Group, JWTKey string, wsUC usecase.WebsocketUseCase) {

	handler := WSHandler{
		WSUseCase: wsUC,
	}
	r := g.Group(path)
	r.GET("/server", handler.WebSocketServer())
}
