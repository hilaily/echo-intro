---
title: 路由
url: guide/routing
menu:
  side:
    parent: guide
    weight: 9
---

## 路由

Echo 的路由基于 [radix tree](http://en.wikipedia.org/wiki/Radix_tree) ，它让路由的查询非常快。路由使用了 [sync pool](https://golang.org/pkg/sync/#Pool) 来重复利用内存并且几乎达到了零内存占用。

路由可以通过特定的 HTTP 方法，url 路径和一个匹配的 handler 来注册。例如，下面的代码则展示了一个注册路由的例子，访问方式为 `Get`，访问路径为 `/hello`，处理结果是返回输出 `Hello World` 的响应。

```go
// 业务处理
func hello(c echo.Context) error {
  	return c.String(http.StatusOK, "Hello, World!")
}

// 路由
e.GET("/hello", hello)
```

你可以用 `Echo.Any(path string, h Handler)` 来为所有的 HTTP 方法发送注册 handler；如果只想为某些方法注册的话则需要用 `Echo.Match(methods []string, path string, h Handler)`。

Echo 通过 `func(echo.Context) error` 定义 handler 方法， `echo.Context` 已经内嵌了 HTTP 请求和响应的接口。

### 匹配所有

匹配零个或多个字符的路径。例如， `/users/*` 将会匹配；

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

### 路由命名
每个注册路由返回一个 `Route` 对象。它可以在注册了之后用来给路由命名。比如：
```go
routeInfo := e.GET("/users/:id", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users/:id")
})
routeInfo.Name = "user"

// 或者这样写
e.GET("/users/new", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users/new")
}).Name = "newuser"
```
当你需要在模版生成 uri 但是又无法获取路由的引用或者多个路由使用的相同的处理器(handler)的时候，使用命名路由就会比较方便。

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
除了 `Echo#URI`，还有 `Echo#Reverse(name string, params ...interface{})` 方法用来根据路由名生成 uri。
比如下面的代码中使用 `Echo#Reverse("foobar", 1234)` 就会生成 `/users/1234`。
```go
// Handler
h := func(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// Route
e.GET("/users/:id", h).Name = "foobar"
```

### 路由列表
`Echo#Routes() []*Route` 会根据定义的顺序列出所有已经注册的路由。每一个路由包含 http 方法，路径和对应的处理器。
示例
```go
// Handlers
func createUser(c echo.Context) error {
}

func findUser(c echo.Context) error {
}

func updateUser(c echo.Context) error {
}

func deleteUser(c echo.Context) error {
}

// Routes
e.POST("/users", createUser)
e.GET("/users", findUser)
e.PUT("/users", updateUser)
e.DELETE("/users", deleteUser)
```
用下面的代码你可以输出所有的路由信息到 json 文件。
```go
data, err := json.MarshalIndent(e.Routes(), "", "  ")
if err != nil {
	return err
}
ioutil.WriteFile("routes.json", data, 0644)
```
```json
[
  {
    "method": "POST",
    "path": "/users",
    "handler": "main.createUser"
  },
  {
    "method": "GET",
    "path": "/users",
    "handler": "main.findUser"
  },
  {
    "method": "PUT",
    "path": "/users",
    "handler": "main.updateUser"
  },
  {
    "method": "DELETE",
    "path": "/users",
    "handler": "main.deleteUser"
  }
]
```


