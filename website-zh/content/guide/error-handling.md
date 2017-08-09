---
title: 错误处理
url : guide/error-handling
menu:
  side:
    parent: guide
    weight: 5
---

## 错误处理

Echo 支持从中间件或者action返回 HTTP 错误集中处理。这样可以允许我们在统一的地方记录日志提供给第三方或者返回自定义的 HTTP 响应给客户端。
你可以返回一个标准的 `error` 或者 `echo.*HTTPError`

例如 一个基本的身份验证中间件验证失败返回 `401 - Unauthorized` 错误, 终止了当前的 HTTP 请求。

```go
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the credentials from HTTP request header and perform a security
			// check

			// For invalid credentials
			return echo.NewHTTPError(http.StatusUnauthorized)

			// For valid credentials call next
			// return next(c)
		}
	})
}
```
也可以不带消息内容调用 `echo.NewHTTPError()`，这样就会使用 error 带的错误信息，比如 `Unauthorized`。

## 默认的 http 错误处理
Echo 提供了默认的 http 错误处理，它以 json 格式发送数据。
```
{
  "message": "error connecting to redis"
}
```
一个错误是 golang 标准的 error，响应是 `500 - Internal Server Error`。当然，如果是在 debug 模式，原始的错误信息会被发送。如果错误是 `*HTTPError`,
响应会有提供的错误码和错误内容。如果日志打开了，错误信息也会被日志记录。

## 自定义 http 错误处理
通过 `e.HTTPErrorHandler` 可以设置自定义的 http 错误处理。
在大部分情况下，默认的 http 错误处理已经够用了，当时如果想根据不同的错误做不同的处理的时候就需要自定义错误处理了。比如发送提醒邮件或者记录日志到应用中心。还可以发送自定义的错误响应给客户端，比如定义错误页面或者返回一段 json 数据。

## 错误页
下面的自定义 http 错误处理器展示了怎么根据不同的错误显示不一样的错误页面和记录日志。错误页的名字应该类似 `<CODE>.html`，比如 `500.html`。你可以在[这里](https://github.com/AndiDittrich/HttpErrorPages)看到 Echo 内置的错误页。
```go
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

e.HTTPErrorHandler = customHTTPErrorHandler
```
> 除了记录日志到 logger，你也可以将日志记录到 Elasticsearch 或者 Splunk。
