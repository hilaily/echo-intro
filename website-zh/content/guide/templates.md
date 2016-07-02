---
title: 模板
menu:
  side:
    parent: guide
    weight: 5
---

## 模板

### 模板渲染

`Context#Render(code int, name string, data interface{}) error` 用于渲染一个模板，然后发送一个 text/html 的状态响应。我们可以使用任何模板引擎，只要使用`Echo.SetRenderer()`注册。

下面是使用Go `html/template`的示例：

1. 实现 `echo.Render` 接口

    ```go
    type Template struct {
        templates *template.Template
    }

    func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    	return t.templates.ExecuteTemplate(w, name, data)
    }
    ```

2. 预编译模板

    `public/views/hello.html`

    ```html
    {{define "hello"}}Hello, {{.}}!{{end}}
    ```

    ```go
    t := &Template{
        templates: template.Must(template.ParseGlob("public/views/*.html")),
    }
    ```

3. 注册模板

    ```go
    e := echo.New()
    e.SetRenderer(t)
    e.GET("/hello", Hello)
    ```

4. 在action中渲染模板

    ```go
    func Hello(c echo.Context) error {
    	return c.Render(http.StatusOK, "hello", "World")
    }
    ```
