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
$ cd <PROJECT IN $GOPATH>
$ go get -u github.com/labstack/echo/...
```

### 使用 dep
```shell
$ cd <PROJECT IN $GOPATH>
$ dep ensure -add github.com/labstack/echo@^3.1
```

### 使用 glide

```shell
$ cd <PROJECT IN $GOPATH>
$ glide get github.com/labstack/echo#~3.1
```

### 使用 govender

```shell
$ cd <PROJECT IN $GOPATH>
$ govendor fetch github.com/labstack/echo@v3.1
```

Echo 使用 Go `1.10.x`开发，并通过了 `1.9.x` 和 `1.10.x` 的测试。

Echo 通过 GitHub releases 进行 [语义化版本(semantic versioning)](http://semver.org) 控制，特定的版本可以使用 [package manager](https://github.com/avelino/awesome-go#package-management) 安装。

