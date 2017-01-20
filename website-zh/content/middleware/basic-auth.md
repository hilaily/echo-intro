+++
title = "BasicAuth 基本认证"
url="/middleware/basic-auth"
[menu.side]
  name = "BasicAuth 基本认证"
  parent = "middleware"
  weight = 3
+++

## BasicAuth (基本认证) 中间件

BasicAuth 中间件提供了 HTTP 的基本认证方式。

- 对于有效的请求则继续执行后面的处理。
- 对于无效的请求，返回"401 - Unauthorized"响应。

*用法*

```go
e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) bool {
	if username == "joe" && password == "secret" {
		return true
	}
	return false
}))
```

## 自定义配置

*用法*

```go
e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}}))
```

*配置*

```go
BasicAuthConfig struct {
  // Skipper 定义了一个跳过中间间的函数
  Skipper Skipper

  // Validator 是一个用来验证 BasicAuth 是否合法的函数
  // Validator 是必须的.
  Validator BasicAuthValidator
}
```

*默认配置*

```go
DefaultBasicAuthConfig = BasicAuthConfig{
	Skipper: defaultSkipper,
}
```











