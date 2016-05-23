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
	Secret []byte

	// Context key to store generated CSRF token into context.
	// Optional. Default value "csrf".
	ContextKey string

	// Extractor is a function that extracts token from the request.
	// Optional. Default value CSRFTokenFromHeader(echo.HeaderXCSRFToken).
	Extractor CSRFTokenExtractor

	// Name of the CSRF cookie. This cookie will store CSRF token.
	// Optional. Default value "csrf".
	CookieName string

	// Domain of the CSRF cookie.
	// Optional. Default value none.
	CookieDomain string

	// Path of the CSRF cookie.
	// Optional. Default value none.
	CookiePath string

	// Expiration time of the CSRF cookie.
	// Optional. Default value 24H.
	CookieExpires time.Time

	// Indicates if CSRF cookie is secure.
	CookieSecure bool
	// Optional. Default value false.

	// Indicates if CSRF cookie is HTTP only.
	// Optional. Default value false.
	CookieHTTPOnly bool
}
```

### Default Configuration

```go
DefaultCSRFConfig = CSRFConfig{
	ContextKey:    "csrf",
	Extractor:     CSRFTokenFromHeader(echo.HeaderXCSRFToken),
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
  Extractor: CSRFTokenFromHeader("X-XSRF-TOKEN"),
}))
```

Example above uses `X-XSRF-TOKEN` request header to extract CSRF token.
