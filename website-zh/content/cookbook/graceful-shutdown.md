---
title: Graceful Shutdown
url : /cookbook/graceful-shutdown
menu:
  side:
    parent: cookbook
    weight: 16
---

## 平滑关闭

### 使用 [grace](https://github.com/facebookgo/grace)

`server.go`

```go
package main

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
)

func main() {
	// Setup
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Six sick bricks tick")
	})
	e.Server.Addr = ":1323"

	// Serve it like a boss
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
```

### 使用 [graceful](https://github.com/tylerb/graceful)

`server.go`

```go
package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/tylerb/graceful"
)

func main() {
	// Setup
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sue sews rose on slow joe crows nose")
	})
	e.Server.Addr = ":1323"

	// Serve it like a boss
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
```