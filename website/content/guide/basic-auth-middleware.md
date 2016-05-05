+++
title = "Basic Auth Middleware"
[menu.side]
  name = "BasicAuth"
  parent = "middleware"
  weight = 5
+++

## BasicAuth Middleware

BasicAuth middleware provides an HTTP basic authentication.

- For valid credentials it calls the next handler.
- For invalid credentials, it sends "401 - Unauthorized" response.
- For empty or invalid `Authorization` header, it sends "400 - Bad Request" response.

*Usage*

```go
e := echo.New()
e.Use(middleware.BasicAuth(func(username, password string) bool {
	if username == "joe" && password == "secret" {
		return true
	}
	return false
}))
```
