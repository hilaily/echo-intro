+++
title = "Gzip"
url = "gzip"
[menu.side]
  name = "Gzip"
  parent = "middleware"
  weight = 5
+++

## Gzip Middleware

Gzip 中间件使用 gzip 压缩方案来对HTTP响应进行压缩。 

### 配置

```go
GzipConfig struct {
  // Gzip compression level.
  // Optional. Default value -1.
  Level int
}
```

### 默认配置

```go
DefaultGzipConfig = GzipConfig{
  Level: -1,
}
```

*Usage*

`e.Use(middleware.Gzip())`

### 自定义配置

*Usage*

```go
e := echo.New()
e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
  Level: 5
}))
```
