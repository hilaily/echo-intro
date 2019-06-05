---
title: 迁移
url : guide/migrating
menu:
  side:
    parent: guide
    weight: 6
---

## 迁移

### 更新日志

- 通过 [Let’s Encrypt](https://letsencrypt.org/) 自动生成 TLS 证书
- 内置优雅停机(graceful shutdown)
- 提供用于封装标准处理程序(standard handler)和中间件(middleware)的功能函数
- `Map` 类型简单表述为 `map[string]interface{}`
- 上下文(context)现在封装了标准的 `net/http` 请求与响应
- 新的配置
  - `Echo#ShutdownTimeout`
  - `Echo#DisableHTTP2`
- 新的 API
  - `Echo#Start()`
  - `Echo#StartTLS()`
  - `Echo#StartAutoTLS()`
  - `Echo#StartServer()`
  - `Echo#Shutdown()`
  - `Echo#ShutdownTLS()`
  - `Context#Scheme()`
  - `Context#RealIP()`
  - `Context#IsTLS()`
- Echo 利用以下属性替代 setter/getter 方法
  - Binder
  - Renderer
  - HTTPErrorHandler
  - Debug
  - Logger 
- 改善重定向和 CORS 中间件
- 由于 `Echo#Static` 的存在废除了 static 中间件
- 删除 API
  - Echo#Run()
  - Context#P()
- 删除标准 `Context` 支持
- 删除 `fasthttp` 
- 删除标记为 deprecated 的 API
- `Logger` 接口移至 root 级别
- 将网站和示例移至主仓库
- 更新文档以及修复了大量问题(issues)

## [Cookbook](https://echo.labstack.com/cookbook)