# 概要
シンプルな一行日記用のCLIツールです。

# インストール方法
```sh
$ go install github.com/midry3/diajog/cmd/dj@latest
```

# 使い方
```sh
$ dj <オプション> "内容"
```

# オプション一覧
### -h, --help
ヘルプメッセージを表示します。

### -v, --view
日記を見返すことが出来ます。出力が長い時は`less`コマンドを一緒に使ってください。
