---
title: 错误处理
menu:
  side:
    parent: guide
    weight: 9
---

## 错误处理

Echo 支持从中间件或者action返回 HTTP 错误 集中处理。

- 在统一的地方记录日志
- 返回自定义的 HTTP 响应

例如 一个基本的身份验证中间件验证失败返回 
`401 - Unauthorized` 错误, 终止了当前的 HTTP 请求。

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Use(func(c echo.Context) error {
		// 从 HTTP 请求中取出 credentials(一般为cookie)然后验证

		// 验证失败
		return echo.NewHTTPError(http.StatusUnauthorized)
	})
	e.GET("/welcome", welcome)
	e.Run(":1323")
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}
```

查看 [HTTPErrorHandler](/guide/customization#http-error-handler) 怎样处理的。
