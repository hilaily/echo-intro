+++
title = "FAQ 常见问题"
url = "guide/faq"
[menu.side]
  name = "FAQ 常见问题"
  parent = "guide"
  weight = 20
+++

Q: 怎样从 `echo.Context` 获取  `*http.Request` 和 `http.ResponseWriter`  ?

- `http.Request` > `c.Request()`
- `http.ResponseWriter` > `c.Response()`

Q: 在 Echo 中怎么使用 go 的标准控制器 `func(http.ResponseWriter, *http.Request)` ?

```go
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Echo!")
}

func main() {
	e := echo.New()
	e.GET("/", echo.WrapHandler(http.HandlerFunc(handler)))
	e.Start(":1323")
}
```

Q: 在 Echo 中怎么使用标准的中间件 `func(http.Handler) http.Handler` ?

```go
func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("middleware")
		h.ServeHTTP(w, r)
	})
}

func main() {
	e := echo.New()
	e.Use(echo.WrapMiddleware(middleware))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo!")
	})
	e.Start(":1323")
}
```

Q: 怎样在指定的 ip 地址运行 Echo ?

```go
e.Start("<ip>:<port>")
```

Q: 怎么在一个随机的端口运行 Echo ?

```go
e.Start(":0")
```
