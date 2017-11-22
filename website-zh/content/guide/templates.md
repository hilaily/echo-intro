---
title: 模板
url : guide/templates
menu:
  side:
    parent: guide
    weight: 11
---

## 模板

### 模板渲染

`Context#Render(code int, name string, data interface{}) error` 用于渲染一个模板，然后发送一个 text/html 的状态响应。我们可以使用任何模板引擎，只要赋值给 `Echo.Renderer`。

下面是使用Go `html/template` 的示例：

1.实现 `echo.Renderer` 接口

```go
type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}
```

2.预编译模板

`public/views/hello.html`

```html
{{define "hello"}}Hello, {{.}}!{{end}}
```

```go
t := &Template{
	templates: template.Must(template.ParseGlob("public/views/*.html")),
}
```

3.注册模板

```go
e := echo.New()
e.Renderer = t
e.GET("/hello", Hello)
```

4.在 action 中渲染模板

```go
 func Hello(c echo.Context) error {
 	return c.Render(http.StatusOK, "hello", "World")
 }
```

### 高级 - 在模版中调用 Echo
在某些情况下可能需要从模版生成 uri。这样你需要在模版中调用 `Echo#Reverse`。Golang 的 `html/template` 包不太适合于这种情况，但是我们可以通过两种方法实现它：给所有的传递到模版的对象提供一个公用的方法或者将 `map[string]interface{}` 作为参数传递给自定义模版。
下面的代码展示后者的处理方式：
`template.html`
```html
<html>
    <body>
        <h1>Hello {{index . "name"}}</h1>

        <p>{{ with $x := index . "reverse" }}
           {{ call $x "foobar" }} &lt;-- this will call the $x with parameter "foobar"
           {{ end }}
        </p>
    </body>
</html>
```
`server.go`
```go
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
  e := echo.New()
  renderer := &TemplateRenderer{
      templates: template.Must(template.ParseGlob("*.html")),
  }
  e.Renderer = renderer

  // Named route "foobar"
  e.GET("/something", func(c echo.Context) error {
      return c.Render(http.StatusOK, "something.html", map[string]interface{}{
          "name": "Dolly!",
      })
  }).Name = "foobar"

  log.Fatal(e.Start(":8000"))
}
```