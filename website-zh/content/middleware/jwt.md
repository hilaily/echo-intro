+++
title = "JWT"
url = "jwt"
[menu.side]
  name = "JWT"
  parent = "middleware"
  weight = 5
+++

## JWT 中间件

JSON Web Token(JWT)是一个非常轻巧的规范。这个规范允许我们使用JWT在用户和服务器之间传递安全可靠的信息。

- 对于有效的令牌，它将用户设置为上下文，并调用下一个处理程序。
- 对于无效的令牌，它发送"401 - Unauthorized"的响应。
- 对于空的或无效的`Authorization`头，它发送"400 - Bad Request"。

### 配置

```go
JWTConfig struct {
  // Signing key to validate token.
  // Required.
  SigningKey []byte `json:"signing_key"`

  // Signing method, used to check token signing method.
  // Optional. Default value HS256.
  SigningMethod string `json:"signing_method"`

  // Context key to store user information from the token into context.
  // Optional. Default value "user".
  ContextKey string `json:"context_key"`

  // TokenLookup is a string in the form of "<source>:<name>" that is used
  // to extract token from the request.
  // Optional. Default value "header:Authorization".
  // Possible values:
  // - "header:<name>"
  // - "form:<name>"
  TokenLookup string `json:"token_lookup"`
}
```

### 默认配置

```go
DefaultJWTConfig = JWTConfig{
  SigningMethod: AlgorithmHS256,
  ContextKey:    "user",
  TokenLookup:   "header:" + echo.HeaderAuthorization,
}
```

*Usage*

`e.Use(middleware.JWT([]byte("secret"))`

### 自定义配置
*Usage*

```go
e := echo.New()
e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
  SigningKey: []byte("secret"),
  TokenLookup: "query:token",
}))
```

### [Recipe]({{< ref "recipes/jwt.md">}})
