+++
title = "Gzip Middleware"
[menu.side]
  name = "Gzip"
  parent = "middleware"
  weight = 5
+++

## Gzip Middleware

Gzip middleware compresses HTTP response using gzip compression scheme.

### Configuration

```go
GzipConfig struct {
  // Level is the gzip level.
  // Optional, with default value as -1.
  Level int
}
```

### Default Configuration

```go
DefaultGzipConfig = GzipConfig{
  Level: -1,
}
```

*Usage*

`e.Use(middleware.Gzip())`

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
  Level: 5
}))
```
