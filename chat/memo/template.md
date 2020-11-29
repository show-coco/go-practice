# tmplate

Go の標準ライブラリにはテンプレートには`text/template`と`html/template`がある。

`html/template`はコンテキストを認識するので、不正なスクリプトを埋め込む攻撃を回避したり、URL で使用できない文字をエンコードすることができる。
