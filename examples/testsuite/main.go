package main

import (
	"context"
	"fmt"
	"github.com/lxzan/gws"
	"net/http"
	"time"
)

func main() {
	var upgrader = gws.Upgrader{CompressEnabled: true, MaxContentLength: 32 * 1024 * 1024}
	var handler = new(WebSocket)

	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = upgrader.Upgrade(context.Background(), writer, request, handler)
	})

	_ = http.ListenAndServe(":3000", nil)
}

type WebSocket struct{}

func (c *WebSocket) OnClose(socket *gws.Conn, message *gws.Message) {
	fmt.Printf("onclose: code=%d, payload=%s\n", message.Code(), string(message.Bytes()))
	_ = socket.Close()
	_ = message.Close()
}

func (c *WebSocket) OnError(socket *gws.Conn, err error) {
	fmt.Printf("onerror: err=%s\n", err.Error())
	_ = socket.Close()
}

func (c *WebSocket) OnOpen(socket *gws.Conn) {
	println("connected")
}

func (c *WebSocket) OnMessage(socket *gws.Conn, message *gws.Message) {
	socket.WriteMessage(message.Typ(), message.Bytes())
	_ = message.Close()
}

func (c *WebSocket) OnPing(socket *gws.Conn, message *gws.Message) {
	fmt.Printf("onping: payload=%s\n", string(message.Bytes()))
	socket.WritePong(message.Bytes())
	socket.SetDeadline(time.Now().Add(30 * time.Second))
	_ = message.Close()
}

func (c *WebSocket) OnPong(socket *gws.Conn, message *gws.Message) {}