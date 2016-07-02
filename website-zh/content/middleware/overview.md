+++
title = "概述"
[menu.side]
  name = "概述"
  parent = "middleware"
  weight = 1
+++

## 中间件

中间件是一个函数，嵌入在HTTP 的请求和响应之间。它可以获得 `Echo#Context` 对象用来进行一些特殊的操作，
比如记录每个请求或者统计请求数。

Action的处理在所有的中间件运行完成之后。

### 中间件级别

#### Root Level (Before router)

`Echo#Pre()` 用于注册一个在路由执行之前运行的中间件，可以用来修改请求的一些属性。比如在请求路径结尾添加或者删除一个'/'来使之能与路由匹配。

下面的这几个内建中间件应该被注册在这一级别：

- AddTrailingSlash
- RemoveTrailingSlash
- MethodOverride

*注意*: 由于在这个级别路由还没有执行，所以这个级别的中间件不能调用任何 `echo.Context` 的 API。

#### Root Level (After router)

大部分时间你将用到 `Echo#Use()` 在这个级别注册中间件。
这个级别的中间件运行在路由处理完请求之后，可以调用所有的 `echo.Context` API。

下面的这几个内建中间件应该被注册在这一级别：

- BodyLimit
- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- Secure
- CORS
- Static

#### Group Level

当在路由中创建一个组的时候，可以为这个组注册一个中间件。例如，给 admin 这个组注册一个 BasicAuth 中间件。

*用法*

```go
e := echo.New()
admin := e.Group("/admin", middleware.BasicAuth())
```

也可以在创建组之后用 `admin.Use()`注册该中间件。

#### Route Level

当你创建了一个新的路由，可以选择性的给这个路由注册一个中间件。

*用法*

```go
e := echo.New()
e.GET("/", <Handler>, <Middleware...>)
```

### [Writing Custom Middleware]({{< ref "recipes/middleware.md">}})
