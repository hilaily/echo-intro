+++
title = "尾部斜杠"
url = "/middleware/trailing-slash"
[menu.side]
  name = "尾部斜杠"
  parent = "middleware"
  weight = 20

+++

# Trailing Slash (尾部斜杠) 中间件

### 添加尾部斜杠

Add trailing slash 中间件会在在请求的 URI 后加上反斜杠

*用法*

```go
e := echo.New()
e.Pre(middleware.AddTrailingSlash())
```

### 去除尾部斜杠

Remove trailing slash 中间件在请求的 uri 后去除反斜杠

*用法*

```go
e := echo.New()
e.Pre(middleware.RemoveTrailingSlash())
```

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```

这个示例将向请求 URI 添加一个尾部斜杠，并使用 `301 - StatusMovedPermanently` 重定向。

###  配置

```go
TrailingSlashConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Status code to be used when redirecting the request.
  // Optional, but when provided the request is redirected using this code.
  RedirectCode int `json:"redirect_code"`
}
```

*默认配置*

```go
DefaultTrailingSlashConfig = TrailingSlashConfig{
  Skipper: defaultSkipper,
}
```
