+++
title = "基本认证"
url="/middleware/basic-auth"
[menu.side]
  name = "基本认证"
  parent = "middleware"
  weight = 2

+++

## Basic Auth (基本认证) 中间件

Basic Auth 中间件提供了 HTTP 的基本认证方式。

- 对于有效的请求，则继续调用下一个处理程序 (handler) 。
- 对于丢失或无效的请求，则返回 "401 - Unauthorized" 响应。

*用法*

```go
e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	if username == "joe" && password == "secret" {
		return true, nil
	}
	return false, nil
}))
```

### 自定义配置

*用法*

```go
e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))
```

### 配置

```go
BasicAuthConfig struct {
  // Skipper 定义了一个跳过中间件的函数
  Skipper Skipper

  // Validator 是一个用来验证 BasicAuth 是否合法的函数
  // Validator 是必须的.
  Validator BasicAuthValidator

  // Realm 是一个用来定义 BasicAuth 的 Realm 属性的字符串
  // 默认值是 "Restricted"
  Realm string
}
```

*默认配置*

```go
DefaultBasicAuthConfig = BasicAuthConfig{
	Skipper: defaultSkipper,
}
```

