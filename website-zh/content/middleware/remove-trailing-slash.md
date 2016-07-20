+++
title = "Remove Trailing Slash(移除结尾斜杠)"
[menu.side]
  name = "RemoveTrailingSlash(移除结尾斜杠)"
  parent = "middleware"
  weight = 5
+++

## RemoveTrailingSlash(移除结尾斜杠) 中间件

RemoveTrailingSlash 中间件 会从请求 URI 里移除到最后的斜杠。

### 配置

```go
TrailingSlashConfig struct {
  // 当重定向一个请求的时候会用到状态码
  // 可选，但是当提供的请求需要跳转则要使用到。
  RedirectCode int
}
```

*用法*

```go
e := echo.New()
e.Pre(middleware.RemoveTrailingSlash())
```

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.RemoveTrailingSlashWithConfig(TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```

这样会移除掉最后的斜杠，同时带上`StatusMovedPermanently`(即301)重定向。

