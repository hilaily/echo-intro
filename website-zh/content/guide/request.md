---
title: HTTP 请求
menu:
  side:
    parent: guide
    weight: 8
---

## HTTP 请求

### Action 路径处理

`Context#Path()` 返回在路由注册的Action的路径，它可以被中间件使用来得到访问的目的地址。

*例如*

```go
e.Use(func(c echo.Context) error {
    println(c.Path()) // Prints `/users/:name`
    return nil
})
e.GET("/users/:name", func(c echo.Context) error) {
    return c.String(http.StatusOK, name)
})
```

### golang.org/x/net/context 对象

`echo.Context` 内嵌了 `context.Context` 接口，所以`context.Context` 的所有方法`echo.Context`都支持。

*例如*

```go
e.GET("/users/:name", func(c echo.Context) error) {
    c.SetNetContext(context.WithValue(nil, "key", "val"))
    // Pass it down...
    // Use it...
    val := c.Value("key").(string)
    return c.String(http.StatusOK, name)
})
```

### URL 参数

URL 参数可以用参数名 `Context#Param(name string) string` 和参数索引(序号) `Context#P(i int) string` 取得。通过参数索引的方式性能会稍微好一点。

*例如*

```go
e.GET("/users/:name", func(c echo.Context) error {
	// By name
	name := c.Param("name")

	// By index
	name := c.P(0)

	return c.String(http.StatusOK, name)
})
```

```sh
$ curl http://localhost:1323/users/joe
```

### 请求参数

请求参数可以通过参数名获取 `Context#QueryParam(name string)`。

*例如*

```go
e.GET("/users", func(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, name)
})
```

```sh
$ curl -G -d "name=joe" http://localhost:1323/users
```

### 表单参数

表单参数可以使用参数名获取 `Context#FormValue(name string)`。

*例如*

```go
e.POST("/users", func(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
})
```

```sh
$ curl -d "name=joe" http://localhost:1323/users
```
