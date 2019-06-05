---
title: 请求
url: guide/request
menu:
  side:
    parent: guide
    weight: 7
---

## 请求

### 数据绑定

可以使用 `Context#Bind(i interface{})` 将请求内容体绑定至 go 的结构体。默认绑定器支持基于 Content-Type 标头包含 application/json，application/xml 和 application/x-www-form-urlencoded 的数据。

下面是绑定请求数据到 `User` 结构体的例子

```go
// User
User struct {
  Name  string `json:"name" form:"name" query:"name"`
  Email string `json:"email" form:"email" query:"email"`
}
```

```go
// Handler
func(c echo.Context) (err error) {
  u := new(User)
  if err = c.Bind(u); err != nil {
    return
  }
  return c.JSON(http.StatusOK, u)
}
```

#### JSON 数据

```go
curl \
  -X POST \
  http://localhost:1323/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe","email":"joe@labstack"}'
```

#### Form 表单数据

```go
curl \
  -X POST \
  http://localhost:1323/users \
  -d 'name=Joe' \
  -d 'email=joe@labstack.com'
```

#### 查询参数 (Query Parameters)

```go
curl \
  -X GET \
  http://localhost:1323/users\?name\=Joe\&email\=joe@labstack.com
```

### 自定义绑定器

可以通过 `Echo#Binder` 注册自定义绑定器。

*示例*

```go
type CustomBinder struct {}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// 你也许会用到默认的绑定器
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}

	// 做你自己的实现

	return
}
```

### 检索数据

#### Form 表单数据

表单数据可以通过名称检索，使用 `Context#FormValue(name string)` 方法。

*示例*

```go
// Handler
func(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
}
```

```go
curl \
  -X POST \
  http://localhost:1323 \
  -d 'name=Joe'
```

当然，你也可以通过实现 `Echo#BindUnmarshaler` 接口来绑定自定义数据类型。

```go
type Timestamp time.Time

func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	*t = Timestamp(ts)
	return err
}
```

#### 查询参数 (Query Parameters)

查询参数可以通过名称获取，使用 `Context#QueryParam(name string)` 方法。

*示例*

```go
// Handler
func(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, name)
})
```

```go
curl \
  -X GET \
  http://localhost:1323\?name\=Joe
```

和表单数据一样，自定义数据类型依然通过 `Context#QueryParam(name string)` 进行绑定。

#### 路径参数 (Path Parameters)

路径参数可以通过 `Context#Param(name string) string` 进行检索。

*示例*

```go
e.GET("/users/:name", func(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
})
```

```go
$ curl http://localhost:1323/users/Joe
```

### 数据验证

Echo 没有内置的数据验证功能，但是可以通过 `Echo#Validator` 和[第三方库](https://github.com/avelino/awesome-go#validation)来注册一个数据验证器。

下面例子使用  [https://github.com/go-playground/validator](https://github.com/go-playground/validator)  所展示的框架来做验证：

```go
type (
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

```go
curl \
  -X POST \
  http://localhost:1323/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe","email":"joe@invalid-domain"}'
{"message":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"}
```

