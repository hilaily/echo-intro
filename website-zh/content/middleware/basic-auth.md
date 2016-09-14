+++
title = "Basic Auth"
url="basic-auth"
[menu.side]
  name = "BasicAuth(基本认证)"
  parent = "middleware"
  weight = 3
+++

## BasicAuti(基本认证) 中间件

BasicAuth 中间件提供了 HTTP 的基本认证方式。

- 对于有效的请求则继续执行后面的处理。
- 对于无效的请求，返回"401 - Unauthorized"响应。
- 对于请求头中"Authorization"无效或者为空的，返回"400 - Bad Request" 响应。

*用法*

```go
e := echo.New()
e.Use(middleware.BasicAuth(func(username, password string) bool {
	if username == "joe" && password == "secret" {
		return true
	}
	return false
}))
```
