+++
title = "Trailing Slash"
url = "/middleware/trailing-slash"
[menu.side]
  name = "Trailing Slash"
  parent = "middleware"
  weight = 20
+++

# Trailing Slash 中间件

### 添加尾部斜杠

在请求的 uri 后加上反斜杠

*使用*

```go
e := echo.New()
e.Pre(middleware.AddTrailingSlash())
```

### 去除尾部斜杠

在请求的 uri 后去除反斜杠

*用法*

```go
e := echo.New()
e.Pre(middleware.RemoveTrailingSlash())
```

### 自定义配置

*使用*

```go
e := echo.New()
e.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```

这个示例将会加上反斜杠，并且使用 `308 - StatusMovedPermanently` 重定向。

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
