+++
title = "Static Middleware"
[menu.side]
  name = "Static"
  parent = "middleware"
  weight = 5
+++

## Static Middleware

Static middleware can be used to serve static files from the provided root directory.

### Configuration

```go
// StaticConfig defines the config for static middleware.
StaticConfig struct {
  // Root is the directory from where the static content is served.
	// Required.
  Root string `json:"root"`

  // Index is the list of index files to be searched and used when serving
  // a directory.
  // Optional, with default value as []string{"index.html"}.
  Index []string `json:"index"`

  // Browse is a flag to enable/disable directory browsing.
  // Optional, with default value as false.
  Browse bool `json:"browse"`
}
```

### Default Configuration

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

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
  Root:   "static",
  Browse: true,
}))
```

This serves static files from `static` directory and enables directory browsing.
