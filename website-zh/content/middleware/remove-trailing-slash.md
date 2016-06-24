+++
title = "Remove Trailing Slash"
[menu.side]
  name = "RemoveTrailingSlash"
  parent = "middleware"
  weight = 5
+++

## RemoveTrailingSlash Middleware

RemoveTrailingSlash middleware removes a trailing slash from the request URI.

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
e.Pre(middleware.RemoveTrailingSlash())
```

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.RemoveTrailingSlashWithConfig(TrailingSlashConfig{
  RedirectCode: http.StatusMovedPermanently,
}))
```

This will remove a trailing slash from the request URI and redirect with `StatusMovedPermanently`.
