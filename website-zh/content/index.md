---
title: Index
url: index
---

## 快速开始

### 安装

```sh
$ go get -u github.com/labstack/echo/...
```

### 编写 Hello, World!

创建 `server.go` 文件

```go
package main

import (
	"net/http"
    
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

启动服务
```sh
$ go run server.go
```
用浏览器访问 [http://localhost:1323](http://localhost:1323) 然后就能在页面上看到 `Hello, World!` 
### 路由

```go
e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
```

### URL路径参数

```go
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID 来自于url `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
```

用浏览器访问 [http://localhost:1323/users/Joe](http://localhost:1323) 然后就能在页面上看到 `Joe` 

### 请求参数

`/show?team=x-men&member=wolverine`

```go
// e.GET("/show", show)
func show(c echo.Context) error {
	// 从请求参数里获取 team 和 member 的值
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
```
从浏览器访问 `http://localhost:1323/show?team=x-men&member=wolverine` 可以看到页面上显示"team:x-men, member:wolverine"

### 表单 `application/x-www-form-urlencoded`

`POST` `/save`

name | value
:--- | :---
name | Joe Smith
email | joe@labstack.com

```go
// e.POST("/save", save)
func save(c echo.Context) error {
	// 获取 name 和 email 的值
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}
```
在命令行里执行下面的语句
```bash
$ curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
```
控制台会输出`name:Joe Smith, email:joe@labstack.com`

### 表单 `multipart/form-data`

`POST` `/save`

name | value
:--- | :---
name | Joe Smith
avatar | avatar

```go
func save(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
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

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}
```
命令行执行下面语句
```bash
$ curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/save
//output => <b>Thank you! Joe Smith</b>
```
使用以下命令查看刚刚上传的图片

```bash
cd <project directory>
ls avatar.png
// => avatar.png
```

### 处理请求

- 根据 Content-Type 请求标头将 `json`，`xml`，`form` 或 `query` 负载绑定到 Go 结构中。
- 通过状态码将响应渲染为 `json` 或者 `xml` 格式。

```go
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
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

下面的代码定义`/static/*`目录为静态资源文件目录

```go
e.Static("/static", "static")
```

##### [了解更多](http://go-echo.org/guide/static-files)

### [模板渲染](http://go-echo.org/guide/templates)

### 中间件

```go
// Root level middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())

// Group level middleware
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (error, bool) {
  if username == "joe" && password == "secret" {
    return nil, true
  }
  return nil, false
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
##### [了解更多](https://echo.labstack.com/middleware)