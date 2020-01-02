# 11_helloworld_on_gae_go
## purpose
- HelloWorldを表示する
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