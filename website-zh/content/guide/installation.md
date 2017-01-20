+++
title = "安装"
slug = "installation"
[menu.side]
  parent = "guide"
  weight = 1
+++

## 安装

### 要求

- 安装 Go
- 设置 GOPATH 环境变量

### 使用 go get

```shell
$ cd <project in $GOPATH>
$ go get github.com/labstack/echo/...
```

###  使用 glide

```shell
$ cd <project in $GOPATH>
$ glide get github.com/labstack/echo#~3.0
```

### 使用 govender

```shell
$ cd <project in $GOPATH>
$ govendor fetch github.com/labstack/echo@v3.0
```

Echo 在 Go 1.7.x 下开发，1.6.x 和 1.7.x 下测试通过。

基于 [semantic versioning](http://semver.org) 的管理版本，Echo 使用 GitHub releases 。特别的版本可以使用 [package manager](https://github.com/avelino/awesome-go#package-management) 安装。