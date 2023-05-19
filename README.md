# WebClip

bachi

![test](./frontend/build/static/favicon.ico)


## Overview
Get HTML files from a website on the command line, convert them to markdown format, and save them in a specified folder.
By specifying the option, image files are also downloaded.

コマンドラインでWebサイトからHTMLファイルを取得し、マークダウン形式に変換して、指定したフォルダに保存します。
オプションを指定することで、画像ファイルもダウンロードします。


## Requirement
環境や、必要なライブラリなどについて言及しています。
……とはいっても、場合によっては動作確認をした環境だけ載せているという感じで、必ずしも必要としてない条件も載せてしまっています。
例えば以下のように示しています。
必要があればバージョンなども載せています。

- macOS
- clang
- Docker 19.03.6


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



## Features
詳しい仕様について、基本的に箇条書きで紹介しています。
Usageで紹介しなかった詳しい使い方も書いています。
箇条書きで書きづらい場合や、長くなりそうな場合には、「Features」ではなく「Description」に変更したほうが個人的にはピンときます。

## Reference
参考URLを書きます。



## Author

[twitter@geniusmaaakun](https://twitter.com/geniusmaaakun)

## Licence

[MIT](https://......)

