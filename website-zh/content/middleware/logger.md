+++
title = "Logger(日志)"
[menu.side]
  name = "Logger(日志)"
  parent = "middleware"
  weight = 5
+++

## Logger(日志) 中间件

Logger 中间件记录了每一个请求的信息。

### 配置

```go
LoggerConfig struct {
  // 日志的格式可以使用下面的标签定义。:
  //
  // - time_rfc3339
  // - id (Request ID - Not implemented)
  // - remote_ip
  // - uri
  // - host
  // - method
  // - path
  // - referer
  // - user_agent
  // - status
  // - latency (In microseconds)
  // - latency_human (Human readable)
  // - rx_bytes (Bytes received)
  // - tx_bytes (Bytes sent)
  //
  // 例如 "${remote_ip} ${status}"
  //
  // 可选。默认值是 DefaultLoggerConfig.Format.
  Format string

  // Output 是记录日志的位置。
  // 可选。默认值是 os.Stdout.
  Output io.Writer
}
```

### 默认配置

```go
DefaultLoggerConfig = LoggerConfig{
  Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}",` +
    `"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
    `"latency_human":"${latency_human}","rx_bytes":${rx_bytes},` +
    `"tx_bytes":${tx_bytes}}` + "\n",
  color:  color.New(),
  Output: os.Stdout,
}
```

*用法*

`e.Use(middleware.Logger())`

*输出样例*

```js
{"time":"2016-05-10T07:02:25-07:00","remote_ip":"::1","method":"GET","uri":"/","status":200, "latency":55653,"latency_human":"55.653µs","rx_bytes":0,"tx_bytes":13}
```

### 自定义配置

*用法*

```go
e := echo.New()
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "method=${method}, uri=${uri}, status=${status}\n",
}))
```

*输出样例*

```sh
method=GET, uri=/hello, status=200
```
