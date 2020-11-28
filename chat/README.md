## http

### http.Handler

ServeHTTP 関数を持つインタフェース。ServeHTTP は HTTP リクエストを受け取りレスポンスを返すことが責務。

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```

### http.Handle

URL と URL に対応する`http.Handler`を`DefaultServeMux`に登録する関数。

```go
func Handle(pattern string, handler Handler) {
  DefaultServeMux.Handle(pattern, handler)
}
```

`DefaultServeMux`は URL に対応した`http.Handler`を実行するルータ。つまり、`localhost:8080/`と処理 A を`DefaultServeMux`に登録することによって、`localhost:8080/`にリクエストがきた時には処理 A が実行される。

```go
http.Handle("/any/", anyHandler)
```

### http.HandlerFunc

func(ResponseWriter, \*Request) の別名の型で ServeHTTP 関数を持つので、関数を定義して http.HandlerFunc にキャストするだけで構造体を宣言することなく http.Handler を用意することができる。

HandlerFunc を使わない場合、以下のように構造体を宣言し、ServeHTTP を実装しなければならない。

```go
type MyHandler struct{}
func (handler MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  /* 処理 */
}

myHandler := &MyHandler{}
http.Handle("/any/", myHandler)
```

`http.HandlerFunc`にキャストすれば構造体を宣言しなくて良い。

```go
myHandler := http.HandlerFunc(
  func(w http.ResponseWriter, r *http.Request) {
    /* 処理 */
  })

http.Handle("/any/", myHandler)
```

### http.HandleFunc

`http.Handle`の使いやすいバージョン。内部で`http.HandleFunc`にキャストしてくれるため、関数を渡すだけで良い。

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
```

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { /* 処理 */ })
```

参考:

- [Go 言語の http パッケージにある Handle とか Handler とか HandleFunc とか HandlerFunc とかよくわからないままとりあえずイディオムとして使ってたのでちゃんと理解したメモ](https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65)
- [http パッケージ](http://golang.jp/pkg/http)
