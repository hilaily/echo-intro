+++
title = "CSRF(跨站请求伪造)"
url = "/middleware/ecsrf"
[menu.side]
  name = "CSRF(跨站请求伪造)"
  parent = "middleware"
  weight = 5
+++

## CSRF Middleware(跨站请求伪造)

CSRF（Cross-site request forgery跨站请求伪造，也被称为“One Click Attack”或者Session Riding，通常缩写为CSRF或者XSRF，是一种挟制用户在当前已登录的Web应用程序上执行非本意的操作的攻击方法。 跟跨网站脚本（XSS）相比，XSS 利用的是用户对指定网站的信任，CSRF 利用的是网站对用户网页浏览器的信任。

### 配置

```go
CSRFConfig struct {
  // Key to create CSRF token.
  Secret []byte `json:"secret"`

  // TokenLookup is a string in the form of "<source>:<key>" that is used
  // to extract token from the request.
  // Optional. Default value "header:X-CSRF-Token".
  // Possible values:
  // - "header:<name>"
  // - "form:<name>"
  // - "header:<name>"
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

  // Expiration time of the CSRF cookie.
  // Optional. Default value 24H.
  CookieExpires time.Time `json:"cookie_expires"`

  // Indicates if CSRF cookie is secure.
  CookieSecure bool `json:"cookie_secure"`
  // Optional. Default value false.

  // Indicates if CSRF cookie is HTTP only.
  // Optional. Default value false.
  CookieHTTPOnly bool `json:"cookie_http_only"`
}
```

### 默认配置

```go
DefaultCSRFConfig = CSRFConfig{
  TokenLookup:   "header:" + echo.HeaderXCSRFToken,
  ContextKey:    "csrf",
  CookieName:    "csrf",
  CookieExpires: time.Now().Add(24 * time.Hour),
}
```

*Usage*

`e.Use(middleware.CSRF("secret"))`

### 自定义配置

*Usage*

```go
e := echo.New()
e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
  TokenLookup: "header:X-XSRF-TOKEN",
}))
```

Example above uses `X-XSRF-TOKEN` request header to extract CSRF token.
