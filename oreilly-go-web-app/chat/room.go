package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/show-coco/oreilly-go-web-app/trace"
)

type room struct {
	forward chan []byte  // 他のクライアントに転送するためのメッセージを保存するチャネル
	join    chan *client // チャットルームに参加しようとしているクライアントのためのチャネル
	leave   chan *client // チャットルームから退出しようとしているクライアントのためのチャネル
	clients map[*client]bool
	tracer  trace.Tracer
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("新しいクライアントが参加しました")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("クラインアントが退室しました")
		case msg := <-r.forward:
			r.tracer.Trace("メッセージを受信しました: ", string(msg))
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージを送信
					r.tracer.Trace(" -- クライアントに送信されました")
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- 送信に失敗しました。クライアントをクリーンアップします")
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// HTTP接続をアップグレードするためにwebsocketのupgraderを用意
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read() // readメソッドを呼び出し、接続を保持
}
