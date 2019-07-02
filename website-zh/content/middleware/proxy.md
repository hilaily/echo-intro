+++
title = "代理"
url = "/middleware/proxy"
[menu.side]
  name = "代理"
  parent = "middleware"
  weight = 13

+++

## Proxy (代理) 中间件

Proxy 提供 HTTP / WebSocket 反向代理中间件。它使用已配置的负载平衡技术将请求转发到上游服务器。

*用法*

```go
url1, err := url.Parse("http://localhost:8081")
if err != nil {
  e.Logger.Fatal(err)
}
url2, err := url.Parse("http://localhost:8082")
if err != nil {
  e.Logger.Fatal(err)
}
e.Use(middleware.Proxy(&middleware.RoundRobinBalancer{
  Targets: []*middleware.ProxyTarget{
    {
      URL: url1,
    },
    {
      URL: url2,
    },
  },
}))
```

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{}))
```

### 配置

```go
// ProxyConfig defines the config for Proxy middleware.
ProxyConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Balancer defines a load balancing technique.
  // Required.
  // Possible values:
  // - RandomBalancer
  // - RoundRobinBalancer
  Balancer ProxyBalancer
}
```

*默认配置*

| 名称    | 值             |
| ------- | -------------- |
| Skipper | DefaultSkipper |

### [示例](https://github.com/labstack/echox/blob/master/cookbook/reverse-proxy)