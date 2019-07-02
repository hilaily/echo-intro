+++
title = "密钥认证"
url = "/middleware/key-auth"
[menu.side]
  name = "密钥认证"
  parent = "middleware"
  weight = 10

+++

## Key Auth (密钥认证) 中间件

Key Auth 中间件提供了一个基于密钥的验证方式。

- 对于有效密钥，它将调用下一个处理程序。
- 对于无效密钥，它会发送 “401 - Unauthorized” 响应。
- 对于丢失密钥，它发送 “400 - Bad Request” 响应。

*用法*

```go
e.Use(middleware.KeyAuth(func(key string) bool {
  return key == "valid-key"
}))
```

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
  KeyLookup: "query:api-key",
}))
```

### 配置

```go
// KeyAuthConfig defines the config for KeyAuth middleware.
KeyAuthConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // KeyLookup is a string in the form of "<source>:<name>" that is used
  // to extract key from the request.
  // Optional. Default value "header:Authorization".
  // Possible values:
  // - "header:<name>"
  // - "query:<name>"
  KeyLookup string `json:"key_lookup"`

  // AuthScheme to be used in the Authorization header.
  // Optional. Default value "Bearer".
  AuthScheme string

  // Validator is a function to validate key.
  // Required.
  Validator KeyAuthValidator
}
```

*默认配置*

```go
DefaultKeyAuthConfig = KeyAuthConfig{
  Skipper:    defaultSkipper,
  KeyLookup:  "header:" + echo.HeaderAuthorization,
  AuthScheme: "Bearer",
}
```

