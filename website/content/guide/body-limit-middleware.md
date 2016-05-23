+++
title = "Body Limit Middleware"
[menu.side]
  name = "BodyLimit"
  parent = "middleware"
  weight = 5
+++

## BodyLimit Middleware

BodyLimit middleware sets the maximum allowed size for a request body, if the
size exceeds the configured limit, it sends "413 - Request Entity Too Large"
response. The body limit is determined based on both `Content-Length` request
header and actual content read, which makes it super secure.

Limit can be specified as `4x` or `4xB`, where x is one of the multiple from K, M,
G, T or P.

*Usage*

```go
e := echo.New()
e.Use(middleware.BodyLimit("2M"))
```
