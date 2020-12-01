# インタフェース

```go
type Tracer interface {
	Trace(...interface{})
}
```

`...interface{}`という引数の型は、任意の型を何個でも受け取れる。`fmtSprint`や`log.Fatal`でもこのような型が使われている

# テスト

名前が Test で始まり、`*testing.T`型の引数を受け取る関数はユニットテストとみなされる

```bash
go test
```
