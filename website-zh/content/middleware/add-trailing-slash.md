+++
title = "添加结尾斜杠"
url = "add-trailing-slash"
[menu.side]
  name = "AddTrailingSlash(添加结尾斜杠)"
  parent = "middleware"
  weight = 2
+++

## AddTrailingSlash(添加结尾斜杠) 中间件

AddTrailingSlash 中间件用语在请求 URI 后添加一个斜杠。

### 配置

```go
TrailingSlashConfig struct {
  // Status code to be used when redirecting the request.
  // Optional, but when provided the request is redirected using this code.
  RedirectCode int
}
```

*用法*

```go
e := echo.New()
e.Pre(middleware.AddTrailingSlash())
```

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.AddTrailingSlashWithConfig(TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```

这将在请求的 URI 的末尾添加一个'/'并且做 `StatusMovedPermanenty` 跳转。
