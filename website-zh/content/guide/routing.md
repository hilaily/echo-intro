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

路由线路可以通过制定的HTTP方法，路径和一个匹配的handler来注册。例如，下面的代码则展示了一个路由线路的注册的例子，访问方式为 `Get`，
访问路径为 `/hello`，处理结果是返回输出 `Hello World` 的响应。

```go
// 业务处理
func hello(c echo.Context) error {
  	return c.String(http.StatusOK, "Hello, World!")
}

// 路由
e.GET("/hello", hello)
```

你可以用 `Echo.Any(path string, h Handler)` 来为所有的HTTP方法发送注册handler；如果只想为某些方法注册的话则需要用 `Echo.Match(methods []string, path string, h Handler)`。

Echo 通过 `func(echo.Context) error` 定义handler方法， `echo.Context` 已经内嵌了 HTTP 请求和响应的接口。

### 匹配所有

匹配零个或多个字符的路径。例如， `/users/*` 将会匹配:

- `/users/`
- `/users/1`
- `/users/1/files/1`
- `/users/anything...`

### 路径匹配顺序

- Static (固定路径)
- Param (参数路径)
- Match any (匹配所有)

#### 示例

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

上面定义的路由将按下面的优先级顺序匹配:

- `/users/new`
- `/users/:id`
- `/users/1/files/*`

> 定义路由的顺序没有限制。

### 组路由

`Echo#Group(prefix string, m ...Middleware) *Group`

拥有相同前缀的路由可以通过中间件定义一个子路由来化为一组。
除了一些特殊的中间件，组路由也会继承父中间件。
在组路由里使用中间件可以用`Group.Use(m ...Middleware)`。
组路由可以嵌套。

下面的代码，我们创建了一个 admin 组，使所有的 `/admin/*` 都要求 HTTP 基本认证。

```go
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(username, password string) bool {
	if username == "joe" && password == "secret" {
		return true
	}
	return false
}))
```

### 构造URI 

`Echo.URI` 可以用来在任何业务处理代码里生成带有特殊参数的URI。
这样当你重构自己的应用的时候，可以很方便的集中处理代码。

下面的代码中 `e.URI(h, 1)` 将生成`/users/1`。

```go
// 业务处理
h := func(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// 路由
e.GET("/users/:id", h)
```
