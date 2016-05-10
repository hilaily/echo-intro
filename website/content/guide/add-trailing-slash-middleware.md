+++
title = "Add Trailing Slash Middleware"
[menu.side]
  name = "AddTrailingSlash"
  parent = "middleware"
  weight = 5
+++

## AddTrailingSlash Middleware

AddTrailingSlash middleware adds a trailing slash to the request URI.

### Configuration

```go
TrailingSlashConfig struct {
  // Status code to be used when redirecting the request.
  // Optional, but when provided the request is redirected using this code.
  RedirectCode int
}
```

*Usage*

```go
e := echo.New()
e.Pre(middleware.AddTrailingSlash())
```

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.AddTrailingSlashWithConfig(TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```

This will add a trailing slash to the request URI and redirect with `StatusMovedPermanently`.
