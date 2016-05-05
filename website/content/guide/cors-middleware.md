+++
title = "CORS Middleware"
[menu.side]
  name = "CORS"
  parent = "middleware"
  weight = 5
+++

## CORS Middleware

CORS middleware implements [CORS](http://www.w3.org/TR/cors) specification.
CORS gives web servers cross-domain access controls, which enable secure cross-domain
data transfers.

### Configuration

```go
CORSConfig struct {
	// AllowOrigin defines a list of origins that may access the resource.
	// Optional, with default value as []string{"*"}.
	AllowOrigins []string

	// AllowMethods defines a list methods allowed when accessing the resource.
	// This is used in response to a preflight request.
	// Optional, with default value as `DefaultCORSConfig.AllowMethods`.
	AllowMethods []string

	// AllowHeaders defines a list of request headers that can be used when
	// making the actual request. This in response to a preflight request.
	// Optional, with default value as []string{}.
	AllowHeaders []string

	// AllowCredentials indicates whether or not the response to the request
	// can be exposed when the credentials flag is true. When used as part of
	// a response to a preflight request, this indicates whether or not the
	// actual request can be made using credentials.
	// Optional, with default value as false.
	AllowCredentials bool

	// ExposeHeaders defines a whitelist headers that clients are allowed to
	// access.
	// Optional, with default value as []string{}.
	ExposeHeaders []string

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached.
	// Optional, with default value as 0.
	MaxAge int
}
```

### Default Configuration

```go
DefaultCORSConfig = CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
}
```

*Usage*

`e.Use(middleware.CORS())`

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
}))
```
