---
title: Index
url: index
---

## 简易高效的 Go(Golang) 语言 HTTP 框架。比其他框架快速10倍。

## 功能概览

- 优化的 HTTP 路由。
- 创建可靠并可伸缩的RESTful API。
- 行于标准的HTTP服务器或FastHTTP服务器。
- 组 APIs.
- 可扩展的middleware框架。
- Define middleware at root, group or route level.
- 为JSON, XML进行数据绑定，产生负荷。
- 提供便捷的方法来发送各种HTTP相应。
- 对HTTP错误进行集中处理。
- Template rendering with any template engine.
- 定义属于你的日志格式。
- 高度个性化。

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
	// User ID 来自于url `users/:id`
	id := c.Param("id")
}
```

### 请求参数

`/show?team=x-men&member=wolverine`

```go
func show(c echo.Context) error {
	// 从请求参数里获取 team 和 member 的值
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
	// 获取 name 和 email 的值
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

- 在数据结构体里设置`JSON` 或 `XML` 或 `form` 直接匹配请求头的 `Content-Type`。 
- 结合响应状态将响应渲染为`JSON` 或者 `XML`。

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
	// 或者
	// return c.XML(http.StatusCreated, u)
})
```

### 静态资源

定义`/static/*`目录为静态资源文件目录

```go
e.Static("/static", "static")
```

##### [更多...](http://go-echo.org/guide/static-files)

### [模板渲染](http://go-echo.org/guide/templates)

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
[BodyLimit]({{< ref "middleware/body-limit.md">}}) | 限制请求体
[Logger]({{< ref "middleware/logger.md">}}) | HTTP请求日志
[Recover]({{< ref "middleware/recover.md">}}) | Recover from panics
[Gzip]({{< ref "middleware/gzip.md">}}) | 发送gzip压缩的 HTTP 响应
[BasicAuth]({{< ref "middleware/basic-auth.md">}}) | HTTP基本身份认证
[JWTAuth]({{< ref "middleware/jwt.md">}}) | JWT身份认证
[Secure]({{< ref "middleware/secure.md">}}) | 防止攻击
[CORS]({{< ref "middleware/cors.md">}}) | 跨源资源共享
[CSRF]({{< ref "middleware/csrf.md">}}) | 跨站请求伪造
[MethodOverride]({{< ref "middleware/method-override.md">}}) | 覆盖请求方法
#### 第三方中间件

中间件 | 描述
:--- | :---
[echoperm](https://github.com/xyproto/echoperm) | 对用户、登陆状态与权限进行追踪。
[echopprof](https://github.com/mtojek/echopprof) | Adapt net/http/pprof to labstack/echo.

##### [学习更多...](http://go-echo.org/guide/middleware)

### 接下来

- 回到 [guide](http://go-echo.org/guide/installation)
- 浏览 [recipes](http://go-echo.org/recipes/hello-world)

### 帮助

- [在线询问](https://gitter.im/labstack/echo)
- [新建一个issue](https://github.com/labstack/echo/issues/new)

## 支持我们

- ☆ 点赞
- [打赏](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=JD5R56K84A8G4&lc=US&item_name=LabStack&item_number=echo&currency_code=USD&bn=PP-DonationsBF:btn_donate_LG.gif:NonHosted)
- 🌐 spread the word
- [改进](#contribute:d680e8a854a7cbad6d490c445cba2eba) 这个程序

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
