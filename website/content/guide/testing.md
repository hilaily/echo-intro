---
title: Testing
menu:
  side:
    parent: guide
    weight: 9
---

### Testing Handler

`GET` `/users/:id`

Handler below retrieves user by ID from the database. If user is not found it returns
`404` error with a message.

- ID comes from path parameter.
- DB comes from `net.context`.

`handler.go`

```go
package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	db := c.NetContext().Value("db").(map[string]string)
	user := db[id]
	if user == "" {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.String(http.StatusOK, db[id])
}
```

Let's write a test case (positive and negative) for our handler by creating a mock
context using `Echo#NewContext()` and load it with the following properties:

- Path parameters.
- Mock DB using `net.context`.

The negative test is accomplished by setting the user ID to a invalid value.

`handler_test.go`

```go
package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/context"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

var (
	mockDB = map[string]string{
		"1": "walle",
		"2": "bolt",
		"3": "tintin",
	}
)

func TestGet(t *testing.T) {
	// Setup
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.SetNetContext(context.WithValue(context.Background(), "db", mockDB))

	// Positive test
	if err := getUser(c); err != nil {
		t.Error(err)
	}

	// Negative test
	c.SetParamValues("4")
	if err := getUser(c); err == nil {
		t.Error(err)
	}
}
```

### Testing Middleware

*TBD*
