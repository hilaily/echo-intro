+++
title = "WebSocket Recipe"
description = "WebSocket recipe / example for Echo"
url="/recipes/websocket"
[menu.side]
  name = "WebSocket"
  parent = "recipes"
  weight = 5
+++

## WebSocket Recipe

> 只支持 `standard` 引擎

### 使用 `net` 库的 WebSocket

#### 服务端

`server.go`

{{< embed "websocket/net/server.go" >}}

### 使用 `gorilla` 的 WebSocket

#### 服务端

`server.go`

{{< embed "websocket/gorilla/server.go" >}}

### 客户端

`index.html`

{{< embed "websocket/public/index.html" >}}

### 输出示例

`Client`

```sh
Hello, Client!
Hello, Client!
Hello, Client!
Hello, Client!
Hello, Client!
```

`Server`

```sh
Hello, Server!
Hello, Server!
Hello, Server!
Hello, Server!
Hello, Server!
```

### 维护者

- [vishr](https://github.com/vishr)

### [Source Code]({{< source "websocket" >}}


