---
title: Static Files
menu:
  side:
    parent: guide
    weight: 3
---

### Serving static files using `Static` middleware

Static middleware serves static content from the provided root directory.

#### Configuration

```go
StaticConfig struct {
  // Root is the directory from where the static content is served.
  Root string `json:"root"`

  // Index is the index file to be used for directory browsing.
  // Default value is `index.html`.
  Index string `json:"index"`

  // Browse is a flag to enable/disable directory browsing.
  Browse bool `json:"browse"`
}
```

#### Default Configuration

```go
DefaultStaticConfig = StaticConfig{
  Root: "",
  Index:  "index.html",
  Browse: false,
}
```

##### Usage

`e.Use(middleware.Static("public"))`

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

Example above uses `Root` as public directory to serve static files and sets `Browse`
to true, enabling directory browsing.

### Serving static files using `Echo#Static()`

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
