package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行なっている1人のユーザを表す
type client struct {
	socket *websocket.Conn // socketはクライアントのためのWebSocket
	send   chan []byte     // sendはメッセージが送られるチャネル
	room   *room           // roomはクライアントが参加しているチャットルーム
}

// socketからメッセージを読み込み、受け取ったメッセージをroomのforwardチャネルに送信する
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// sendチャネルに溜まったメッセージをsocketへ書き込む
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
