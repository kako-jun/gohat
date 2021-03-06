[English](https://github.com/kako-jun/gohat)

# :no_good: gohat

[![Build Status](https://travis-ci.org/kako-jun/gohat.svg?branch=master)](https://travis-ci.org/kako-jun/gohat)

`gohat` は、危険でシンプルなコマンドラインツールです

gohat（御法度）は、「絶対やっちゃダメ」という意味の日本語です

引数で与えたスクリプトファイル（`.sh`、`.rb`、`.py` など）を、パスワードを入力することなく、root権限で実行します

Goで書かれているため、多くのOSで動作します

　

## Description

### Demo

![demo](https://raw.githubusercontent.com/kako-jun/gohat/master/assets/screen_1.gif)

### VS.

root権限が必要なコマンドは多くあります

それらは、実行時にパスワードを求められるため安全です

しかし、長時間かかる処理をスクリプトで自動化したい場合、root権限が必要なコマンドが含まれていると、そこで実行が止まってしまい不便です

そのスクリプトを `sudo` 付きで実行しても、最初にパスワードを求められるため解決しません

　

スクリプトファイルにSUIDを付けることで、解決するでしょうか？

実行ファイルならば、それで解決です

しかし、`.sh`、`.rb`、`.py` などに、SUIDを付けることはできません

　

`gohat` は、任意のスクリプトファイルを実行できるラッパーです

`gohat` にSUIDを付け、`gohat` 経由で呼び出すことにより、パスワード入力をパスできます

ただし、`gohat` は危険なため、乗っ取られる予定のある人は使っちゃダメです

誤ってファイルシステムをすべて消してしまっても、保証しません

自己責任でお使いください

　

## Installation

### Requirements

- Operating System

    - macOS
    - Linux

### Download binaries

- macOS: [gohat_mac.dmg](https://github.com/kako-jun/gohat/releases/latest)
- Linux ( `chmod u+x gohat` required)

    - x64: [gohat_linux_amd64.tar.gz](https://github.com/kako-jun/gohat/releases/latest)
    - ARM: [gohat_linux_arm64.tar.gz](https://github.com/kako-jun/gohat/releases/latest)
    - Raspberry Pi: [gohat_linux_armv7l.tar.gz](https://github.com/kako-jun/gohat/releases/latest)

### go get

```sh
$ go get github.com/kako-jun/gohat
```

　

## Features

### Usage

初回に1度だけ、`sudo` 付きで実行しておくことが必要です

```sh
$ sudo gohat
```

2回目からは、`sudo` 無しで実行します

```sh
$ gohat foo.sh

$ gohat bar.rb

$ gohat baz.py

$ gohat lightyear.pl
```

　

初回に何をしているかというと、`gohat` が `gohat` 自身にSUIDを付けています

（この時、引数は必要ありません）

処理の中身は、以下を実行しているのと同じです

```sh
chown root:root gohat
chmod u+s gohat
```

　

「なぜ `gohat` が便利なのか……？」の例を、以下に挙げます

#### Examples

##### e.g. スクリプトファイルに実行可能フラグが必要なく、オーナーも誰でも良い

```sh
$ chmod u+x foo.sh
```

しておくことは、必要ありません

オーナーがrootである必要もありません

##### e.g. remount に便利

Chromebookの[Crouton](https://github.com/dnschneid/crouton)では、SDカードは `noexec`、`nosuid` 付きでマウントされます

そのままでは、SDカードにリポジトリを置くのに不便なので、

```sh
mount -o remount,exec,suid /media/removable/SD\ Card
```

で解除したいのですが、`mount -o` はroot権限が必要なコマンドです

また、スリープすると `noexec`、`nosuid` が復活するので、毎回解除しなくてはいけません

何度もパスワードを入力するのは面倒です

　

このような場合は

```sh
mount -o remount,exec,suid /media/removable/SD\ Card
```

という内容で `remount.sh` を作り、`.bashrc` などの末尾で

```sh
gohat remount.sh
```

することで、パスワード入力を省略できます

　

#### Unsupported

##### 共有PCで使うのは超危険

あらゆるコマンドをroot権限で実行できるため、システムを破壊できてしまいます

PATHに追加するのは超危険です

SUIDが付いた状態でディスク上にただ存在するだけでも、存在を知られると悪用される可能性があります

少なくとも `gohat` からリネームしておきましょう

　

### Coding

```golang
import "github.com/kako-jun/gohat/gohat-core"

gohat.Exec(scriptPath)
```

### Contributing

Pull Requestを歓迎します

- `gohat` をより便利にする機能の追加
- より洗練されたGoでの書き方
- バグの発見、修正
- もっと良い英訳、日本語訳があると教えたい

など、アイデアを教えてください

　

## Authors

kako-jun

- :octocat: https://github.com/kako-jun
- :notebook: https://gist.github.com/kako-jun
- :house: https://llll-ll.com
- :bird: https://twitter.com/kako_jun_42

### :lemon: Lemonade stand

寄付を頂けたら、少し豪華な猫エサを買おうと思います

下のリンクから、Amazonギフト券（Eメールタイプ）を送ってください

「受取人」欄には `kako.hydrajin@gmail.com` と入力してください

　**[:hearts: Donate](https://www.amazon.co.jp/gp/product/B004N3APGO/ref=as_li_tl?ie=UTF8&tag=llll-ll-22&camp=247&creative=1211&linkCode=as2&creativeASIN=B004N3APGO&linkId=4aab440d9dbd9b06bbe014aaafb88d6f)**

- 「メッセージ」欄を使って、感想を伝えることもできます
- 送り主が誰かは分かりません
- ¥15 から送れます

　

## License

This project is licensed under the MIT License.

See the [LICENSE](https://github.com/kako-jun/gohat/blob/master/LICENSE) file for details.

## Acknowledgments

- [Go](https://golang.org/)
- and you
