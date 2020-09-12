# 解説
## これはなに？
HTMLで画像を表示するサービス
## 何を学ぶのか？
HTMLの基礎
静的ファイルの取り扱い
## 前回との差分
- 追記 HTMLの中に<img src="kmc.png" style="height:100px;width:auto;">と書かれている

- ファイル staticディレクトリが作られ、中にkmc.pngがはいっている

- 追記 app.yamlの中に
```
- url: /(.*\..*)
  static_files: static/\1
  upload: static/.*
```

## 教える知識
HTMLのIMGタグは画像を表示できる  
app.yaml内の
```
- url: /(.*\..*)
  static_files: static/\1
  upload: static/.*
```
という文言は拡張子をもつURLをstaticディレクトリ内のファイルに割り当てている
/(.*\..*)は正規表現といい拡張子を持つURLと合致する。正規表現は下から自分で調べてください。
app.yamlの文法を読めば全体の意味を理解できます。app.yamlの文法は下から自分で調べてください
## 自分で調べる知識
* [HTML入門](http://www.ink.or.jp/~bigblock/html/) ここに<img>についての解説もある
* [正規表現](https://ja.wikipedia.org/wiki/%E6%AD%A3%E8%A6%8F%E8%A1%A8%E7%8F%BE)
* [app.yaml](https://cloud.google.com/appengine/docs/standard/python/config/appref?hl=ja)
