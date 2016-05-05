+++
title = "Remove Trailing Slash Middleware"
[menu.side]
  name = "RemoveTrailingSlash"
  parent = "middleware"
  weight = 5
+++

## RemoveTrailingSlash Middleware

RemoveTrailingSlash middleware removes a trailing slash from the request URI.

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
