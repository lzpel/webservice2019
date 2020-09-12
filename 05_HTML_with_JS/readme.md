# 解説
## これはなに？
HTMLでjavascriptを読み込むサンプル
動的に日付が表示される
## 何を学ぶのか？
japascriptとはブラウザ上で動作するスクリプト言語でページに動きをつけるために使われる。
## 教える知識と前回との差分
- 追記 index.htmlでスクリプトを読み込んだ<script src="/script.js"></script>
```
ここで/static/style.cssに書かれた内容を読み込んでいる
*.cssはHTMLの表示の仕方を変える
```
- 変更 ID属性を指定した<p id="name">この文章は日付に置き換えられる</p>
```
このid="name"という部分でnameというidを宣言している
後述するscript.jsでdocument.getElementById("name")という関数呼び出しで指定される
```
- ファイル script.jsというjavascriptがstaticディレクトリ内に作られている
```
この関数は現在の日時を2019/05/21という形式の文字列で返す
function getNowDate(){
  var dt = new Date();
  var y = dt.getFullYear();
  var m = ("00" + (dt.getMonth()+1)).slice(-2);
  var d = ("00" + dt.getDate()).slice(-2);
  var result = y + "/" + m + "/" + d;
  return result;
}
この関数はgetNowDate()で取得した日付の文字列をnameというidの要素のテキストに設定する
function nowTime(){
  var node = document.getElementById("name");
  node.textContent = getNowDate();
}
nowTime();
```
巷にあふれる動くサイトはjavascriptが使われている。この講座でもJavascriptのライブラリであるＪｑｕｅｒｙなどを今後使用する。
JavaScriptの書き方などはここで解説する余裕がない。各自調べてください。リンクは張っておきます。
## 自分で調べる知識
* [HTML入門](https://reference.hyper-text.org/html5/attribute/id/)
* [javascript](https://www.javadrive.jp/javascript/) 
* [javascriptとDOM](https://www.javadrive.jp/javascript/dom/) 