+++
title = "Key 验证"
url = "/middleware/key-auth"
[menu.side]
  name = "Key Auth"
  parent = "middleware"
  weight = 5
+++

## Key Auth 中间件

Key Auth 中间件提供了一个基于 key 的验证方式。

- 对于有效的 key，它将调用下一个处理程序。
- 对于无效的 key，它发送"401 - Unauthorized"的响应。
- 对于空的 key，它发送"400 - Bad Request"。

*使用*

```go
e.Use(middleware.KeyAuth(func(key string) bool {
  return key == "valid-key"
}))
```

### 自定义配置

*使用*

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

### 默认配置

```go
DefaultKeyAuthConfig = KeyAuthConfig{
  Skipper:    defaultSkipper,
  KeyLookup:  "header:" + echo.HeaderAuthorization,
  AuthScheme: "Bearer",
}
```

