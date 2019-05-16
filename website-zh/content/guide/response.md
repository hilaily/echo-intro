---
title: HTTP 响应
url: guide/response
menu:
  side:
    parent: guide
    weight: 8
---

## HTTP 响应

### 发送 string 数据

`Context#String(code int, s string)` 用于发送一个带有状态码的纯文本响应。

```go
func(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
```

### 发送 HTML 响应

`Context#HTML(code int, html string)` 用于发送一个带状态码的简单 html 响应。如果你需要动态生成 html 内容请查看[模版](https://echo.labstack.com/guide/templates)。

```go
func(c echo.Context) error {
  return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
}
```

### 发送 HTML Blog

`Context#HTMLBlob(code int, b []byte)` 用于发送一个带状态码的 html blob(二进制大对象)响应。它和输出 []byte 类型内容的模版引擎配合使用非常方便。

###  模版引擎渲染

[查看](https://echo.labstack.com/guide/templates)

### 发送 JSON 数据

`Context#JSON(code int, i interface{})` 用于发送一个带状态码的 json 对象。它会将 golang 的对象转换成 json 字符串。

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

### JSON 流

`Context#JSON()` 内部使用 `json.Marshl` 来转换 json 数据，对于多大的数据来说性能不够好，这种情况你可以直接使用 json 流。

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

### JSON Pretty

`Context#JSONPretty(code int, i interface{}, indent string)` 也是用于发送 json 数据。不过它打印出的 json 数据带有缩进（可以使用空格和 tab），更为好看。

发送带有空格锁进的 json 数据。

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

### JSON Blob

`Context#JSONBlob(code int, b []byte)` 用来直接发送一个已经转换好的 json 对象。

```go
func(c echo.Context) error {
  encodedJSON := []byte{} // Encoded JSON from external source
  return c.JSONBlob(http.StatusOK, encodedJSON)
}
```

### 发送 JSONP 数据

`Context#JSONP(code int, callback string, i interface{})`  用来把 golang 的数据类型转换成 json 并通过回调以 jsonp 的结构发送出去。

[示例](https://echo.labstack.com/examples/jsonp)

### 发送 XML 数据

`Context#XML(code int, i interface{})` 用来转换 golang 对象为 xml 数据发送响应。

```go
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  return c.XML(http.StatusOK, u)
}
```

### Stream XML

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

### XML Pretty

`Context#XMLPretty(code int, i interface{}, indent string)`

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

### XML Blob

`Context#XMLBlob(code int, b []byte)` 

```go
func(c echo.Context) error {
  encodedXML := []byte{} // Encoded XML from external source
  return c.XMLBlob(http.StatusOK, encodedXML)
}
```

### 发送文件

`Context#File(file string)` 用来发送一个文件为内容的响应。

```go
func(c echo.Context) error {
  return c.File("<文件路径>")
}
```

### 发送附件

`Context#Attachment(file, name string)` 和发送文件的方法类似，只是它会多提供一个名称。

```go
func(c echo.Context) error {
  return c.Attachment("<PATH_TO_YOUR_FILE>")
}
```

### Send Inline

`Context#Inline(file, name string)` 

```go
func(c echo.Context) error {
  return c.Inline("<PATH_TO_YOUR_FILE>")
}
```

### Send Blob

``Context#Blob(code int, contentType string, b []byte)`  用来发送任意类型的数据。需要提供 content type。

```go
func(c echo.Context) (err error) {
  data := []byte(`0306703,0035866,NO_ACTION,06/19/2006
	  0086003,"0005866",UPDATED,06/19/2006`)
	return c.Blob(http.StatusOK, "text/csv", data)
}
```

### 发送流数据

`Context#Stream(code int, contentType string, r io.Reader)` 用来发送任意数据流响应。需要提供 content type，io.Reader 和状态码。

```go
func(c echo.Context) error {
  f, err := os.Open("<PATH_TO_IMAGE>")
  if err != nil {
    return err
  }
  return c.Stream(http.StatusOK, "image/png", f)
}
```

### 发送空内容

`Context#NoContent(code int)`

```go
func(c echo.Context) error {
  return c.NoContent(http.StatusOK)
}
```

### 重定向

`Context#Redirect(code int, url string)`，提供一个 url 用于重定向。

```go
func(c echo.Context) error {
  return c.Redirect(http.StatusMovedPermanently, "<URL>")
}
```

