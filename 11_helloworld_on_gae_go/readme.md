# 11_helloworld_on_gae_go
## 11.1 目的

- Hello Worldを表示する
- 基本的なウェブアプリケーションの仕組みを理解する

## 11.2 基本的なウェブアプリケーションの仕組み

ウェブアプリケーション - Wikipedia から引用

>代表的なウェブアプリケーションでは、WebブラウザがHTTPを利用してHTMLを取得・表示、それをDOMを介してJavaScriptが操作し、必要に応じてWebサーバと通信をおこなってデータを更新する。このようにウェブ（World Wide Web）を基盤として作られる応用ソフトウェアをウェブアプリケーション（Webアプリ）と総称する。上記の例はあくまでウェブアプリケーションを実現する技術スタックの一例であり、他の様々な技術を用いてWebアプリを作成できる。またウェブアプリケーションの明確な定義は存在しない（動的なウェブページとの差異は不明瞭である）。
>ウェブアプリケーションの一例としては、ウィキペディアなどで使われているウィキやブログ、電子掲示板、銀行のインターネットバンキング、証券会社のオンライントレード、電子商店街などネット販売のショッピングカートなどを挙げることができる。
>ウェブアプリケーションに対して、ローカルのデスクトップ環境上で動作するアプリケーションは、デスクトップアプリケーションやスタンドアロンアプリケーション、スマートフォンで動作するアプリケーションはネイティブアプリと呼ばれる。
>Webアプリはクライアント-サーバーモデルを基本としており、WWWを基盤とする分散コンピューティングの一形態ともみれる。2010年代後半には多数のマイクロサービスをAPIを介して連携させ構成されるWebアプリも増えており、分散コンピューティングとしての側面がより強くなっている。

### 11.2.1 最低限の仕組み

HTTPプロトコルに従いインターネットからリクエストと呼ばれる情報の取得の要求を受け付けレスポンスと呼ばれる返答を返すプログラム

### 11.2.2 リクエストとレスポンス

以下のコマンドを実行すると下記の生のリクエストとレスポンスを含む文章が表示される


```
$curl --http1.1 --get -v https://www.kmc.gr.jp/
（中略、以下リクエスト）
 GET / HTTP/1.1
 Host: qiita.com
 User-Agent: curl/7.55.1
 Accept: */*
 
 (中略、以下レスポンス)
 HTTP/1.1 200 OK
 Server: nginx/1.17.6
 Date: Thu, 02 Jan 2020 15:05:02 GMT
 Content-Type: text/html
 Content-Length: 10051
 Last-Modified: Thu, 02 Jan 2020 04:29:12 GMT
 Connection: keep-alive
 Vary: Accept-Encoding
 ETag: "5e0d7198-2743"
 Strict-Transport-Security: max-age=15768000
 Accept-Ranges: bytes
 
<!DOCTYPE html>
<html lang='ja' data-path='index'>
  <head>
    <meta charset='utf-8'>
    <meta name='viewport' content='width=device-width,initial-scale=1'>
    <link rel="stylesheet" href="/assets/kmc-d37c216ea54258989bb2b5cfcba505b7.css">
(以下HTML)
```

これが生のHTTPリクエストとレスポンスである。

リクエストから解説する

- `GET`は(メソッド|method)という情報であり情報に対する処理を意味する。
    - `GET`は取得を意味する
    - 他にも情報を書き込む`POST`や情報を削除する`DELETE`などがある
- `/` は(相対パス|path|URI)という情報であり情報の場所を意味する
    - `/` は最上位の場所を意味する。
    - 例えば`/a/b/c`はaの中のbの中のcという場所を意味する
- `HTTP/1.1` はプロトコルとプロトコルバージョンの宣言である。
    - 現在のブラウザはほぼ`HTTP/1.1`で通信している
- `Host: www.kmc.gr.jp` は(ホスト|host)という情報でありIPアドレスかドメインで接続先コンピュータを指定する。
    - `www.kmc.gr.jp`は接続先のコンピューターを指定するドメインである。
    - ドメインに対応するIPアドレスは `$nslookup www.kmc.gr.jp` コマンドで表示できる
- 他はリクエストのHTTPヘッダである。
    - ` User-Agent: curl/7.55.1`はリクエストを発行したプログラムを表す。
    - ` Accept: */*` は要求するレスポンスの形式を表す。
    - 他にも様々な種類のヘッダが存在する。

レスポンスを解説する

- `HTTP/1.1` はレスポンスと同様にプロトコルとそのバージョン
- `200 OK` はステータスと呼ばれてレスポンスの概要をステータスコードと呼ばれる番号で表す
    - 200から299までは成功を意味する
        - 200 OK : リクエストは成功しレスポンスとともに要求に応じた情報が返される。
    - 400から499まではクライアントエラーを意味する
        - 400 Bad Request : リクエストが不正である。
        - 404 Not Found : 未検出。リソースが見つからなかった。
    - 500から599 まではサーバーエラーを意味する
- 続きはレスポンスのヘッダであり様々な種類がある
    - Date: Thu, 02 Jan 2020 15:11:54 GMT はサーバーで観測したアクセス時刻
- ヘッダーの後一行空けてレスポンスの内容が格納されている
    - 今回はトップページを要求しているのでHTMLを返している

インターネットでリクエストを受け付けレスポンスを送るのがウェブサーバーである。


## files
- .gcloudignore
  - gitやIDE関連の不要なファイルをサーバーに送らない設定ファイル
  - 無くても動くが遅い
- app.yaml
  - AppEngineの設定ファイル
- index.html
  - htmlファイル
- main.go
  - メインソースコード
- readme.md
  - このファイル
- server.go
  - main.goを短縮するためのライブラリ
  - サーバー起動処理、型の別名付与、ユーティリティ関数をまとめた
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