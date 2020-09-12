# 解説
## これはなに？
HTMLにスタイルを設定し色を変えたり角度を変えたりする
## 何を学ぶのか？
スタイル及びcssの取り扱い
## 教える知識と前回との差分
- 追記 HTMLの中に<link href="/style.css" rel="stylesheet">と書かれている
```
ここで/static/style.cssに書かれた内容を読み込んでいる
*.cssはHTMLの表示の仕方を変える
```
- 変更 HTML内の<h2>クラスにclass属性を付与した
```
<!-- クラスを指定した -->
<h2 class="redclass">サブタイトル</h2>
```
- ファイル staticディレクトリにstyle.cssが作られている
```
/static/style.cssの中身は以下の通り
.から始まる内容はクラスを表す{}内の情報がクラスに付与される
実際のHTMLにどのclassを割り当てるかはHTMLタグのclass属性で決まる
ここではcolor : red; で色を変えている
.redclass {
    color: red;
}
.や#以外で始まる内容はタグを表している。ここでは４５度傾けている
img {
    transform: rotate(-45deg);
}
```
巷にあふれる美しいデザインのサイトやウェブサービスもHTMLとCSSの組み合わせで構成されている
## 自分で調べる知識
* [HTML入門](http://www.ink.or.jp/~bigblock/html/) ここに<link>タグについての解説もある
* [CSS入門](https://saruwakakun.com/html-css/basic/css)
