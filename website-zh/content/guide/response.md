---
title: 响应
url: guide/response
menu:
  side:
    parent: guide
    weight: 8
---

## 响应

### 发送 string 数据

`Context#String(code int, s string)` 用于发送一个带有状态码的纯文本响应。

*示例*

```go
func(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
```

### 发送 HTML 响应（参考模板）

`Context#HTML(code int, html string)` 用于发送一个带有状态码的简单 HTML 响应。如果你需要动态生成 HTML 内容请查看[模版](https://echo.labstack.com/guide/templates)。

*示例*

```go
func(c echo.Context) error {
  return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
}
```

#### 发送 HTML Blob

`Context#HTMLBlob(code int, b []byte)` 用于发送一个带状态码的 HTML blob（二进制长对象）响应。可以发现，使用输出 `[]byte` 的模版引擎很方便。

###  模版引擎渲染

[查看](https://echo.labstack.com/guide/templates)

### 发送 JSON 数据

`Context#JSON(code int, i interface{})` 用于发送一个带状态码的 JSON 对象，它会将 Golang 的对象转换成 JSON 字符串。

*示例*

```go
// User
type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

// Handler
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  return c.JSON(http.StatusOK, u)
}
```

#### JSON 流

`Context#JSON()` 内部使用 `json.Marshal` 来转换 JSON 数据，但该方法面对大量的 JSON 数据会显得效率不足，对于这种情况可以直接使用 JSON 流。

*示例*

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(u)
}
```

#### JSON 美化（JSON Pretty）

`Context#JSONPretty(code int, i interface{}, indent string)` 可以发送带有缩进（可以使用空格和 tab）的更为好看的 JSON 数据。

*示例*

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "joe@labstack.com",
  }
  return c.JSONPretty(http.StatusOK, u, "  ")
}
```

```go
{
  "email": "joe@labstack.com",
  "name": "Jon"
}
```

> 通过在请求URL查询字符串中附加 `pretty` ，你也可以使用 `Context#JSON()` 来输出带有缩进的 JSON 数据。

*示例*

```bash
curl http://localhost:1323/users/1?pretty
```


#### JSON Blob

`Context#JSONBlob(code int, b []byte)` 可用来从外部源（例如数据库）直接发送预编码的 JSON 对象。

*示例*

```go
func(c echo.Context) error {
  encodedJSON := []byte{} // Encoded JSON from external source
  return c.JSONBlob(http.StatusOK, encodedJSON)
}
```

### 发送 JSONP 数据

`Context#JSONP(code int, callback string, i interface{})`  可以将 Golang 的数据类型转换成 JSON 类型，并通过回调以带有状态码的 JSONNP 结构发送。

[查看示例](https://echo.labstack.com/cookbook/jsonp)

### 发送 XML 数据

`Context#XML(code int, i interface{})` 可以将 Golang 对象转换成 XML 类型，并带上状态码发送响应。

*示例*

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  return c.XML(http.StatusOK, u)
}
```

#### XML 流

`Context#XML` 内部使用 `xml.Marshal` 来转换 XML 数据，但该方法面对大量的 XML 数据会显得效率不足，对于这种情况可以直接使用 XML 流。

*示例*

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return xml.NewEncoder(c.Response()).Encode(u)
}
```

#### XML 美化（XML Pretty）

`Context#XMLPretty(code int, i interface{}, indent string)` 可以发送带有缩进（可以使用空格和 tab）的更为好看的  XML 数据。

*示例*

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "joe@labstack.com",
  }
  return c.XMLPretty(http.StatusOK, u, "  ")
}
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<User>
  <Name>Jon</Name>
  <Email>joe@labstack.com</Email>
</User>
```

> 通过在请求URL查询字符串中附加 `pretty` ，你也可以使用 `Context#XML()` 来输出带有缩进的 XML 数据。

*示例*

```bash
curl http://localhost:1323/users/1?pretty
```

#### XML Blob

`Context#XMLBlob(code int, b []byte)` 可用来从外部源（例如数据库）直接发送预编码的 XML 对象。

*示例*

```go
func(c echo.Context) error {
  encodedXML := []byte{} // Encoded XML from external source
  return c.XMLBlob(http.StatusOK, encodedXML)
}
```

### 发送文件

`Context#File(file string)` 可用来发送内容为文件的响应，并且它能自动设置正确的内容类型、优雅地处理缓存。

*示例*

```go
func(c echo.Context) error {
  return c.File("<文件路径>")
}
```

### 发送附件

`Context#Attachment(file, name string)` 和发送文件 `File()` 的方法类似，只是它的方法名称不同。

*示例*

```go
func(c echo.Context) error {
  return c.Attachment("<PATH_TO_YOUR_FILE>")
}
```

### 发送内嵌（Inline）

`Context#Inline(file, name string)` 和发送文件 `File()` 的方法类似，只是它的方法名称不同。

*示例*

```go
func(c echo.Context) error {
  return c.Inline("<PATH_TO_YOUR_FILE>")
}
```

### 发送二进制长文件（Blob）

``Context#Blob(code int, contentType string, b []byte)`  可用于发送带有内容类型(content type)和状态代码的任意类型数据。

*示例*

```go
func(c echo.Context) (err error) {
  data := []byte(`0306703,0035866,NO_ACTION,06/19/2006
	  0086003,"0005866",UPDATED,06/19/2006`)
	return c.Blob(http.StatusOK, "text/csv", data)
}
```

### 发送流（Stream）

`Context#Stream(code int, contentType string, r io.Reader)` 可用于发送带有内容类型(content type)、状态代码、`io.Reader` 的任意类型数据流。

*示例*

```go
func(c echo.Context) error {
  f, err := os.Open("<PATH_TO_IMAGE>")
  if err != nil {
    return err
  }
  return c.Stream(http.StatusOK, "image/png", f)
}
```

### 发送空内容（No Content）

`Context#NoContent(code int)` 可用于发送带有状态码的空内容。

*示例*

```go
func(c echo.Context) error {
  return c.NoContent(http.StatusOK)
}
```

### 重定向

`Context#Redirect(code int, url string)` 可用于重定向至一个带有状态码的 URL。

*示例*

```go
func(c echo.Context) error {
  return c.Redirect(http.StatusMovedPermanently, "<URL>")
}
```

### Hooks

#### 响应之前

`Context#Response#Before(func())` 可以用来注册在写入响应之前调用的函数。

#### 响应之后

`Context#Response#After(func())` 可以用来注册在写入响应之后调用的函数。但是如果“Content-Length”是未知状态，则不会有任何方法会被执行。

*示例*

```go
func(c echo.Context) error {
  c.Response().Before(func() {
    println("before response")
  })
  c.Response().After(func() {
    println("after response")
  })
  return c.NoContent(http.StatusNoContent)
}
```

> 可以在响应之前与之后注册多个方法