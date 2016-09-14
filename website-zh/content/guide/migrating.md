---
title: 迁移
url : migrating
menu:
  side:
    parent: guide
    weight: 2
---

## 从 v1 迁移

### 更新日志

- 85%的API都保持和以前一样
- `Engine` 接口抽象的实现了 `HTTP` 服务，允许我们在 GO 标准库上使用 HTTP 服务。它现在支持标准服务和 [fasthttp](https://github.com/valyala/fasthttp) 服务。
- Context, Request 和 Response 转换成了接口[详细...](https://github.com/labstack/echo/issues/146)。
- 处理程序签名改为了 `func (c echo.Context) error`。
- Dropped auto wrapping of handler and middleware to enforce compile time check.
- 在路由之前运行的中间件的 API，现在不需要`Echo#Hook` 了。
- 能够在 route 级别定义中间件了。
- `Echo#HTTPError` 暴露了它的`Code` 和 `Message` 值。
- 可选的指定中间件日志和默认日志的格式。

#### API

v1 | v2
--- | ---
`Context#Query()` | `Context#QueryParam()`
`Context#Form()`  | `Context#FormValue()`

### FAQ（常见问题）

Q. 如何通过接口访问原始对象

A. 如果你需要的话...

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

Q. 如何使用标准的业务处理和中间件

A.如下所示

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// 标准中间件
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("standard middleware")
		next.ServeHTTP(w, r)
	})
}

// 标准业务处理
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

- 浏览刚刚转换到 v2 的[recipes](/recipes/hello-world)。
- 阅读文档，深入理解测试实例。
