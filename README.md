# webclipをP2Pで共有


再起ダウンロード設計
スクレイピング
imgタグを見つけたらダウンロード、置き換え


マークダウン形式
	画像を表示
　　 h1 対応表
    https://www.sirochro.com/note/html-markdown-list/
    https://bayashita.com/p/entry/show/109
	markdown_to_html
    https://github.com/JohannesKaufmann/html-to-markdown
	md2html
    https://github.com/nocd5/md2html
	markdown
    https://github.com/gomarkdown/markdown
	godown
    https://github.com/mattn/godown



    htmlを保存する
    htmlをmdに変換
    html削除


    日付、タイトルで保存
    画像はimageに保存
    コマンドライン引数を受け取る -u -o


# cliを使う　済み
# 変数を構造体にまとめる　済み
# デフォルトの保存パスを設定しておく　
# 画像の保存のフラグを作る　済み
# 文字コード変換　済み
    https://qiita.com/koki_develop/items/dab4bcbb1df1271a17b6

# すでに存在するフォルダの場合は追加  →親ディレクトリを作っておく
    os.Ixistでエラーの中身を判定

    もしくは、複数ダウンロード。フォルダ分けする yaml形式で渡すこともできる





# メモアプリを作成。
        スレッド形式　チャットのような感じ
        マークダウン記述

# chatgptに検索させて、取り込む
    
# functional optionパターンでオプション設定
    ```
    //｀With~返り値の関数にsrvを渡す感じになる
	for _, opt := range options {
		opt(srv)
	}
    ```

# テスト

# P2Pで共有



# パスをDBに保存　パスが無ければ削除
　id name path srcurl 保存　
    一覧取得
　　clean  
    search タイトルを検索

# ブラウザで表示　再度html 
    データは最初にすべて取得ブラウザ表示。
    順次取得の方がパフォーマンス上がる？

# アプリでも見れる。  flutter




# 構造
App
downloader
    imageDownloader
    htmlDownloader

converter
    toMarkdown()
    mdSave()

repo


# APIを作成 gollira/mux

# React
やりたいこと
テスト、Ts

## プロジェクトの作成
npx create-react-app frontend --template typescript

テキストボックス
入力をuseStateで管理
入力ボックスに入力しエンターを押したアラaxiosでデータを取得開始
useStateで結果を配列に入れて、表示



サーチバー
親コンポーネントからステートを受け取る。更新関数

結果のボックス
ステートの中身をリスト表示


テスト