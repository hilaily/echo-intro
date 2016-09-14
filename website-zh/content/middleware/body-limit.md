+++
title = "Body Limit"
url="body-limit"
[menu.side]
  name = "BodyLimit(请求体限制)"
  parent = "middleware"
  weight = 4
+++

## BodyLimit(请求体限制) 中间件

BodyLimit 中间件用于设置允许的请求体的最大长度，如果请求的大小超过了该值，则返回"413 － Request Entity Too Large"响应。
这个限制的判断取决于请求头的`Content-Length`和实际读取到的请求体内容两方面，是非常安全的。

限制可以指定`4x`或者`4xB`，x是"K, M, G, T, P"中的一个。

*用法*

```go
e := echo.New()
e.Use(middleware.BodyLimit("2M"))
```
