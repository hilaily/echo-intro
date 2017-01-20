---
title: 错误处理
url : guide/error-handling
menu:
  side:
    parent: guide
    weight: 11
---

## 错误处理

Echo 倾向于从中间件或者action返回 HTTP 错误集中处理。它也允许我们去在统一的地方记录日志或者返回自定义的 HTTP 响应给其他的服务。

例如 一个基本的身份验证中间件验证失败返回 `401 - Unauthorized` 错误, 终止了当前的 HTTP 请求。

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the credentials from HTTP request header and perform a security
			// check

			// For invalid credentials
			return echo.NewHTTPError(http.StatusUnauthorized)

			// For valid credentials call next
			// return next(c)
		}
	})
	e.GET("/", welcome)
	e.Logger.Fatal(e.Start(":1323"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}
```

查看 [HTTPErrorHandler](https://echo.labstack.com/guide/customization#http-error-handler) 怎样处理的。
