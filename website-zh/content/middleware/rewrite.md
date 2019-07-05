+++
title = "重写"
url = "/middleware/rewrite"
[menu.side]
  name = "重写"
  parent = "middleware"
  weight = 17

+++

## Rewrite (重写) 中间件

Rewrite 中间件会根据提供的规则重写URL路径，它更适用于向后兼容与创建更清晰、更具描述性的链接。

*用法*

```go
e.Pre(middleware.Rewrite(map[string]string{
  "/old":              "/new",
  "/api/*":            "/$1",
  "/js/*":             "/public/javascripts/$1",
  "/users/*/orders/*": "/user/$1/order/$2",
}))
```

星号中捕获的值可以通过索引检索，例如 $1, $2 等等。

### Custom Configuration

*用法*

```go
e := echo.New()
e.Pre(middleware.RewriteWithConfig(middleware.RewriteConfig{}))
```

### 配置

```go
// RewriteConfig defines the config for Rewrite middleware.
RewriteConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Rules defines the URL path rewrite rules.
  Rules map[string]string `yaml:"rules"`
}
```

*默认配置*

| 名称    | 值             |
| :------ | :------------- |
| Skipper | DefaultSkipper |

> 重写中间件应该在路由之前通过 Echo#Pre() 注册从而触发。

