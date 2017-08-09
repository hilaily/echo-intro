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

`Echo#Debug` 用来开启/关闭 debug 模式。Debug 模式下的日志级别是 **DEBUG**。

### 日志

#### 日志输出

`Echo#Logger.SetOutput(io.Writer)` 用于设置日志输出的位置，默认是 `os.Stdout`。

使用 `Echo#Logger.SetOutput(ioutil.Discard)` 或者 `Echo#Logger.SetLevel(log.OFF)`完全禁用日志。

#### 日志级别

`Echo#Logger.SetLevel(log.Lvl)` 用于设置日志级别，默认是 `OFF`。

可以使用的值：

- `DEBUG`
- `INFO`
- `WARN`
- `ERROR`
- `OFF`

#### 自定义日志

Echo 的日志实现了  `echo.Logger`  接口，你也可以使用 `Echo#Logger`实现该接口来注册一个自定义的日志。

### 自定义 Server

#### 使用 Echo#StartServer()

示例：

```go
s := &http.Server{
  Addr:         ":1323",
  ReadTimeout:  20 * time.Minute,
  WriteTimeout: 20 * time.Minute,
}
e.Logger.Fatal(e.StartServer(s))
```

#### 启动横幅
可以使用 `Echo#HideBanner` 关闭启动时候的横幅LOGO。

#### 自定义监听器
可以使用 `Echo#*Listener`启动一个自定义的 listener。
示例：
```go
l, err := net.Listen("tcp", ":1323")
if err != nil {
  e.Logger.Fatal(l)
}
e.Listener = l
e.Logger.Fatal(e.Start(""))
```

### 禁用 HTTP/2

`Echo#DisableHTTP2` 用于关闭 HTTP/2 协议。

### 读取超时

`Echo#*Server#ReadTimeout` 用于设置读取请求的最大时间。

### 写入超时

`Echo#*Server#WriteTimeout` 用于设置写入响应的最大时间。

### 验证

`Echo#Validator` 用来注册一个验证器，它可以在载入请求的时候做数据验证。

[查看更多](https://echo.labstack.com/guide/request#validate-data)

### 自定义绑定

`Echo#Binder` 用于注册一个绑定器来绑定请求。

[查看更多](https://echo.labstack.com/guide/request/#custom-binder)

### 渲染

`Echo#Renderer` 用来注册一个渲染引擎来渲染模版。

### HTTP 错误处理

`Echo#HTTPErrorHandler` 用于注册一个 http 错误处理器。

[查看更多](https://echo.labstack.com/guide/error-handling)