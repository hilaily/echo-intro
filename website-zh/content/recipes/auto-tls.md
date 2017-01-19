---
title: Auto TLS
url : /recipes/auto-tls
menu:
  side:
    parent: cookbook
    weight: 1
---

## Auto TLS

这个例子演示如何自动从 Let's Encrypt 获得 TLS 证书。 `Echo#StartAutoTLS` 接受一个接听 443 端口的网络地址。类似 `<DOMAIN>:443` 这样。

如果没有错误，访问 `https://<DOMAIN>` ，可以看到一个 TLS 加密的欢迎界面。

### 服务器

`server.go`

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<your_domain>")
	// Store the certificate to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits/)
	// e.AutoTLSManager.Cache = autocert.DirCache("<path to store key and certificate>")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})
	e.Logger.Fatal(e.StartAutoTLS(":443"))
}
```