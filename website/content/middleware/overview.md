+++
title = "Overview"
[menu.side]
  name = "Overview"
  parent = "middleware"
  weight = 1
+++

## Middleware Overview

Middleware is a function chained in the HTTP request-response cycle with access
to `Echo#Context` which it uses to perform a specific action, for example, logging
every request or limiting the number of requests.

Handler is processed in the end after all middleware are finished executing.

### Middleware Levels

#### Root Level (Before router)

`Echo#Pre()` can be used to register a middleware which is executed before router
processes the request. It is helpful to make any changes to the request properties,
for example, adding or removing a trailing slash from the path so it matches the
route.

The following built-in middleware should be registered at this level:

- AddTrailingSlash
- RemoveTrailingSlash
- MethodOverride

*Note*: As router has not processed the request, middleware at this level won't
have access to any path related API from `echo.Context`.

#### Root Level (After router)

Most of the time you will register a middleware at this level using `Echo#Use()`.
This middleware is executed after router processes the request and has full access
to `echo.Context` API.

The following built-in middleware should be registered at this level:

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

When creating a new group, you can register middleware just for that group. For
example, you can have an admin group which is secured by registering a BasicAuth
middleware for it.

*Usage*

```go
e := echo.New()
admin := e.Group("/admin", middleware.BasicAuth())
```

You can also add a middleware after creating a group via `admin.Use()`.

#### Route Level

When defining a new route, you can optionally register middleware just for it.

*Usage*

```go
e := echo.New()
e.GET("/", <Handler>, <Middleware...>)
```

Example above uses a `Level` 5 for gzip compression.

### [Writing Custom Middleware]({{< ref "recipes/middleware.md">}})
