+++
title = "CSRF Middleware"
[menu.side]
  name = "CSRF"
  parent = "middleware"
  weight = 5
+++

## CSRF Middleware

Cross-site request forgery, also known as one-click attack or session riding and
abbreviated as CSRF (sometimes pronounced sea-surf) or XSRF, is a type of malicious
exploit of a website where unauthorized commands are transmitted from a user that
the website trusts.

### Configuration

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

### Default Configuration

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

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
  TokenLookup: "header:X-XSRF-TOKEN",
}))
```

Example above uses `X-XSRF-TOKEN` request header to extract CSRF token.
