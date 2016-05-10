+++
title = "JWT Middleware"
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
  SigningKey []byte

  // Signing method, used to check token signing method.
  // Optional. Default value HS256.
  SigningMethod string

  // Context key to store user information from the token into context.
  // Optional. Default value "user".
  ContextKey string

  // Extractor is a function that extracts token from the request.
  // Optional. Default value JWTFromHeader.
  Extractor JWTExtractor
}
```

### Default Configuration

```go
DefaultJWTConfig = JWTConfig{
	SigningMethod: AlgorithmHS256,
	ContextKey:    "user",
	Extractor:     JWTFromHeader,
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
  Extractor: JWTFromQuery,
}))
```

### [Recipe]({{< ref "recipes/jwt.md">}})
