+++
title = "Method Override"
url = "method-override"
[menu.side]
  name = "MethodOverride"
  parent = "middleware"
  weight = 5
+++

## MethodOverride 中间件

MethodOverride中间件检查从请求中重写的方法，并使用它来代替原来的方法。

出于安全原因，只有`POST`方法可以被重写。

### 配置

```go
MethodOverrideConfig struct {
  // Getter is a function that gets overridden method from the request.
  // Optional. Default values MethodFromHeader(echo.HeaderXHTTPMethodOverride).
  Getter MethodOverrideGetter
}
```

### 默认配置

```go
DefaultMethodOverrideConfig = MethodOverrideConfig{
  Getter: MethodFromHeader(echo.HeaderXHTTPMethodOverride),
}
```

*Usage*

`e.Pre(middleware.MethodOverride())`

### 自定义配置

*Usage*

```go
e := echo.New()
e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
  Getter: middleware.MethodFromForm("_method"),
}))
```
