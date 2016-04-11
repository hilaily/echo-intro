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

#### Configuration

```go
// StaticConfig defines the config for static middleware.
StaticConfig struct {
  // Root is the directory from where the static content is served.
	// Required.
  Root string `json:"root"`

  // Index is the list of index files to be searched and used when serving
  // a directory.
  // Optional with default value as []string{"index.html"}.
  Index []string `json:"index"`

  // Browse is a flag to enable/disable directory browsing.
  // Optional with default value as false.
  Browse bool `json:"browse"`
}
```

#### Default Configuration

```go
DefaultStaticConfig = StaticConfig{
  Index:  []string{"index.html"},
  Browse: false,
}
```

*Usage*

```go
e := echo.New()
e.Use(middleware.Static("/static"))
```

This serves static files from `static` directory. For example, a request to `/js/main.js`
will fetch and serve `static/js/main.js` file.

#### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
  Root:   "static",
  Browse: true,
}))
```

This serves static files from `static` directory and enables directory browsing.

### Using `Echo#Static()`

`Echo#Static(prefix, root string)` registers a new route with path prefix to serve
static files from the provided root directory.

*Usage 1*

```go
e := echo.New()
e.Static("/static", "assets")
```

This will serve any file from the assets directory for path `/static/*`. For example,
a request to `/static/js/main.js` will fetch and serve `assets/js/main.js` file.

*Usage 2*

```go
e := echo.New()
e.Static("/", "assets")
```

This will serve any file from the assets directory for path `/*`. For example,
a request to `/js/main.js` will fetch and serve `assets/js/main.js` file.

### Using `Echo#File()`

`Echo#File(path, file string)` registers a new route with path to serve a static
file.

*Usage 1*

Serving an index page from `public/index.html`

```go
e.File("/", "public/index.html")
```

*Usage 2*

Serving a favicon from `images/facicon.ico`

```go
e.File("/favicon.ico", "images/facicon.ico")
```
