---
title: 迁移
url : guide/migrating
menu:
  side:
    parent: guide
    weight: 6
---

## 从 v1 迁移

### 更新日志

- 使用 Let's Encrypt 自动生成证书
- 支持优雅关闭服务
- 多功能的函数来包裹标准处理器和中间件(Utility functions to wrap standard handler and middleware)
- Map 类型简单表述为 `map[string]interface{}`
- Context 现在内嵌了标准库 `net/http` 的请求和响应类型
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
- Echo 实例增加了这些可导出的属性，不再使用 setter/getter 方法
  - Binder
  - Renderer
  - HTTPErrorHandler
  - Debug
  - Logger 
- 改善重定向和 CORS 中间件
- 由于 `Echo#Static` 的存在废除了 static 中间件
- 废除的 API
  - Echo#Run()
  - Context#P()
- 废除了标准 `Context` 的支持
- 废除了 `fasthttp` 
- 去除了标记为 deprecated 的 API
- 将 `Logger` 接口移到 root 级别
- 将网站和示例移动到主仓库
- 更新文档以及修复了大量 issues。  