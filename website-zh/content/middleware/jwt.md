+++
title = "JWT"
url = "/middleware/jwt"
[menu.side]
  name = "JWT"
  parent = "middleware"
  weight = 9
+++

## JWT 中间件

JWT 提供了一个 JSON Web Token (JWT) 认证中间件。

- 对于有效的 token，它将用户置于上下文中并调用下一个处理程序。
- 对于无效的 token，它会发送 "401 - Unauthorized" 响应。
- 对于丢失或无效的 `Authorization` 标头，它会发送 "400 - Bad Request" 。

*用法*

`e.Use(middleware.JWT([]byte("secret"))`

### 自定义配置

*用法*

```go
e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
  SigningKey: []byte("secret"),
  TokenLookup: "query:token",
}))
```

### 配置

```go
// JWTConfig defines the config for JWT middleware.
JWTConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Signing key to validate token.
  // Required.
  SigningKey interface{}

  // Signing method, used to check token signing method.
  // Optional. Default value HS256.
  SigningMethod string

  // Context key to store user information from the token into context.
  // Optional. Default value "user".
  ContextKey string

  // Claims are extendable claims data defining token content.
  // Optional. Default value jwt.MapClaims
  Claims jwt.Claims

  // TokenLookup is a string in the form of "<source>:<name>" that is used
  // to extract token from the request.
  // Optional. Default value "header:Authorization".
  // Possible values:
  // - "header:<name>"
  // - "query:<name>"
  // - "cookie:<name>"
  TokenLookup string

  // AuthScheme to be used in the Authorization header.
  // Optional. Default value "Bearer".
  AuthScheme string
}
```

*默认配置*

```go
DefaultJWTConfig = JWTConfig{
  Skipper:       defaultSkipper,
  SigningMethod: AlgorithmHS256,
  ContextKey:    "user",
  TokenLookup:   "header:" + echo.HeaderAuthorization,
  AuthScheme:    "Bearer",
  Claims:        jwt.MapClaims{},
}
```

### [示例](https://echo.labstack.com/cookbook/jwt)