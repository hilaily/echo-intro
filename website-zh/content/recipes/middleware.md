---
title: Middleware
url: middleware
menu:
  side:
    identifier: "recipes-middleware"
    parent: recipes
    weight: 3
---

## 中间件示例

### 如何写一个自定义中间件？

- 使用中间件收集请求数、状态和正常运行时间。
- 使用中间件来写自定义服务器响应头。

#### 服务端

`server.go`

{{< embed "middleware/server.go" >}}

#### 响应

*Headers*
```sh
Content-Length:122
Content-Type:application/json; charset=utf-8
Date:Thu, 14 Apr 2016 20:31:46 GMT
Server:Echo/2.0
```

*Body*

```js
{
  "uptime": "2016-04-14T13:28:48.486548936-07:00",
  "requestCount": 5,
  "statuses": {
    "200": 4,
    "404": 1
  }
}
```

### 维护者

- [vishr](https://github.com/vishr)

### [源码]({{< source "middleware" >}})
