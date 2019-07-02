+++
title = "重定向"
url = "/middleware/redirect"
[menu.side]
  name = "重定向"
  parent = "middleware"
  weight = 15

+++

## Redirect (重定向) 中间件

#### HTTPS 重定向

HTTPS 重定向中间件将 http 请求重定向到 https。例如，http://laily.net 将被重定向到 https://laily.net。

*用法*

```go
e := echo.New()
e.Pre(middleware.HTTPSRedirect())
```

#### HTTPS WWW 重定向

HTTPS WWW 重定向将 http 请求重定向到带 www 的 https 请求。例如，http://laily.net 将被重定向到 https://www.laily.net。

*用法*

```go
e := echo.New()
e.Pre(middleware.HTTPSWWWRedirect())
```

#### HTTPS NonWWW 重定向

HTTPS NonWWW 将 http 请求重定向到不带 www 的 https 请求。例如，http://www.laily.net 将被重定向到 https://laily.net。

*用法*

```go
e := echo.New()
e.Pre(middleware.HTTPSNonWWWRedirect())
```

### WWW 重定向

WWW 重定向将不带 www 的请求重定向到带 www 的请求。

例如，http://laily.net 重定向到 http://www.laily.net

*用法*

```go
e := echo.New()
e.Pre(middleware.WWWRedirect())
```

### NonWWW 重定向

NonWWW 重定向将带 www 的请求重定向到不带 www 的请求。

例如，http://www.laily.net 重定向到 http://laily.net

*用法*

```go
e := echo.New()
e.Pre(middleware.NonWWWRedirect())
```

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
  Code: http.StatusTemporaryRedirect,
}))
```

上面的示例将 HTTP 的请求重定向到 HTTPS，使用 `307 - StatusTemporaryRedirect`  状态码跳转。

### 配置

```go
RedirectConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Status code to be used when redirecting the request.
  // Optional. Default value http.StatusMovedPermanently.
  Code int `json:"code"`
}
```

*默认配置*

```go
DefaultRedirectConfig = RedirectConfig{
  Skipper: defaultSkipper,
  Code:    http.StatusMovedPermanently,
}
```

