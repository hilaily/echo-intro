---
title: 测试
url: guide/testing
menu:
  side:
    parent: guide
    weight: 10
---

## 测试

### 测试业务逻辑

`GET` `/users/:id`

这个业务是根据用户的 id 从数据库取到该用户数据，如果用户不存在则返回`404`和提示语句。

#### 创建 User

`POST` `/users`

- 接受 JSON 格式的数据。
- 创建成功返回 `201 - Created`。
- 发生错误返回 `500 - Internal Server Error`。

#### 获取 User

`GET` `/users/:email`

- 获取成功返回 `200 - OK`
- 发生错误返回 `404 - Not Found` if user is not found otherwise `500 - Internal Server Error`

`handler.go`

```go
package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	User struct {
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email"`
	}
	handler struct {
		db map[string]*User
	}
)

func (h *handler) createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *handler) getUser(c echo.Context) error {
	email := c.Param("email")
	user := h.db[email]
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}
```

`handler_test.go`

```go
package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*User{
		"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
	}
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestCreateUser(t *testing.T) {
	// 设置
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/users", strings.NewReader(userJSON))
	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		h := &handler{mockDB}

		// 断言
		if assert.NoError(t, h.createUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, userJSON, rec.Body.String())
		}
	}
}

func TestGetUser(t *testing.T) {
	// 设置
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/users/:email")
	c.SetParamNames("email")
	c.SetParamValues("jon@labstack.com")
	h := &handler{mockDB}

	// 断言
	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
```

#### 使用 Form 表单提交

```go
f := make(url.Values)
f.Set("name", "Jon Snow")
f.Set("email", "jon@labstack.com")
req, err := http.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
```

#### 设置 URL 参数 

```go
c.SetParamNames("id", "email")
c.SetParamValues("1", "jon@labstack.com")
```

#### 设置请求参数

```go
q := make(url.Values)
q.Set("email", "jon@labstack.com")
req, err := http.NewRequest(echo.POST, "/?"+q.Encode(), nil)
```

### 测试中间件

*待更新*

你可以在[这里](https://github.com/labstack/echo/tree/master/middleware)查看框架自带中间件的测试代码。
