+++
title = "Add Trailing Slash Middleware"
[menu.side]
  name = "AddTrailingSlash"
  parent = "middleware"
  weight = 5
+++

## AddTrailingSlash Middleware

AddTrailingSlash middleware adds a trailing slash to the request URI.

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
