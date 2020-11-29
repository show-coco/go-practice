package main

type room struct {
	forward chan []byte // forwardは他のクライアントに転送するためのメッセージを保存するチャネル
}
