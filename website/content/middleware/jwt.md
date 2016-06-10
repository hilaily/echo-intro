+++
title = "JWT"
[menu.side]
  name = "JWT"
  parent = "middleware"
  weight = 5
+++

## JWT Middleware

JWT provides a JSON Web Token (JWT) authentication middleware.

- For valid token, it sets the user in context and calls next handler.
- For invalid token, it sends "401 - Unauthorized" response.
- For empty or invalid `Authorization` header, it sends "400 - Bad Request".

### Configuration

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

### Default Configuration

```go
DefaultJWTConfig = JWTConfig{
  SigningMethod: AlgorithmHS256,
  ContextKey:    "user",
  TokenLookup:   "header:" + echo.HeaderAuthorization,
}
```

*Usage*

`e.Use(middleware.JWT([]byte("secret"))`

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
  SigningKey: []byte("secret"),
  TokenLookup: "query:token",
}))
```

### [Recipe]({{< ref "recipes/jwt.md">}})
