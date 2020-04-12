

# tglo toggl出力ツール 

本プログラムは、Webサービス「toggl」に入力したエントリを取得し、画面に出力するツールである。

## 準備

### togglのAPIトークン、WORKSPACEIDの指定

本プログラム実行時に対象とするtogglアカウントのAPIトークン、WORKSPACEIDを指定する必要がある。

#### 環境変数

環境変数で指定する場合の例は以下の通り。

```
> TGLO_TOGGL_APITOKEN=<あなたのtogglのAPIトークン文字列> TGLO_TOGGL_WORKSPACEID=<あなたのtogglのWORKSPACEID番号> ./tglo yesterday
```

#### .envファイル

以下のような記述をした.envファイルを本プログラム実行ディレクトリに配置する。

```
TGLO_TOGGL_APITOKEN=<あなたのtogglのAPIトークン文字列>
TGLO_TOGGL_WORKSPACEID=<あなたのtogglのWORKSPACEID番号>
```

## 実行

実行例は以下の通り。
```
> ./tglo help
> ./tglo today
> ./tglo yesterday
> ./tglo day --date 2020-04-10
> ./tglo thisweek
> ./tglo lastweek
> ./tglo thisweek -s
> ./tglo lastweek -s
```