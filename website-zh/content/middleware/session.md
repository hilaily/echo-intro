+++
title = "会话"
url = "/middleware/session"
[menu.side]
  name = "会话"
  parent = "middleware"
  weight = 19

+++

## Session (会话) 中间件

Session 中间件促进了 [gorilla/sessions](https://github.com/gorilla/sessions) 支持的 HTTP 会话管理。默认提供了基于 cookie 与文件系统的会话存储；然而，你也可以访问 [community maintained implementation](https://github.com/gorilla/sessions#store-implementations) 来参考其各式各样的后端实现。

> Echo 社区贡献

### 依赖

```go
import (
  "github.com/gorilla/sessions"
  "github.com/labstack/echo-contrib/session"
)
```

*用法*

```go
e := echo.New()
e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

e.GET("/", func(c echo.Context) error {
  sess, _ := session.Get("session", c)
  sess.Options = &sessions.Options{
    Path:     "/",
    MaxAge:   86400 * 7,
    HttpOnly: true,
  }
  sess.Values["foo"] = "bar"
  sess.Save(c.Request(), c.Response())
  return c.NoContent(http.StatusOK)
})
```

### Custom Configuration

*用法*

```go
e := echo.New()
e.Use(session.MiddlewareWithConfig(session.Config{}))
```

### 配置

```go
Config struct {
  // Skipper defines a function to skip middleware.
  Skipper middleware.Skipper

  // Session store.
  // Required.
  Store sessions.Store
}
```

*默认配置*

```go
DefaultConfig = Config{
  Skipper: DefaultSkipper,
}
```

