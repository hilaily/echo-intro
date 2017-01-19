+++
title = "JWT"
url = "/middleware/jwt"
[menu.side]
  name = "JWT"
  parent = "middleware"
  weight = 5
+++

## JWT 中间件

JSON Web Token(JWT) 是一个非常轻巧的规范。这个规范允许我们使用JWT在用户和服务器之间传递安全可靠的信息。

- 对于有效的令牌，它将用户存储进上下文，并调用下一个处理程序。
- 对于无效的令牌，它发送"401 - Unauthorized"的响应。
- 对于空的或无效的`Authorization`头，它发送"400 - Bad Request"。

*使用*

```go
e.Use(middleware.JWT([]byte("secret"))
```

### 自定义配置

*使用*

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

### 默认配置

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
