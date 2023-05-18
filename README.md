# WebClip
コマンドラインでWebサイトからHTMLファイルを取得し、マークダウン形式に変換して、指定したフォルダに保存します。
オプションを指定することで、画像ファイルもダウンロードします

# 使い方
## コマンドオプション一覧
-u: URLを指定します
-o: 保存ディレクトリを指定します
-save: データベースに保存します
-d: 対象ページの画像ファイルも保存します

## サブコマンド一覧

search: 保存したファイルから条件に合うファイルをデータベースから探します。--save を指定した場合に限ります
```
webclip search -b "{body}"
webclip search -t "{title}"
```

server: 保存したファイルの一覧ページを表示します。あくまでコマンドラインツールなので簡易的です
```
webclip server
```

clean: ファイルパスが存在しない場合、DBからデータを削除します。DBの容量を節約できます。
```
webclip clean
```

zip: ファイルをzip化します
```
webclip zip
webclip zip -t "{title}"
webclip zip -b "{body}"
webclip zip -t "{title}" -b "{body}"
```

