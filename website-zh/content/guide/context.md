+++
title = "Context"
url= "/guide/context"
[menu.side]
  name= "Context"
  parent = "guide"
  weight = 3

+++

## Context

echo.Context 表示当前 HTTP 请求的上下文。通过路径、路径参数、数据、注册处理程序和相关 API 进行请求的读取与响应的输出。由于 Context 是一个接口，也可以轻松地使用自定义 API 进行扩展。

### 扩展 Context

**自定义一个 context**
```go
type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}
```
**创建一个中间件来扩展默认的 context**

```go
e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}
		return h(cc)
	}
})
```
> 这个中间件要在所有其它中间件之前注册到路由上。

**在业务处理中使用**

```go
e.Get("/", func(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.Foo()
	cc.Bar()
	return cc.String(200, "OK")
})
```
