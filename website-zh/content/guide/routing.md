---
title: 路由
menu:
  side:
    parent: guide
    weight: 6
---

## 路由

Echo 的路由[性能非常高]({{< ref "index.md#performance">}}) 而且扩展性强。
该路由基于 [radix tree](http://en.wikipedia.org/wiki/Radix_tree) 这个数据结构。它让路由的查询非常快。路由使用了
 [sync pool](https://golang.org/pkg/sync/#Pool) 来重复利用内存并且几乎达到了零内存占用。

Routes can be registered by specifying HTTP method, path and a matching handler.
For example, code below registers a route for method `GET`, path `/hello` and a
handler which sends `Hello, World!` HTTP response.

```go
// Handler
func hello(c echo.Context) error {
  	return c.String(http.StatusOK, "Hello, World!")
}

// Route
e.GET("/hello", hello)
```

You can use `Echo.Any(path string, h Handler)` to register a handler for all HTTP methods.
If you want to register it for some methods use `Echo.Match(methods []string, path string, h Handler)`.

Echo defines handler function as `func(echo.Context) error` where `echo.Context` primarily
holds HTTP request and response interfaces.

### Match-any

Matches zero or more characters in the path. For example, pattern `/users/*` will
match:

- `/users/`
- `/users/1`
- `/users/1/files/1`
- `/users/anything...`

### Path matching order

- Static
- Param
- Match any

#### Example

```go
e.GET("/users/:id", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users/:id")
})

e.GET("/users/new", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users/new")
})

e.GET("/users/1/files/*", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users/1/files/*")
})
```

Above routes would resolve in the following order:

- `/users/new`
- `/users/:id`
- `/users/1/files/*`

> Routes can be written in any order.

### Group

`Echo#Group(prefix string, m ...Middleware) *Group`

Routes with common prefix can be grouped to define a new sub-router with optional
middleware. In addition to specified middleware group also inherits parent middleware.
To add middleware later in the group you can use `Group.Use(m ...Middleware)`.
Groups can also be nested.

In the code below, we create an admin group which requires basic HTTP authentication
for routes `/admin/*`.

```go
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(username, password string) bool {
	if username == "joe" && password == "secret" {
		return true
	}
	return false
}))
```

### URI building

`Echo.URI` can be used to generate URI for any handler with specified path parameters.
It's helpful to centralize all your URI patterns which ease in refactoring your
application.

`e.URI(h, 1)` will generate `/users/1` for the route registered below

```go
// Handler
h := func(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// Route
e.GET("/users/:id", h)
```
