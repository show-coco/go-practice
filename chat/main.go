package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templは1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理します。複数のgoroutineから呼び出されても同じServeHTTPを使用するのでレシーバはポインタ
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		// コンパイルされたテンプレートへの参照を保持
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// チャットルームを開始
	go r.run()

	// Webサーバを起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
