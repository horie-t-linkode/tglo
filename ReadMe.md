

# tglo toggl出力ツール 

本プログラムは、Webサービス「toggl」に入力したエントリを取得し、画面に出力するツールである。

## 準備

### ビルド方法

(現時点では)ビルドはlinux環境で行う。

開発用ビルド
```
$ cd tglo
$ make
$ ls tglo
tglo
```

リリース用ビルド
```
$ cd tglo
$ make prod
$ ls bin/
linux  osx  windows
$ ls bin/linux/
tglo
```

### togglのAPIトークン、WORKSPACEIDの指定

本プログラム実行時に対象とするtogglアカウントのAPIトークン、WORKSPACEIDを指定する必要がある。  
指定する方法は以下の2通りある。

#### .envファイル

以下のような記述をした.envファイルを本プログラム実行ディレクトリに配置する。

```
TGLO_TOGGL_APITOKEN=<あなたのtogglのAPIトークン文字列>
TGLO_TOGGL_WORKSPACEID=<あなたのtogglのWORKSPACEID番号>
TGLO_DOCBASE_DOMAIN=<あなたのdocbaseのドメイン>
TGLO_DOCBASE_ACCESSTOKEN=<あなたのdocbaseのアクセストークン>
TGLO_DOCBASE_POSTING_TITLE=<docbaseメモ投稿時のタイトル>
TGLO_DOCBASE_POSTING_TAGS=<docbaseメモ投稿時のタグ。「,」区切りで指定。>
TGLO_DOCBASE_POSTING_GROUPS=<docbaseメモ投稿先のグループID。「,」区切りで指定。>
```

#### 環境変数

環境変数として以下の名前を設定する。

- TGLO_TOGGL_APITOKEN
- TGLO_TOGGL_WORKSPACEID
- TGLO_DOCBASE_DOMAIN
- TGLO_DOCBASE_ACCESSTOKEN
- TGLO_DOCBASE_POSTING_TITLE
- TGLO_DOCBASE_POSTING_TAGS
- TGLO_DOCBASE_POSTING_GROUPS

各変数の値の内容については「.envファイル」を参照。

## 実行

実行例は以下の通り。

ヘルプを表示。
```
> ./tglo help
togglエントリ/サマリを出力する

Usage:
  tglo [flags]
  tglo [command]

Available Commands:
  day         指定日のtogglエントリを出力する
  help        Help about any command
  lastweek    先週分のtogglエントリのサマリを出力する
  thisweek    今週分のtogglエントリのサマリを出力する
  today       本日のtogglエントリを出力する
  version     バージョンを表示
  week        指定日を含む週分のtogglエントリのサマリを出力する
  yesterday   昨日のtogglエントリを出力する

Flags:
  -h, --help      help for tglo
  -v, --verbose   開発者用デバッグ出力

Use "tglo [command] --help" for more information about a command.
```

本日分、昨日分、指定日分のtogglエントリを出力する。
```
> ./tglo today
> ./tglo yesterday
> ./tglo day --date 2020-04-10
```

今週分、先週分、指定日を含む週分のエントリを出力する。
```
> ./tglo thisweek
> ./tglo thisweek -s
> ./tglo lastweek
> ./tglo lastweek -s
> ./tglo week --date 2020-04-10
> ./tglo week --date 2020-04-10 -s
```

週分のエントリをdocbaseにメモとして作成する。
```
> ./tglo lastweek --postDocbase
```