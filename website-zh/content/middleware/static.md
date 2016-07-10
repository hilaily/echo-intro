+++
title = "Static"
[menu.side]
  name = "Static"
  parent = "middleware"
  weight = 5
+++

## Static 中间件

Static 中间件用于从给定的根目录提供静态资源服务。

### 配置

```go
StaticConfig struct {
  // Root 目录是静态资源所在的目录。
  // 必须配置。
  Root string `json:"root"`

  // Index file for serving a directory.
  // 目录的默认索引文件(index file)
  // 可选。默认为 "index.html"。
  Index string `json:"index"`

  // 开启 HTML5 模式，所有无法匹配的请求都重定向到root目录使单页面应用可以处理该请求。
  // 可选。默认为 "false"。
  HTML5 bool `json:"html5"`

  // 允许浏览目录。
  // 可选。默认为 false。
  Browse bool `json:"browse"`
}
```

### 默认配置

```go
DefaultStaticConfig = StaticConfig{
  Index: "index.html",
}
```

*用法*

```go
e := echo.New()
e.Use(middleware.Static("/static"))
```

这里将`static`作为静态资源目录。例如，一个访问`/js/main.js`的请求将匹配到`static/js/main.js`文件。

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
  Root:   "static",
  Browse: true,
}))
```

这里从`static`目录提供静态资源访问。并且允许浏览该目录。
