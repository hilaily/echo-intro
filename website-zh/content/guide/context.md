+++
title = "Context"
url= "/guide/context"
[menu.side]
  name= "Context"
  parent = "guide"
  weight = 5
+++

## Context

echo.Context 代表了当前 HTTP 请求的context（上下文？这里看个人理解吧，就不翻译了）。
它含有请求和相应的引用，路径，路径参数，数据，注册的业务处理方法和 读取请求和输出响应的API。
Context 100% 和标准的 context.Context 兼容。
由于 Context 是一个接口，所以也可以很方便的使用自定义的 API 扩展。

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
### 标准的 Context
echo.Context 内嵌了标准的 context.Context 接口，所以 context.Context 的所有方法 echo.Context 都可以使用。

*示例*

```go
e.GET("/users/:name", func(c echo.Context) error) {
    c.SetContext(context.WithValue(nil, "key", "val"))
    // Pass it down...
    // Use it...
    val := c.Value("key").(string)
    return c.String(http.StatusOK, name)
})
```
