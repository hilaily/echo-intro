+++
title = "JWT Auth Middleware"
[menu.side]
  name = "JWTAuth"
  parent = "middleware"
  weight = 5
+++

## JWTAuth Middleware

JWTAuth provides a JSON Web Token (JWT) authentication middleware.

- For valid token, it sets the user in context and calls next handler.
- For invalid token, it sends "401 - Unauthorized" response.
- For empty or invalid `Authorization` header, it sends "400 - Bad Request".

### Configuration

```go
JWTAuthConfig struct {
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
DefaultJWTAuthConfig = JWTAuthConfig{
	SigningMethod: AlgorithmHS256,
	ContextKey:    "user",
	Extractor:     JWTFromHeader,
}
```

*Usage*

`e.Use(middleware.JWTAuth([]byte("secret"))`

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.JWTAuthWithConfig(middleware.JWTAuthConfig{
  SigningKey: []byte("secret"),
  Extractor: JWTFromQuery,
}))
```

### [Recipe]({{< ref "recipes/jwt-authentication.md">}})
