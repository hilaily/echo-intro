---
title: Index
---

## 简易高效的 Go(Golang) 语言 HTTP 框架。比其他框架快速10倍。

## 功能概览

- 优化的 HTTP 路由
- Build robust and scalable RESTful APIs.
- Run with standard HTTP server or FastHTTP server.
- Group APIs.
- Extensible middleware framework.
- Define middleware at root, group or route level.
- Data binding for JSON, XML and form payload.
- Handy functions to send variety of HTTP responses.
- Centralized HTTP error handling.
- Template rendering with any template engine.
- Define your format for the logger.
- Highly customizable.

## 性能

- 环境:
	- Go 1.6
	- wrk 4.0.0
	- 2 GB, 2 Core (DigitalOcean)
- 测试方案: https://github.com/vishr/web-framework-benchmark
- 日期: 4/4/2016

<img width="600" height="371" src="https://o8l6oohcu.qnssl.com/go-echo/echo.png">

## 快速开始

### 安装

```sh
$ go get github.com/labstack/echo/...
```

### 编写 Hello, World!

创建 `server.go` 文件

```go
package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Run(standard.New(":1323"))
}
```

开启服务

```sh
$ go run server.go
```

在浏览器访问 [http://localhost:1323](http://localhost:1323) 然后你就能在页面上看到 
Hello, World! 
### 路由

```go
e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
```

### URL路径参数

```go
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
}
```

### 请求参数

`/show?team=x-men&member=wolverine`

```go
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
}
```

### 表单 `application/x-www-form-urlencoded`

`POST` `/save`

name | value
:--- | :---
name | Joe Smith
email | joe@labstack.com

```go
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
}
```

### 表单 `multipart/form-data`

`POST` `/save`

name | value
:--- | :---
name | Joe Smith
email | joe@labstack.com
avatar | avatar

```go
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you!</b>")
}
```

### 处理请求

- Bind `JSON` or `XML` or `form` payload into Go struct based on `Content-Type` request header.
- Render response as `JSON` or `XML` with status code.

```go
type User struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
}

e.POST("/users", func(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
	// or
	// return c.XML(http.StatusCreated, u)
})
```

### 静态资源

定义`/static/*`目录为静态资源文件目录

```go
e.Static("/static", "static")
```

##### [更多...](https://echo.labstack.com/guide/static-files)

### [模板渲染](https://echo.labstack.com/guide/templates)

### 中间件

```go
// Root level middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())

// Group level middleware
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(username, password string) bool {
	if username == "joe" && password == "secret" {
		return true
	}
	return false
}))

// Route level middleware
track := func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /users")
		return next(c)
	}
}
e.GET("/users", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users")
}, track)
```

#### echo 自带的中间件

中间件| 描述
:--- | :---
[BodyLimit]({{< ref "middleware/body-limit.md">}}) | Limit request body
[Logger]({{< ref "middleware/logger.md">}}) | Log HTTP requests
[Recover]({{< ref "middleware/recover.md">}}) | Recover from panics
[Gzip]({{< ref "middleware/gzip.md">}}) | Send gzip HTTP response
[BasicAuth]({{< ref "middleware/basic-auth.md">}}) | HTTP basic authentication
[JWTAuth]({{< ref "middleware/jwt.md">}}) | JWT authentication
[Secure]({{< ref "middleware/secure.md">}}) | Protection against attacks
[CORS]({{< ref "middleware/cors.md">}}) | Cross-Origin Resource Sharing
[CSRF]({{< ref "middleware/csrf.md">}}) | Cross-Site Request Forgery
[Static]({{< ref "middleware/static.md">}}) | Serve static files
[AddTrailingSlash]({{< ref "middleware/add-trailing-slash.md">}}) | Add trailing slash to the request URI
[RemoveTrailingSlash]({{< ref "middleware/remove-trailing-slash.md">}}) | Remove trailing slash from the request URI
[MethodOverride]({{< ref "middleware/method-override.md">}}) | Override request method

#### 第三方中间件

中间件 | 描述
:--- | :---
[echoperm](https://github.com/xyproto/echoperm) | Keeping track of users, login states and permissions.
[echopprof](https://github.com/mtojek/echopprof) | Adapt net/http/pprof to labstack/echo.

##### [更多...](https://echo.labstack.com/guide/middleware)

### 接下来

- Head over to [guide](https://echo.labstack.com/guide/installation)
- Browse [recipes](https://echo.labstack.com/recipes/hello-world)

### 帮助

- [Hop on to chat](https://gitter.im/labstack/echo)
- [Open an issue](https://github.com/labstack/echo/issues/new)

## 支持我们

- ☆ the project
- [Donate](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=JD5R56K84A8G4&lc=US&item_name=LabStack&item_number=echo&currency_code=USD&bn=PP-DonationsBF:btn_donate_LG.gif:NonHosted)
- 🌐 spread the word
- [Contribute](#contribute:d680e8a854a7cbad6d490c445cba2eba) to the project

## 贡献

**用 issues 完成所有工作**


- 提交 issues
- 发 PR 前的讨论
- 提出新功能或者优化的建议
- 完善／修复 文档

## 开发人员

- [Vishal Rana](https://github.com/vishr) - 作者
- [Nitin Rana](https://github.com/nr17) - 顾问
- [其他贡献者](https://github.com/labstack/echo/graphs/contributors)

## License

[MIT](https://github.com/labstack/echo/blob/master/LICENSE)
