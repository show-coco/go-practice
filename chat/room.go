package main

type room struct {
	forward chan []byte  // 他のクライアントに転送するためのメッセージを保存するチャネル
	join    chan *client // チャットルームに参加しようとしているクライアントのためのチャネル
	leave   chan *client // チャットルームから退出しようとしているクライアントのためのチャネル
	clients map[*client]bool
}
