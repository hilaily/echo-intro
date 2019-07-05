+++
title = "请求ID"
url = "/middleware/request-id"
[menu.side]
  name = "请求ID"
  parent = "middleware"
  weight = 16

+++

## Request ID (请求ID) 中间件

Request ID 中间件为请求生成唯一的 ID。

*用法*

```go
e.Use(middleware.RequestID())
```

### 自定义配置

*用法*

```go
e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
  Generator: func() string {
    return customGenerator()
  },
}))
```

### 配置

```go
RequestIDConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Generator defines a function to generate an ID.
  // Optional. Default value random.String(32).
  Generator func() string
}
```

*默认配置*

```go
DefaultRequestIDConfig = RequestIDConfig{
  Skipper:   DefaultSkipper,
  Generator: generator,
}
```

