# WebClip
コマンドラインでWebサイトからHTMLファイルを取得し、マークダウン形式に変換して、指定したフォルダに保存します。

# コマンドオプション一覧
-u: URLを指定します
-o: 保存ディレクトリを指定します
-save: データベースに保存します
-d: 対象ページの画像ファイルも保存します

# サブコマンド一覧

search: 保存したファイルから条件に合うファイルを探します
```
webclip search -b "{body}"
webclip search -t "{title}"
```

server: 保存したファイルの一覧ページを表示します。あくまでコマンドラインツールなので簡易的です
```
webclip server
```

clean: ファイルパスが存在しない場合、DBからデータを削除します。DBの容量を節約できます。


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



```
project_root/
│
├── src/
│   ├── domain/
│   │   ├── entities/
│   │   │   └── (エンティティとビジネスルールを定義するファイル)
│   │   └── value_objects/
│   │       └── (バリューオブジェクトを定義するファイル)
│   │
│   ├── use_cases/
│   │   ├── interfaces/
│   │   │   └── (ユースケースのインターフェースを定義するファイル)
│   │   └── implementations/
│   │       └── (ユースケースの具体的な実装を定義するファイル)
│   │
│   ├── interfaces/
│   │   ├── controllers/
│   │   │   └── (コントローラーを定義するファイル)
│   │   ├── presenters/
│   │   │   └── (プレゼンターを定義するファイル)
│   │   └── views/
│   │       └── (ビューを定義するファイル)
│   │
│   └── infrastructure/
│       ├── repositories/
│       │   └── (リポジトリの実装を定義するファイル)
│       ├── services/
│       │   └── (外部サービスとの連携を担当するファイル)
│       └── (データベース設定、外部API設定などの設定ファイル)
│
├── tests/
│   ├── domain/
│   ├── use_cases/
│   ├── interfaces/
│   └── infrastructure/
│
└── (プロジェクトに関連するその他のファイル, 例: README.md, .gitignore, etc.)


```

```
ディレクトリ構成と役割：

domain - ドメイン層に関連するエンティティ、バリューオブジェクト、およびビジネスルールを定義します。 

use_cases - ユースケース層に関連するインターフェースと実装を定義します。これにはアプリケーションの主要な機能や操作が含まれます。 

interfaces - インターフェース層に関連するコントローラー、プレゼンター、およびビューを定義します。これらはユーザーや外部システムとのやり取りを担当し、入力を受け取り、出力を表示または送信します。

infrastructure - インフラ層に関連するリポジトリの実装や外部サービスとの連携、データベース設定、外部API設定などを定義します。この層はアプリケーションの技術的な詳細を担当し、他の層とは独立して変更・進化させることができます。

tests - 各層に対応するテストコードを配置します。適切なテストカバレッジを維持することで、アプリケーションの安定性や保守性が向上します。

このディレクトリ構成は、クリーンアーキテクチャの原則に基づいていますが、実際のプロジェクトでは言語やフレームワークに応じて適宜調整してください。また、プロジェクトが大規模になると、各層をさらにサブディレクトリに分割することで、コードの整理や管理が容易になります。
```