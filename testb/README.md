# SrcUrl: https://zenn.dev/robes/articles/c31dc875ae29c3

【Python】Youtube動画から音声ファイルを取得する

[Zenn](/)

[![osn_Lofi](https://storage.googleapis.com/zenn-user-upload/avatar/607ef66b0a.jpeg)osn\_Lofi](/robes)

👻

# 【Python】Youtube動画から音声ファイルを取得する

2022/10/16に公開・約400字

[ツイート](https://twitter.com/intent/tweet?url=https://zenn.dev/robes/articles/c31dc875ae29c3&text=%E3%80%90Python%E3%80%91Youtube%E5%8B%95%E7%94%BB%E3%81%8B%E3%82%89%E9%9F%B3%E5%A3%B0%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B%EF%BD%9Cosn_Lofi&hashtags=zenn)

[![](https://storage.googleapis.com/zenn-user-upload/topics/ebddf7c6bb.png)\
\
Python](/topics/python) [![](https://storage.googleapis.com/zenn-user-upload/topics/17167d84a4.png)\
\
YouTube](/topics/youtube) [#\
音声](/topics/%E9%9F%B3%E5%A3%B0) [#\
mp3](/topics/mp3) [#\
ytdlp](/topics/ytdlp) [![](https://zenn.dev/images/drawing/tech-icon.svg)\
\
tech](/tech-or-idea)

**Youtube動画から音声ファイルのみをダウンロードする方法です。**

**GoogleColaboratoryで簡単に実装できます。**

## 1\. ライブラリーをインストール

```python
!pip install yt-dlp

```

## 2\. コマンドラインで実行

`https://youtu.be/vR5GXyklvm4` のところに、 **ダウンロードしたいYouTubeのアドレスを入力** します。

GoogleColaboratoryですと、contentディレクトリにmp3ファイルが保存されます。

以下は首相官邸のYouTubeになります。

```python
!yt-dlp -x --audio-format mp3 https://youtu.be/vR5GXyklvm4

```

### 公式はこちら

[https://github.com/yt-dlp/yt-dlp](https://github.com/yt-dlp/yt-dlp)

[ツイート](https://twitter.com/intent/tweet?url=https://zenn.dev/robes/articles/c31dc875ae29c3&text=%E3%80%90Python%E3%80%91Youtube%E5%8B%95%E7%94%BB%E3%81%8B%E3%82%89%E9%9F%B3%E5%A3%B0%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B%EF%BD%9Cosn_Lofi&hashtags=zenn)

[![osn_Lofi](https://storage.googleapis.com/zenn-user-upload/avatar/607ef66b0a.jpeg)](/robes)

[osn\_Lofi](/robes)

Python始めて2年目です。機械学習やディープラーニング・自然言語処理を勉強しています。最近はKaggleと統計にハマっています。専門は金融です。

### Discussion

![](https://zenn.dev/images/drawing/discussion.png)

ログインするとコメントできます

Login

著者[![@robes](https://storage.googleapis.com/zenn-user-upload/avatar/607ef66b0a.jpeg)\
osn\_Lofi](/robes)

公開2022/10/16

文章量約400字

[![osn_Lofi](https://storage.googleapis.com/zenn-user-upload/avatar/607ef66b0a.jpeg)](/robes)

[osn\_Lofi](/robes)

Python始めて2年目です。機械学習やディープラーニング・自然言語処理を勉強しています。最近はKaggleと統計にハマっています。専門は金融です。

目次

1. [1\. ライブラリーをインストール](#1.-%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E3%83%BC%E3%82%92%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB)
2. [2\. コマンドラインで実行](#2.-%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%83%A9%E3%82%A4%E3%83%B3%E3%81%A7%E5%AE%9F%E8%A1%8C)
1. [公式はこちら](#%E5%85%AC%E5%BC%8F%E3%81%AF%E3%81%93%E3%81%A1%E3%82%89)

[Zenn](/)

エンジニアのための

情報共有コミュニティ

#### About

- [Zennについて](/about)
- [運営会社](https://classmethod.jp)
- [お知らせ・リリース](https://info.zenn.dev)

#### Guides

- [使い方](/zenn)
- [Publication](/publications)
- [よくある質問](/faq)

#### Links

- [Twitter](https://twitter.com/zenn_dev)
- [GitHub](https://github.com/zenn-dev)
- [メディアキット](/mediakit)

#### Legal

- [利用規約](/terms)
- [プライバシーポリシー](/privacy)
- [特商法表記](/terms/transaction-law)

[![Classmethod inc.](https://zenn.dev/images/classmethod-logo-small.svg)](https://classmethod.jp/)