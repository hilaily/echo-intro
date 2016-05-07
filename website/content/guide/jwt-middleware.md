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
	// SigningKey is the key to validate token.
	// Required.
	SigningKey []byte

	// SigningMethod is used to check token signing method.
	// Optional, with default value as `HS256`.
	SigningMethod string

	// ContextKey is the key to be used for storing user information from the
	// token into context.
	// Optional, with default value as `user`.
	ContextKey string

	// Extractor is a function that extracts token from the request
	// Optional, with default values as `JWTFromHeader`.
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
