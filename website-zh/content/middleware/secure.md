+++
title = "Secure(安全)"
url = "/middleware/secure"
[menu.side]
  name = "Secure(安全)"
  parent = "middleware"
  weight = 5
+++

## Secure 中间件

Secure 中间件用于阻止跨站脚本攻击(XSS)，内容嗅探，点击劫持，不安全链接等其他代码注入攻击。

### 配置

```go
SecureConfig struct {
  // XSSProtection 通过设置`X-XSS-Protection`头
  // 来提供XSS攻击的防护。
  // 可选。默认值 "1; mode=block"。
  XSSProtection string

  // ContentTypeNosniff 通过设置`X-Content-Type-Options`头
  // 来防止内容嗅探。
  // 可选。默认值 "nosniff"。
  ContentTypeNosniff string

  // XFrameOptions 被用来指示是否允许浏览器在<fram>，<iframe>或者<object>中渲染页面。
  // 网站可以通过这样来避免点击劫持，保证网站的内容不会被其他站点嵌入。
  // 可选。默认值 "SAMEORIGIN".
  // 可使用的值:
  // `SAMEORIGIN` - 页面只能在同域名的页面下被渲染。
  // `DENY` - 页面不允许在 frame 中显示。
  // `ALLOW-FROM uri` - 页面只能在指定域名的 frame 中显示。
  XFrameOptions string

  // HSTSMaxAge 设置 `Strict-Transport-Security` 头来指示浏览器需要记住这个网站只能通过HTTPS来访问的时间(单位秒)。
  // 这样可以减少遭受中间人攻击(HITM)的几率。
  // 可选。默认值 0。
  HSTSMaxAge int

  // HSTSExcludeSubdomains 不会在`Strict Transport Security`中设置`includeSubdomains`标签。
  // 即从安全规则中排除所有子域名。
  // header, excluding all subdomains from security policy. It has no effect
  // 只有在HSTSMaxAge 被设置为非0值时该参数才有效。
  // 可选。默认值 false。
  HSTSExcludeSubdomains bool

  // ContentSecurityPolicy 用来设置 `Content-Security-Policy` 头。
  // `Content-Security-Policy` 用来定义页面可以加载哪些资源，减少XSS等通过运行不安全代码的注入攻击。
  // 可选。默认值 "".
  ContentSecurityPolicy string
}
```

### 默认配置

*用法*

```go
DefaultSecureConfig = SecureConfig{
  XSSProtection:      "1; mode=block",
  ContentTypeNosniff: "nosniff",
  XFrameOptions:      "SAMEORIGIN",
}
```

`e.Use(middleware.Secure())`

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	XSSProtection:         "",
	ContentTypeNosniff:    "",
	XFrameOptions:         "",
	HSTSMaxAge:            3600,
	ContentSecurityPolicy: "default-src 'self'",
}))
```

给 `XSSProtection`, `ContentTypeNosniff`, `XFrameOptions` or `ContentSecurityPolicy` 设定空值即禁用该规则。
