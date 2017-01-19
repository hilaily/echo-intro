+++
title = "Gzip"
url = "/middleware/gzip"
[menu.side]
  name = "Gzip"
  parent = "middleware"
  weight = 5
+++

## Gzip Middleware

Gzip 中间件使用 gzip 压缩方案来对HTTP响应进行压缩。 

*使用*

```go
e.Use(middleware.Gzip())
```

### 自定义配置

*使用*

```go
e := echo.New()
e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
  Level: 5,
}))
```

### 配置

```go
GzipConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Gzip compression level.
  // Optional. Default value -1.
  Level int `json:"level"`
}
```

### 默认配置

```go
DefaultGzipConfig = GzipConfig{
  Skipper: defaultSkipper,
  Level:   -1,
}
```
