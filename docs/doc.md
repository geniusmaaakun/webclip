# 今後の取り組み

* reactから保存機能
    保存ボタン設置、APIに投げる

* CICD
    github actions

* UI
    かっこいいUIに整える

* FlutterApp
    markdown
    chatgptに投げる
    保存したデータを表示する

```
    import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:webclip/models/markdown.dart';
import 'package:webclip/providers/markdown_provider.dart';
import 'package:webclip/screens/markdown_screen.dart';

class SearchScreen extends HookConsumerWidget {
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final markdowns = ref.watch(markdownsProvider);

    final inputController = useTextEditingController();
    final resultData = useState<List<Markdown>>([]);

    final navigate = useNavigator();

    useEffect(() {
      ref.listen<AsyncValue<List<Markdown>>>(markdownsProvider, (state) {
        state.when(
          data: (data) {
            resultData.value = data;
          },
          loading: () {},
          error: (error, stackTrace) {},
        );
      });
      ref.read(markdownsProvider.notifier).loadMarkdowns();
    }, []);

    void search(String value) {
      if (value.isEmpty) {
        resultData.value = markdowns.data?.value ?? [];
        return;
      }

      final searchKeywords = value.trim().toLowerCase().split(RegExp(r'\s+'));

      if (searchKeywords.isEmpty) {
        resultData.value = markdowns.data?.value ?? [];
        return;
      }

      final serchedPosts = markdowns.data?.value
          ?.where((data) =>
              searchKeywords.every((kw) =>
                  data.title
                      .toString()
                      .toUpperCase()
                      .contains(kw.toUpperCase())))
          .toList() ??
          [];

      resultData.value = serchedPosts;
    }

    void onClickMarkdown(String id) {
      navigate.push(MaterialPageRoute(builder: (_) => MarkdownScreen(id: id)));
    }

    void onClickHome() {
      navigate.pop();
    }

    return Scaffold(
      body: Column(
        children: [
          GestureDetector(
            onTap: onClickHome,
            child: Text('WebClip', style: Theme.of(context).textTheme.headline4),
          ),
          TextField(
            controller: inputController,
            onChanged: (value) => search(value),
            decoration: InputDecoration(hintText: '検索キーワードを入力'),
          ),
          Expanded(
            child: ListView.builder(
              itemCount: resultData.value.length,
              itemBuilder: (BuildContext context, int index) {
                final item = resultData.value[index];
                return GestureDetector(
                  onTap: () => onClickMarkdown(item.id),
                  child: ListTile(
                    title: Text(item.title ?? ''),
                    subtitle: Text(item.path ?? ''),
                    trailing: Text(item.createdAt?.toString() ?? ''),
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}
```

* テスト
    E2Eテスト
    ユニットテスト

* カスタムエラー
    エラーをカスタムする


# Flutter 設計
golangのサーバーはそのまま使う。（内部サーバー） + github api
もしくはwebclipのサーバーを外部サーバーとして運用 一人5GBまで無料 aws
外部サーバーとして運用するだけなら、UIだけでOK
サーチバーとエディターのみ 実装する

別のリポジトリとして構築

　動画保存　url?
　flutter github連携
	リポジトリを作成　設定
	リポジトリに追加、取得、表示

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
E2Eテスト　月〜金
	Cypress. 　　調べる
CI/CD. 土日. GitHub actions

# 通知機能　

https://qiita.com/ryotanny/items/670b3aba7ea8e57eb776

webpush 
https://blog.capilano-fw.com/?p=11536

# 状態管理　Zustand. 



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

# UI
	左にリストと検索バー、右にマークダウン
	保存機能　reactで保存ボタンを作る。保存ボタンを押したら、APIに送信
    2 詳細ページ　	マークダウンエディタ実装
　　		ハイライトは無視、エンターで保存。表示だけでも良いかも
    そこそこかっこいいUI探す
    ChacraUI

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



# clean architecture 理想
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