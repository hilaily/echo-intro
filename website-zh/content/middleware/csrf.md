+++
title = "CSRF 跨站请求伪造"
url = "/middleware/ecsrf"
[menu.side]
  name = "CSRF 跨站请求伪造"
  parent = "middleware"
  weight = 5
+++

## CSRF Middleware(跨站请求伪造)

CSRF（Cross-site request forgery 跨站请求伪造，也被称为“One Click Attack”或者Session Riding，通常缩写为CSRF或者XSRF，是一种挟制用户在当前已登录的Web应用程序上执行非本意的操作的攻击方法。 跟跨网站脚本（XSS）相比，XSS 利用的是用户对指定网站的信任，CSRF 利用的是网站对用户网页浏览器的信任。

*使用*

```go
e.Use(middleware.CSRF())
```

### 自定义配置

*使用*

```go
e := echo.New()
e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
  TokenLookup: "header:X-XSRF-TOKEN",
}))
```

上面例子将使用`X-XSRF-TOKEN` 请求头取出 CSRF 的 token 值。

### 获取 CSRF Token

#### 服务器端

服务器端可以使用 `ContextKey `从 `Echo#Context` 拿到 CSRF token 然后通过模版传给客户端。

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

### 默认配置

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

