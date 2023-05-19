# WebClip

![ソースコード言語](https://img.shields.io/github/languages/top/geniusmaaakun/webclip)

![ソースコード数](https://img.shields.io/github/languages/count/geniusmaaakun/webclip)

![ソースコードサイズ](https://img.shields.io/github/languages/code-size/geniusmaaakun/webclip)


![logo](./frontend/public/favicon.ico)


## Overview
Get HTML files from a website on the command line, convert them to markdown format, and save them in a specified folder.
By specifying the option, image files are also downloaded.

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
* -u: specify the URL
* -o: specify the save directory
* -save: save to database
* -d: also save the image file of the target page

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

* search: Search the database for files that match the conditions from the saved files. Only if --save is specified
* search: 保存したファイルから条件に合うファイルをデータベースから探します。--save を指定した場合に限ります
```
webclip search -b "{body}"
webclip search -t "{title}"
```

* server: Display the list page of saved files. Simple as it is a command line tool
* server: 保存したファイルの一覧ページを表示します。あくまでコマンドラインツールなので簡易的です
```
webclip server
```

* clean: Delete data from DB if file path does not exist. You can save DB space.
* clean: ファイルパスが存在しない場合、DBからデータを削除します。DBの容量を節約できます。
```
webclip clean
```

* zip: zip the file
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

- README書き方
* https://qiita.com/legitwhiz/items/bb34ef20ba23336e0c87
* https://qiita.com/Kotabrog/items/fb328b72ac94137897af#requirement

- バッジ
* https://kic-yuuki.hatenablog.com/entry/2019/06/29/173256
* https://shields.io/


## Author

[twitter@geniusmaaakun](https://twitter.com/geniusmaaakun)

## Licence

[MIT](https://opensource.org/license/mit/)

