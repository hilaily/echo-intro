+++
title = "CSRF "
url = "/middleware/csrf"
[menu.side]
  name = "CSRF "
  parent = "middleware"
  weight = 6

+++

## CSRF (跨域请求伪造) 中间件

CSRF (Cross-site request forgery) 跨域请求伪造，也被称为 **one-click attack** 或者 **session riding**，通常缩写为 **CSRF** 或者 **XSRF**， 是一种挟制用户在当前已登录的Web应用程序上执行非本意的操作的攻击方法。[[1\]](https://zh.wikipedia.org/wiki/跨站请求伪造#cite_note-Ristic-1) 跟[跨网站脚本](https://zh.wikipedia.org/wiki/跨網站指令碼) (XSS) 相比，**XSS** 利用的是用户对指定网站的信任，CSRF 利用的是网站对用户网页浏览器的信任。

*用法*

`e.Use(middleware.CSRF())`

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
  TokenLookup: "header:X-XSRF-TOKEN",
}))
```

上面的例子使用 `X-XSRF-TOKEN` 请求头取出 CSRF 的 token 值。

### 获取 CSRF Token

#### 服务器端

服务器端可以使用 `ContextKey ` 从 `Echo#Context` 拿到 CSRF token 然后通过模版传给客户端。

#### 客户端

客户端可以通过 CSRF cookie 拿到 token 值。

### 配置

```go
// CSRFConfig defines the config for CSRF middleware.
CSRFConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // TokenLength is the length of the generated token.
  TokenLength uint8 `json:"token_length"`
  // Optional. Default value 32.

  // TokenLookup is a string in the form of "<source>:<key>" that is used
  // to extract token from the request.
  // Optional. Default value "header:X-CSRF-Token".
  // Possible values:
  // - "header:<name>"
  // - "form:<name>"
  // - "query:<name>"
  TokenLookup string `json:"token_lookup"`

  // Context key to store generated CSRF token into context.
  // Optional. Default value "csrf".
  ContextKey string `json:"context_key"`

  // Name of the CSRF cookie. This cookie will store CSRF token.
  // Optional. Default value "csrf".
  CookieName string `json:"cookie_name"`

  // Domain of the CSRF cookie.
  // Optional. Default value none.
  CookieDomain string `json:"cookie_domain"`

  // Path of the CSRF cookie.
  // Optional. Default value none.
  CookiePath string `json:"cookie_path"`

  // Max age (in seconds) of the CSRF cookie.
  // Optional. Default value 86400 (24hr).
  CookieMaxAge int `json:"cookie_max_age"`

  // Indicates if CSRF cookie is secure.
  // Optional. Default value false.
  CookieSecure bool `json:"cookie_secure"`

  // Indicates if CSRF cookie is HTTP only.
  // Optional. Default value false.
  CookieHTTPOnly bool `json:"cookie_http_only"`
}
```

*默认配置*

```go
DefaultCSRFConfig = CSRFConfig{
  Skipper:      defaultSkipper,
  TokenLength:  32,
  TokenLookup:  "header:" + echo.HeaderXCSRFToken,
  ContextKey:   "csrf",
  CookieName:   "_csrf",
  CookieMaxAge: 86400,
}
```

