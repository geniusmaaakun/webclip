<div style="text-align: center;">

# WebClip

![ソースコード言語](https://img.shields.io/github/languages/top/geniusmaaakun/webclip)

![ソースコード数](https://img.shields.io/github/languages/count/geniusmaaakun/webclip)

![ソースコードサイズ](https://img.shields.io/github/languages/code-size/geniusmaaakun/webclip)


![logo](./frontend/public/favicon.ico)

</div>



## Overview
コマンドラインでWebサイトからHTMLファイルを取得し、マークダウン形式に変換して、指定したフォルダに保存します。
オプションを指定することで、画像ファイルもダウンロードします。
特定のURLの資料などを、ローカルに保存しておきたいときに使います。

## Requirement
- MacOS
- golang: 1.19
- react 17.0.2

## install
```
make install
```

## Usage
### コマンドオプション一覧
* -u: URLを指定します
* -o: 保存ディレクトリを指定します
* -save: データベースに保存します
* -d: 対象ページの画像ファイルも保存します

```
// ファイルを変換して保存します
webclip -u "{URL}" -o "{OUTPUT_DIR}"

// ファイルを変換して保存し、データベースに保存します
webclip -u "{URL}" -o "{OUTPUT_DIR}" -save

// ファイルを変換して保存し、データベースに保存し、対象ページの画像ファイルも保存します
webclip -u "{URL}" -o "{OUTPUT_DIR}" -save -d
```

###  サブコマンド一覧
* search: 保存したファイルから条件に合うファイルをデータベースから探します。--save を指定した場合に限ります
```
webclip search -b "{body}"
webclip search -t "{title}"
```

* server: 保存したファイルの一覧ページを表示します。あくまでコマンドラインツールなので簡易的です
```
webclip server
```

* clean: ファイルパスが存在しない場合、DBからデータを削除します。DBの容量を節約できます。
```
webclip clean
```

* zip: ファイルをzip化します
```
webclip zip
webclip zip -t "{title}"
webclip zip -b "{body}"
webclip zip -t "{title}" -b "{body}"
```



## Features and Description

- 資料のコピーを手動でコピペするを自動化するために作成しました。
コマンドラインで実行することで、指定したURLのHTMLファイルを取得し、マークダウン形式に変換して、指定したフォルダに保存します。

- clean architecture \
クリーンアーキテクチャの練習用に作成しました。

- react + golang \
reactとgolangを使ってみたかったので、フロントエンドにreactを採用しました。


## Reference

- how to write README 
* https://qiita.com/legitwhiz/items/bb34ef20ba23336e0c87
* https://qiita.com/Kotabrog/items/fb328b72ac94137897af#requirement

- github shields
* https://kic-yuuki.hatenablog.com/entry/2019/06/29/173256
* https://shields.io/

* How to Embed React App into Go Binary
https://medium.com/@pavelfokin/how-to-embed-react-app-into-go-binary-12905d5963f0

https://www.smartinary.com/blog/how-to-embed-a-react-app-in-a-go-binary/


## Author

[twitter@geniusmaaakun](https://twitter.com/geniusmaaakun)

## Licence

[MIT](https://opensource.org/license/mit/)
