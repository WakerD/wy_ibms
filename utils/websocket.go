package utils

import (

	"fmt"

	"github.com/gorilla/websocket"
	"net/http"
)

type Websocket struct {}

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

var WsConn *websocket.Conn

func (ws Websocket) Wshandler(w http.ResponseWriter, r *http.Request) {
    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    WsConn = conn

    for {
        t, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        conn.WriteMessage(t, msg)
    }
}

func GetWs() *websocket.Conn{
	return WsConn
}