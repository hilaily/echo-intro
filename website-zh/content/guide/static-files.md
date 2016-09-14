---
title: 静态文件
url : static-files
menu:
  side:
    parent: guide
    weight: 4
---

图片，JavaScript，CSS，PDF，字体文件等等...

## 静态文件

### [使用 Static 中间件]({{< ref "middleware/static.md">}})

### 使用`Echo#Static()`

`Echo#Static(prefix, root string)` 用一个路径前缀注册了一个新的路由来提供静态文件的访问服务。改路径前缀作为根目录。

*用法 1*

```go
e := echo.New()
e.Static("/static", "assets")
```

这样会将所有访问`/static/*`的请求去 访问`assets`目录。例如，一个访问`/static/js/main.js`的请求会匹配到`assets/js/main.js`这个文件。

*用法 2*

```go
e := echo.New()
e.Static("/", "assets")
```

这样会将所有`assets`目录的文件使用`/*`去访问。例如，一个访问`/js/main.js`的请求将会匹配到`assets/js/main.js`文件。

### 使用`Echo#File()`

`Echo#File(path, file string)` 使用一个路径注册一个新的路由去访问某个静态文件。

*用法 1*

将`public/index.html`作为主页。

```go
e.File("/", "public/index.html")
```

*用法 2*

给`images/favicon.ico`一个静态路径。

```go
e.File("/favicon.ico", "images/favicon.ico")
```
