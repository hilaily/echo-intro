---
title: HTTP2
url : http2
menu:
  side:
    identifier: http2
    parent: recipes
    weight: 3
---

## 什么是 HTTP2?

HTTP/2 (原本的名字是 HTTP/2.0) 是万维网使用的 HTTP 网络协议的第二个主要版本。

### 特性

- 使用二进制格式传输数据，而不是文本。使得在解析和优化扩展上更为方便。
- 多路复用，所有的请求都是通过一个 TCP 连接并发完成。
- 对消息头采用 HPACK 进行压缩传输，能够节省消息头占用的网络的流量。
- Server Push：服务端能够更快的把资源推送给客户端。

## 怎样运行 HTTP2 和 HTTPS 服务?

> 只用 `standard` 引擎支持.

### 生成一个自签名的 X.509 TLS 证书(HTTP/2 需要 TLS 才能运行)

```sh
go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
```

上面的明亮会生一个`cert.pem` 和`key.pem` 文件。

> 这里只是展示使用，所以我们用了自签名的证书，正式环境建议去
[CA](https://zh.wikipedia.org/wiki/%E6%95%B0%E5%AD%97%E8%AF%81%E4%B9%A6%E8%AE%A4%E8%AF%81%E6%9C%BA%E6%9E%84)申请证书。

### 配置引擎 `engine.Config`

`server.go`

{{< embed "http2/server.go" >}}

### 最后

- https://localhost:1323/request (显示 HTTP 请求信息)
- https://localhost:1323/stream (实时展示当前时间)

### 维护者

- [vishr](https://github.com/vishr)

### [Source Code]({{< source "http2" >}})
