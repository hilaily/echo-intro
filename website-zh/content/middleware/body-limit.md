+++
title = "Body Limit"
url="/middleware/body-limit"
[menu.side]
  name = "BodyLimit(请求体限制)"
  parent = "middleware"
  weight = 4
+++

## BodyLimit(请求体限制) 中间件

BodyLimit 中间件用于设置允许的请求体的最大长度，如果请求体的大小超过了该值，则返回"413 － Request Entity Too Large"响应。
这个限制的判断取决于请求头的 `Content-Length` 和实际读取到的请求体内容两方面，尽可能的保证安全。

限制可以指定 `4x` 或者 `4xB`，x是"K, M, G, T, P"中的一个。

*用法*

```go
e := echo.New()
e.Use(middleware.BodyLimit("2M"))
```

## 自定义配置

*使用*

```go
e := echo.New()
e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{}))
```

*配置*

```go
BodyLimitConfig struct {
  // Skipper 定义了一个跳过中间间的函数
  Skipper Skipper

  // 请求体被允许的最大值，可以被指定为类似“4x”和“4xB”这样的值，
  // x 是 K，M，G，T，P 中的一个。
  Limit string `json:"limit"`
}
```

*默认配置*

```go
DefaultBodyLimitConfig = BodyLimitConfig{
  Skipper: defaultSkipper,
}
```









