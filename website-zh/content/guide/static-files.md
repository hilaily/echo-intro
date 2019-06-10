---
title: 静态文件
url : guide/static-files
menu:
  side:
    parent: guide
    weight: 10
---

## 静态文件

例如图片，JavaScript，CSS，PDF，字体文件等等...

### [使用静态中间件](https://echo.labstack.com/middleware/static/)

### 使用 Echo#Static()

`Echo#Static(prefix, root string)` 使用路径前缀注册一个新路由，以便由根目录提供静态文件。

*用法 1*

```go
e := echo.New()
e.Static("/static", "assets")
```

如上所示， assets 目录中 `/static/*` 路径下的任何文件都会被访问。例如，一个访问 `/static/js/main.js` 的请求会匹配到 `assets/js/main.js` 这个文件。

*用法 2*

```go
e := echo.New()
e.Static("/", "assets")
```

如上所示，  assets 目录中 `/*` 路径下的任何文件都会被访问。例如，一个访问 `/js/main.js` 的请求将会匹配到 `assets/js/main.js` 文件。

### 使用 Echo#File()

`Echo#File(path, file string)` 使用路径注册新路由以提供静态文件。

*用法 1*

使用 `public/index.html` 提供索引页面

```go
e.File("/", "public/index.html")
```

*用法 2*

使用 `images/favicon.ico` 提供一个图标

```go
e.File("/favicon.ico", "images/favicon.ico")
```
