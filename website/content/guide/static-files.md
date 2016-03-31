---
title: Static Files
menu:
  side:
    parent: guide
    weight: 3
---

Images, JavaScript, CSS, PDF, Fonts and so on...

### Using Static Middleware

Static middleware can be used to serve static files from the provided root directory.

##### Usage

```go
e := echo.New()
e.Use(middleware.Static("public"))
```

This will serve any file from the public directory. For example, a request `/js/main.js`
will fetch and serve `public/js/main.js` file.

##### Caveat

With this setup, each and every request goes through this middleware which might
affect the performance. To overcome that, you can use `Echo#Group()` or [`Echo#Static()`]({{< relref "#using-echo-static">}})
APIs.

#### Using `Echo#Group()`

```go
e := echo.New()
static := e.Group("/static/*", middleware.Static("public"))
```

This will serve any file from the public directory. For example, a request `/static/js/main.js`
will fetch and serve `public/js/main.js` file.

#### Custom Configuration

##### Usage

```go
e := echo.New()
e.Use(middleware.StaticFromConfig(middleware.StaticConfig{
  Root:   "public",
  Browse: true,
  Index:  middleware.DefaultStaticConfig.Index,
}))
```

This uses `Root` as public directory to serve the static files and sets `Browse`
to true, enabling directory browsing.

##### Configuration

```go
StaticConfig struct {
  // Root is the directory from where the static content is served.
  Root string `json:"root"`

  // Index is the list of index files to be searched and used when serving
  // a directory.
  // Default value is `[]string{"index.html"}`.
  Index []string `json:"index"`

  // Browse is a flag to enable/disable directory browsing.
  Browse bool `json:"browse"`
}
```

##### Default Configuration

```go
DefaultStaticConfig = StaticConfig{
  Root:   "",
  Index:  []string{"index.html"},
  Browse: false,
}
```

### Using `Echo#Static()`

`Echo#Use(middleware.Static(root string))`

Serves static files from the provided `root` directory.

`Echo#Static(prefix, root string)`

Serves files from provided `root` directory for `/<prefix>*` HTTP path.

`Echo#File(path, file string)`

Serves provided `file` for `/<path>` HTTP path.

*Examples*

- Serving static files with no prefix `e.Use(middleware.Static("public"))`
- Serving static files with a prefix `e.Static("/static", "assets")`
- Serving an index page `e.File("/", "public/index.html")`
- Serving a favicon `e.File("/favicon.ico", "images/facicon.ico")`
