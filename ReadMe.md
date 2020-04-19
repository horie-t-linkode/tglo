

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
```

#### 環境変数

環境変数で以下の値を設定する。

- TGLO_TOGGL_APITOKEN : <あなたのtogglのAPIトークン文字列>
- TGLO_TOGGL_WORKSPACEID : <あなたのtogglのWORKSPACEID番号>

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