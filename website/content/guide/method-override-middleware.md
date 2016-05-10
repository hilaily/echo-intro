+++
title = "Method Override Middleware"
[menu.side]
  name = "MethodOverride"
  parent = "middleware"
  weight = 5
+++

## MethodOverride Middleware

MethodOverride middleware checks for the overridden method from the request and
uses it instead of the original method.

For security reasons, only `POST` method can be overridden.

### Configuration

```go
MethodOverrideConfig struct {
  // Getter is a function that gets overridden method from the request.
  // Optional. Default values MethodFromHeader(echo.HeaderXHTTPMethodOverride).
  Getter MethodOverrideGetter
}
```

### Default Configuration

```go
DefaultMethodOverrideConfig = MethodOverrideConfig{
  Getter: MethodFromHeader(echo.HeaderXHTTPMethodOverride),
}
```

*Usage*

`e.Pre(middleware.MethodOverride())`

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
  Getter: middleware.MethodFromForm("_method"),
}))
```
