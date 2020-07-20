package utils

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

var Server *socketio.Server
var wserr error
var WebSocket *socketio.Conn

func InitWs() {
	Server, wserr = socketio.NewServer(nil)
	if wserr != nil {
		log.Fatal(wserr)
	}

	Server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	Server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
		WebSocket = &s
	})

	go Server.Serve()
}
