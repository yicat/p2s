#! /bin/bash

# 翻墙
git config http.proxy http://localhost:1080
git config https.proxy http://localhost:1080
export http_proxy=http://localhost:1080
export https_proxy=http://localhost:1080

# 安装 vscode 依赖
go get -u -v github.com/mdempsky/gocode
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/derekparker/delve/cmd/dlv
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/sqs/goreturns
go get -u -v  github.com/mdempsky/gocode