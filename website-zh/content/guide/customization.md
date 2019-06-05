---
title : 自定义
slug: aaa
url: guide/customization
menu:
  side:
    name: 自定义
    parent: guide
    weight: 2
---

## 自定义

### Debug

可使用`Echo#Debug` 来开启/关闭 debug 模式。Debug 模式下的日志级别是 **DEBUG**。

### 日志

日志默认使用 JSON 格式，可从通过修改标头进行格式修改。

#### 日志标头
 `Echo#Logger.SetHeader(io.Writer)` 用于日志标头的设置，默认值为：

```json
{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}
```

*示例* 

```go
import "github.com/labstack/gommon/log"

/* ... */

if l, ok := e.Logger.(*log.Logger); ok {
  l.SetHeader("${time_rfc3339} ${level}")
}
```

```bash
2018-05-08T20:30:06-07:00 INFO info
```

**可用标签**

- `time_rfc3339`
- `time_rfc3339_nano`
- `level`
- `prefix`
- `long_file`
- `short_file`
- `line`


#### 日志输出

`Echo#Logger.SetOutput(io.Writer)` 用于设置日志输出的位置，默认是 `os.Stdout`。

若需禁用日志，可使用 `Echo#Logger.SetOutput(ioutil.Discard)` 或 `Echo#Logger.SetLevel(log.OFF)`来完成。

#### 日志级别

`Echo#Logger.SetLevel(log.Lvl)` 用于设置日志级别，默认是 `ERROR`。可选值包括：

- `DEBUG`
- `INFO`
- `WARN`
- `ERROR`
- `OFF`

#### 自定义日志

Echo 的日志实现了  `echo.Logger`  接口，该接口允许使用 `Echo#Logger`注册自定义日志。

### 自定义 Server

使用`Echo#StartServer()`进行自定义 Server 的启动

*示例*

```go
s := &http.Server{
  Addr:         ":1323",
  ReadTimeout:  20 * time.Minute,
  WriteTimeout: 20 * time.Minute,
}
e.Logger.Fatal(e.StartServer(s))
```

### 启动横幅
使用 `Echo#HideBanner` 隐藏启动横幅。

### 自定义监听器
使用 `Echo#*Listener`启动一个自定义的 listener。
*示例*

```go
l, err := net.Listen("tcp", ":1323")
if err != nil {
  e.Logger.Fatal(l)
}
e.Listener = l
e.Logger.Fatal(e.Start(""))
```

### 禁用 HTTP/2

使用`Echo#DisableHTTP2` 关闭 HTTP/2 协议。

### 读取超时

使用`Echo#*Server#ReadTimeout` 设置读取请求的最大时间。

### 写入超时

使用`Echo#*Server#WriteTimeout` 设置写入响应的最大时间。

### 验证

使用`Echo#Validator` 注册一个验证器，从而对请求负载执行数据验证。

[查看更多](https://echo.labstack.com/guide/request#validate-data)

### 自定义绑定

使用`Echo#Binder` 注册一个绑定器，从而绑定请求负载。

[查看更多](https://echo.labstack.com/guide/request/#custom-binder)

### 渲染

使用`Echo#Renderer` 注册一个渲染引擎，从而进行模板渲染。

### HTTP 错误处理

使用`Echo#HTTPErrorHandler` 注册一个 http 错误处理器。

[查看更多](https://echo.labstack.com/guide/error-handling)