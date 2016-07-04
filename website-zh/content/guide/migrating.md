---
title: 迁移
menu:
  side:
    parent: guide
    weight: 2
---

## 从 v1 迁移

### 更新日志

- 85%的API都保持和以前一样
- `Engine` interface to abstract `HTTP` server implementation, allowing
us to use HTTP servers beyond Go standard library. It currently supports standard and [fasthttp](https://github.com/valyala/fasthttp) server.
- Context, Request 和 Response 转换成了 interface [详细...](https://github.com/labstack/echo/issues/146)

- Handler signature is changed to `func (c echo.Context) error`.（处理程序签名改为了 `func (c echo.Context) error`）
- Dropped auto wrapping of handler and middleware to enforce compile time check.
- APIs to run middleware before or after the router, which doesn't require `Echo#Hook` API now.
- Ability to define middleware at route level.
- `Echo#HTTPError` exposed it's fields `Code` and `Message`.
- Option to specify log format in logger middleware and default logger.

#### API

v1 | v2
--- | ---
`Context#Query()` | `Context#QueryParam()`
`Context#Form()`  | `Context#FormValue()`

### FAQ（常见问题）

Q. How to access original objects from interfaces?（如何通过接口访问原始对象）

A. Only if you need to...（如果你需要的话……）

```go
// `*http.Request`
c.Request().(*standard.Request).Request

// `*http.URL`
c.Request().URL().(*standard.URL).URL

// Request `http.Header`
c.Request().Header().(*standard.Header).Header

// `http.ResponseWriter`
c.Response().(*standard.Response).ResponseWriter

// Response `http.Header`
c.Response().Header().(*standard.Header).Header
```

Q. 如何使用标准的handler和middleware

A.如下所示

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// Standard middleware
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("standard middleware")
		next.ServeHTTP(w, r)
	})
}

// Standard handler
func handler(w http.ResponseWriter, r *http.Request) {
	println("standard handler")
}

func main() {
	e := echo.New()
	e.Use(standard.WrapMiddleware(middleware))
	e.GET("/", standard.WrapHandler(http.HandlerFunc(handler)))
	e.Run(standard.New(":1323"))
}
```

### 继续?

- Browse through [recipes](/recipes/hello-world) freshly converted to v2.（浏览[recipes](/recipes/hello-world)）
- 阅读文档，深入理解测试实例。
