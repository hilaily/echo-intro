---
title : 自定义
menu:
  side:
    name: 自定义
    parent: guide
    weight: 3
---

## 自定义

### HTTP 错误处理

`Echo#SetHTTPErrorHandler(h HTTPErrorHandler)` 注册了一个自定义的 `Echo#HTTPErrorHandler`.

默认的 HTTP 错误处理规则如下：

- 如果错误是`Echo#HTTPError`类型则发送一个状态码为 `HTTPError.Code`，内容为 `HTTPError.Message` 的HTTP 响应。
- 否则返回 `500 - Internal Server Error`。
- 如果开启了 debug 模式，则使用 `error.Error()` 返回消息。

### Debug

`Echo#SetDebug(on bool)` 开启/关闭 debug 模式。

### 日志

#### 自定义日志

`Echo#SetLogger(l log.Logger)`

SetLogger 用来定义一个 自定义日志。

#### 日志输出

`Echo#SetLogOutput(w io.Writer)` 设置日志输出的位置，默认是 `os.Stdout`。
使用 `Echo#SetLogOutput(io.Discard)` 完全禁用日志。

#### 日志级别

`Echo#SetLogLevel(l log.Level)`

SetLogLevel 用于设置日志级别，默认是 `3` (ERROR).
可以使用的值：

- `0` (DEBUG)
- `1` (INFO)
- `2` (WARN)
- `3`	(ERROR)
- `4`	(FATAL)
- `5` (OFF)

### HTTP Engine

Echo 现在支持使用 standard 和 [fasthttp](https://github.com/valyala/fasthttp) 提供 HTTP 服务。
Echo 内部实现了这两个引擎的接口，可以根据需要无缝的切换这两种 HTTP 服务。

#### 运行一个 standard HTTP server

`e.Run(standard.New(":1323"))`

#### 运行一个 fasthttp server

`e.Run(fasthttp.New(":1323"))`

#### 运行一个带有 TLS 配置的 HTTP 服务

`e.Run(<engine>.WithTLS(":1323", "<certFile>", "<keyFile>"))`

#### 运行一个带有服务配置文件的 HTTP 服务

`e.Run(<engine>.WithConfig(<config>))`

##### 配置

```go
Config struct {
  Address      string        // TCP address to listen on.
  Listener     net.Listener  // Custom `net.Listener`. If set, server accepts connections on it.
  TLSCertFile  string        // TLS certificate file path.
  TLSKeyFile   string        // TLS key file path.
  ReadTimeout  time.Duration // Maximum duration before timing out read of the request.
  WriteTimeout time.Duration // Maximum duration before timing out write of the response.
}
```

#### 服务实例自己配置属性

```go
s := standard.New(":1323")
s.MaxHeaderBytes = 1 << 20
e.Run(s)
```
