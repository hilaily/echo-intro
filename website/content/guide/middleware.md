---
title: Middleware
menu:
  side:
    parent: guide
    weight: 5
---

**Middleware** is a function chained in the HTTP request-response cycle with access
to `Echo#Context` which it uses to perform a specific action, for example, logging
every request or limiting the number of requests.

Handler is processed in the end after all middleware are finished executing.

### Middleware Levels

#### Root Level (Before router)

`Echo#Pre()` can be used to register a middleware which is executed before router
processes the request. It is helpful to make any changes to the request properties,
for example, adding or removing a trailing slash from the path so it matches the
route.

The following built-in middleware should be registered at this level:

- AddTrailingSlash
- RemoveTrailingSlash

*Note*: As router has not processed the request, middleware at this level won't
have access to any path related API from `echo.Context`.

#### Root Level (After router)

Most of the time you will register a middleware at this level using `Echo#Use()`.
This middleware is executed after router processes the request and has full access
to `echo.Context` API.

The following built-in middleware should be registered at this level:

- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- CORS
- Static

#### Group Level

When creating a new group, you can register middleware just for that group. For
example, you can have an admin group which is secured by registering a BasicAuth
middleware for it.

*Usage*

```go
e := echo.New()
admin := e.Group("/admin", middleware.BasicAuth())
```

You can also add a middleware after creating a group via `admin.Use()`.

#### Route Level

When defining a new route, you can optionally register middleware just for it.

*Usage*

```go
e := echo.New()
e.GET("/", <Handler>, <Middleware...>)
```

### Logger Middleware

Logger middleware logs the information about each HTTP request.

#### Configuration

```go
LoggerConfig struct {
  // Format is the log format which can be constructed using the following tags:
  //
  // - time_rfc3339
  // - remote_ip
  // - uri
  // - method
  // - path
  // - status
  // - response_time
  // - response_size
  //
  // Example "${remote_id} ${status}"
  //
  // Optional, with default value as `DefaultLoggerConfig.Format`.
  Format string

  // Output is the writer where logs are written.
  // Optional, with default value as os.Stdout.
  Output io.Writer
}
```

#### Default Configuration

```go
DefaultLoggerConfig = LoggerConfig{
  Format: "time=${time_rfc3339}, remote_ip=${remote_ip}, method=${method}, " +
    "uri=${uri}, status=${status}, took=${response_time}, sent=${response_size} bytes\n",
  color:  color.New(),
  Output: os.Stdout,
}
```

*Usage*

`e.Use(middleware.Logger())`

*Sample Output*

`time=2016-03-22T10:33:59-07:00, remote_ip=::1, method=GET, uri=/hello, status=200, took=54.957µs, sent=20 bytes`

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "method=${method}, uri=${uri}, status=${status}\n",
}))
```

Example above uses a `Format` which logs request method and request URI.

*Sample Output*

`method=GET, uri=/hello, status=200`

### Recover Middleware

Recover middleware recovers from panics anywhere in the chain, prints stack trace
and handles the control to the centralized
[HTTPErrorHandler]({{< ref "guide/customization.md#http-error-handler">}}).

#### Configuration

```go
RecoverConfig struct {
  // StackSize is the stack size to be printed.
  // Optional, with default value as 4 KB.
  StackSize int

  // DisableStackAll disables formatting stack traces of all other goroutines
  // into buffer after the trace for the current goroutine.
  // Optional, with default value as false.
  DisableStackAll bool

  // DisablePrintStack disables printing stack trace.
  // Optional, with default value as false.
  DisablePrintStack bool
}
```

#### Default Configuration

```go
DefaultRecoverConfig = RecoverConfig{
	StackSize:  4 << 10, // 4 KB
	StackAll:   true,
	PrintStack: true,
}
```

*Usage*

`e.Use(middleware.Recover())`

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
  StackSize:  1 << 10, // 1 KB
}))
```

Example above uses a `StackSize` of 1 KB and default values for `DisableStackAll`
and `DisablePrintStack`.

### Gzip Middleware

Gzip middleware compresses HTTP response using gzip compression scheme.

#### Configuration

```go
GzipConfig struct {
  // Level is the gzip level.
  // Optional, with default value as -1.
  Level int
}
```

#### Default Configuration

```go
DefaultGzipConfig = GzipConfig{
  Level: -1,
}
```

*Usage*

`e.Use(middleware.Gzip())`

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
  Level: 5
}))
```

Example above uses a `Level` 5 for gzip compression.

### BasicAuth Middleware

BasicAuth middleware provides an HTTP basic authentication.

- For valid credentials it calls the next handler.
- For invalid credentials, it sends "401 - Unauthorized" response.
- For empty or invalid `Authorization` header, it sends "400 - Bad Request" response.

#### Configuration

```go
BasicAuthConfig struct {
  // Validator is the function to validate basic auth credentials.
  Validator BasicAuthValidator
}
```

*Usage*

```go
e := echo.New()
e.Use(middleware.BasicAuth(func(username, password string) bool {
	if username == "joe" && password == "secret" {
		return true
	}
	return false
}))
```

### JWTAuth Middleware

JWTAuth provides a JSON Web Token (JWT) authentication middleware.

- For valid token, it sets the user in context and calls next handler.
- For invalid token, it sends "401 - Unauthorized" response.
- For empty or invalid `Authorization` header, it sends "400 - Bad Request".

#### Configuration

```go
JWTAuthConfig struct {
	// SigningKey is the key to validate token.
	// Required.
	SigningKey []byte

	// SigningMethod is used to check token signing method.
	// Optional, with default value as `HS256`.
	SigningMethod string

	// ContextKey is the key to be used for storing user information from the
	// token into context.
	// Optional, with default value as `user`.
	ContextKey string

	// Extractor is a function that extracts token from the request
	// Optional, with default values as `JWTFromHeader`.
	Extractor JWTExtractor
}
```

#### Default Configuration

```go
DefaultJWTAuthConfig = JWTAuthConfig{
	SigningMethod: AlgorithmHS256,
	ContextKey:    "user",
	Extractor:     JWTFromHeader,
}
```

*Usage*

`e.Use(middleware.JWTAuth([]byte("secret"))`

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.JWTAuthWithConfig(middleware.JWTAuthConfig{
  SigningKey: []byte("secret"),
  Extractor: JWTFromQuery,
}))
```

#### [Recipe]({{< ref "recipes/jwt-authentication.md">}})

### CORS Middleware

CORS middleware implements [CORS](http://www.w3.org/TR/cors/) specification.
CORS gives web servers cross-domain access controls, which enable secure cross-domain
data transfers.

#### Configuration

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

#### Default Configuration

```go
DefaultCORSConfig = CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
}
```

*Usage*

`e.Use(middleware.CORS())`

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
}))
```

### [Static Middleware]({{< ref "guide/static-files.md#using-static-middleware">}})

### AddTrailingSlash Middleware

AddTrailingSlash adds a trailing slash to the request URI.

*Usage*

```go
e := echo.New()
e.Pre(middleware.AddTrailingSlash())
```

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.AddTrailingSlashWithConfig(TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```

This will add a trailing slash to the request URI and redirect with `StatusMovedPermanently`.

### RemoveTrailingSlash Middleware

RemoveTrailingSlash removes a trailing slash from the request URI.

*Usage*

```go
e := echo.New()
e.Pre(middleware.RemoveTrailingSlash())
```

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.RemoveTrailingSlashWithConfig(TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```

This will remove a trailing slash from the request URI and redirect with `StatusMovedPermanently`.

### Method Override Middleware

Method override middleware checks for the overridden method from the request and
uses it instead of the original method.

For security reasons, only `POST` method can be overridden.

#### Configuration

```go
MethodOverrideConfig struct {
  // Getter is a function that gets overridden method from the request.
  Getter MethodOverrideGetter
}
```

#### Default Configuration

```go
DefaultMethodOverrideConfig = MethodOverrideConfig{
  Getter: MethodFromHeader(echo.HeaderXHTTPMethodOverride),
}
```

*Usage*

`e.Pre(middleware.MethodOverride())`

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
  Getter: middleware.MethodFromForm("_method"),
}))
```

### [Writing Custom Middleware]({{< ref "recipes/middleware.md">}})
