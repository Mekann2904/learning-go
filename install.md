# golangを導入する

## Homebrewを使用してインストールする。

```
brew install go
```

## macの環境変数

```
export GOROOT="$(brew --prefix go)/libexec"
export GOPATH="$HOME/go"
export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
```

## インストール確認

```
go version
```
