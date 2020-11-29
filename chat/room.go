package main

type room struct {
	forward chan []byte  // 他のクライアントに転送するためのメッセージを保存するチャネル
	join    chan *client // チャットルームに参加しようとしているクライアントのためのチャネル
	leave   chan *client // チャットルームから退出しようとしているクライアントのためのチャネル
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージを送信
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
