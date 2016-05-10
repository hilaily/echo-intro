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
  // Gzip compression level.
  // Optional. Default value -1.
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
