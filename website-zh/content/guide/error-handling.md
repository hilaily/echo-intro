+++
title: 错误处理
url : guide/error-handling
menu:
  side:
    parent: guide
    weight: 5

+++

## 错误处理程序

Echo 提倡通过中间件或处理程序 (handler) 返回 HTTP 错误集中处理。集中式错误处理程序允许我们从统一位置将错误记录到外部服务，并向客户端发送自定义 HTTP 响应。

你可以返回一个标准的 `error` 或者 `echo.*HTTPError`。

例如，当基本身份验证中间件找到无效凭据时，会返回 401未授权错误 (401-Unauthorized)，并终止当前的 HTTP 请求。

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
你也可以不带消息内容调用 `echo.NewHTTPError()`，这种情况下状态文本会被用作错误信息，例如 `Unauthorized`。

## 默认 HTTP 错误处理程序
Echo 提供了默认的 HTTP 错误处理程序，它用 JSON 格式发送错误。
```
{
  "message": "error connecting to redis"
}
```
标准错误 `error` 的响应是 `500 - Internal Server Error`。然而在调试 (debug) 模式模式下，原始的错误信息会被发送。如果错误是 `*HTTPError`，则使用设置的状态代码和消息发送响应。如果启用了日志记录，则还会记录错误消息。

## 自定义 HTTP 错误处理程序
通过 `e.HTTPErrorHandler` 可以设置自定义的 HTTP 错误处理程序 (error handler) 。

通常默认的 HTTP 错误处理程序已经够用；然而如果要获取不同类型的错误并采取相应的操作，则可以使用自定义 HTTP 错误处理程序，例如发送通知邮件或记录日志到应用中心的场景。最后，你还可以发送自定义的错误页面或 JSON 响应给客户端。

### 错误页
利用自定义 HTTP 错误处理程序，可以在显示不同种类的错误页面的同时，记录错误日志。错误页的名称可写作 `<CODE>.html`，例如 `500.html`。你可以在[https://github.com/AndiDittrich/HttpErrorPages](https://github.com/AndiDittrich/HttpErrorPages)看到 Echo 内置的错误页。
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
> 日志除了记录到 logger，也可以记录到第三方服务，例如 Elasticsearch 或者 Splunk。
