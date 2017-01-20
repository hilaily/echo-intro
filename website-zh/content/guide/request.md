---
title: HTTP 请求
url: guide/request
menu:
  side:
    parent: guide
    weight: 9
---

## HTTP 请求

### 数据绑定

使用 `Context#Bind(i interface{})` 绑定一个请求内容体到 go 的结构体。默认的绑定器支持解析 Content-Type 是 application/json，application/xml 和 application/x-www-form-urlencoded 的数据。

下面是绑定请求数据到 User 结构体的例子

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

#### From 表单数据

```go
curl \
  -X POST \
  http://localhost:1323/users \
  -d 'name=Joe' \
  -d 'email=joe@labstack.com'
```

### url 请求数据 (Query Parameters)

```go
curl \
  -X GET \
  http://localhost:1323/users\?name\=Joe\&email\=joe@labstack.com
```

### 自定义绑定器

可以通过 `Echo#Binder`自定义绑定器。

示例

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

### 读取数据

#### Form 表达数据

表单数据可以通过名称读取，使用这个方法 `Context#FormValue(name string)`。

示例

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

你可以实现 `Echo#BindUnmarshaler` 接口去去绑定自己的数据结构。

```go
type Timestamp time.Time

func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	*t = Timestamp(ts)
	return err
}
```

### url 请求数据 (Query Parameters)

url 请求参数可以通过名称获取，使用 `Context#QueryParam(name string)` 方法。

示例

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

和表单数据一样，自定义数据也可以通过 `Context#QueryParam(name string)` 绑定。

### url 请求参数 (Path Parameters)

url 请求参数可以通过 `Context#Param(name string) string` 获取。

示例

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

Echo 没有内置数据验证功能，但是可以通过 `Echo#Validator` 和[第三方库](https://github.com/avelino/awesome-go#validation)自己注册一个数据验证器。

下面例子  [https://github.com/go-playground/validator](https://github.com/go-playground/validator)  使用做验证

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

