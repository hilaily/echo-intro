---
title: Static Files
menu:
  side:
    parent: guide
    weight: 3
---

Images, JavaScript, CSS, PDF, Fonts and so on...

### Using `Echo#Static()`

`Echo#Static(path, root string)` serves static files from the provided `root`
directory for path `/<path>*`.

#### Example 1

```go
e := echo.New()
e.Static("/static", "assets")
```

This will serve any file from the assets directory for path `/static/*`. For example,
a request to `/static/js/main.js` will fetch and serve `assets/js/main.js` file.

#### Example 2

```go
e := echo.New()
e.Static("/", "assets")
```

This will serve any file from the assets directory for path `/*`. For example,
a request to `/js/main.js` will fetch and serve `assets/js/main.js` file.

### Using `Echo#StaticWithConfig()`

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

##### Example

```go
e := echo.New()
e.StaticWithConfig("", echo.StaticConfig{
  Root:   "public",
  Browse: true,
})
```

This uses `Root` as public directory to serve the static files and sets `Browse`
to true, enabling directory browsing.

### Using `Echo#File()`

`Echo#File(path, file string)` serves static file from the provided `file` for `/path`.

##### Example 1

Serving an index page from `public/index.html`

```go
e.File("/", "public/index.html")
```

##### Example 2

Serving a favicon from `images/facicon.ico`

```go
e.File("/favicon.ico", "images/facicon.ico")
```
