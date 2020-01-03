# 11_helloworld_on_gae_go
## 目的

- Hello Worldを表示する
- 基本的なウェブアプリケーションの仕組みを理解する

## main.go, server.go

Goの仕様として同一ディレクトリ直下のソースコードは同一パッケージに所属する必要がある。
実際このディレクトリのソースコードは `main.go` `server.go` の二つであり共に`package main`で始まる。

Goの仕様として最初に実行される関数は`package main`の`func main`つまり`main.main()`である。

`main.go`は以下のように定義されている。
`index.html`をレスポンスに書き込んでいるだけである。

```go
package main

func main() {
	Handle("/", mainHandler)
	Listen(map[string]string{})
}
func mainHandler(w Response, r Request) {
	Writef(w, nil, "index.html")
}
```

これは以下のソースコードと等価である。
server.goの`Handle()`はパスと関数の対応を行う`http.HandleFunc()`、
`Listen()`は接続待機を行う`http.ListenAndServe()`と環境変数周辺を、
ラップし隠蔽している。

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("PORT")==""{
		os.Setenv("PORT","8080")
	}
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	b,_:=ioutil.ReadFile("index.html")
	w.Write(b)
}
```

server.goは他にも今後登場する様々な主要機能を実装した。

## app.yaml

app.yamlはAppEngine対する設定ファイルである。
AppEngineにデプロイする際に参照される。
ローカルホストで開発している際は参照されない。

以下が`app.yaml`の中身である。
goのバージョン1.12で動作させることと
`/`以下のパスのリクエスト（つまり全てのリクエスト）は
開発したモジュールでハンドルすることを定めている。


```yaml
runtime: go112

handlers:
- url: /
  script: auto
```

今回は基本だけであるが、
特定の条件を満たしたパスを別のモジュールに割り当てたり、
直接静的ファイルを配信したりと様々な設定ができる

## download
- go get -u github.com/lzpel/webservice2019
  - ダウンロードコマンド
  - 以下のディレクトリに配置される
  - $HOME\go\src\github.com\lzpel\webservice2019\11_helloworld_on_gae_go
## build
- Goland
  - ビルド設定
    - Edit Configurations
    - Add New Configuration/Go build
      - Name = app
      - Run kind = Directory
      - Directory and Working Directory = ～/13_go_client_for_cloud_storage
  - Run 'app' / Debug 'app'
- shell
  - windows：go build -i -o a.exe . ; a.exe
- tips
  - 管理者権限を求めるダイアログが出ればOKする
    - 外部と通信するのでウィルスと疑われやすい
## deploy
- gcloud app deploy app.yaml を実行するだけ
  - プロジェクトを環境変数GOROOT内に配置すること
    - 配置例
      - ~/go/src/github.com/lzpel/webservice2019/12_go_client_for_cloud_datastore
    - さもなくばエラー発生
      - failed to build app: Your app is not on your GOPATH, please move it there and try again.
  - 環境によって発生するエラーとその対処法
    - gcloud cannot find package "golang.org/x/sys/unix"
      - go get golang.org/x/sys/unix を実行すればいい
        - https://github.com/GoogleCloudPlatform/golang-samples/issues/590