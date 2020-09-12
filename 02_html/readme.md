# 解説
## これはなに？
画面にhtmlを表示するサービス
## 何を学ぶのか？
HTMLの基礎
HTMLを表示する仕組み
## 前回との差分
>追加
from google.appengine.ext.webapp import template

>変更  
return webapp2.Response("hello")  
の代わりに  
return webapp2.Response(template.render("index.html", None))
と書かれている
 
>ファイル変更  
index.htmlという新しいファイルがある

## 教える知識
from google.appengine.ext.webapp import template でtemplateというライブラリを導入している  
templateによってindex.htmlが読み込まれてその内容が"Hello"の部分に展開される

ここではHTMLをそのまま表示しているがテンプレートというライブラリには動的にHTMLの内容の一部を変更する機能がある。今後使用する
## 自分で調べる知識
* [HTML入門](http://www.ink.or.jp/~bigblock/html/) 基本的なHTMLの描き方を自習してください
* [template](http://westplain.sakuraweb.com/translate/GAE/Python/PythonTutorial/Using%2520Templates.cgi)　