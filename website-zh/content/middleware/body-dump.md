+++
title = "BodyDump 请求体转储"
url="/middleware/body-dump"
[menu.side]
  name = "BodyDump 请求体转储"
  parent = "middleware"
  weight = 4

+++

## Body Dump (请求体转储) 中间件

Body dump 中间件通常在调试 / 记录的情况下被使用，它可以捕获请求并调用已注册的处理程序 (handler) 响应有效负载。然而，当您的请求 / 响应有效负载很大时（例如上传 / 下载文件）需避免使用它；但如果避免不了，可在 skipper 函数中为端点添加异常。

*用法*

```go
e := echo.New()
e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
}))
```

## 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{}))
```

## 配置

```go
BodyDumpConfig struct {
  // Skipper 定义了一个跳过中间件的函数
  Skipper Skipper

  // Handler 接收请求和响应有效负载
  // Required.
  Handler BodyDumpHandler
}
```

*默认配置*

```go
DefaultBodyDumpConfig = BodyDumpConfig{
  Skipper: DefaultSkipper,
}
```