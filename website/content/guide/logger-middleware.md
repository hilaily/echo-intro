+++
title = "Logger Middleware"
[menu.side]
  name = "Logger"
  parent = "middleware"
  weight = 5
+++

## Logger Middleware

Logger middleware logs the information about each HTTP request.

### Configuration

```go
LoggerConfig struct {
  // Format is the log format which can be constructed using the following tags:
  //
  // - time_rfc3339
  // - remote_ip
  // - uri
  // - method
  // - path
  // - status
  // - response_time
  // - response_size
  //
  // Example "${remote_id} ${status}"
  //
  // Optional, with default value as `DefaultLoggerConfig.Format`.
  Format string

  // Output is the writer where logs are written.
  // Optional, with default value as os.Stdout.
  Output io.Writer
}
```

### Default Configuration

```go
DefaultLoggerConfig = LoggerConfig{
  Format: `{"time": "${time_rfc3339}", "remote_ip": "${remote_ip}", ` +
    `"method": "${method}", "uri": "${uri}", "status": ${status}, ` +
    `"response_time": "${response_time}", "response_size": "${response_size}B"}` +
    "\n",
  color:  color.New(),
  Output: os.Stdout,
}
```

*Usage*

`e.Use(middleware.Logger())`

*Sample Output*

```js
{"time": "2016-05-09T19:19:42-07:00", "remote_ip": "::1", "method": "GET", "uri": "/", "status": 200, "response_time": "63.82µs", "response_size": "13B"}
```

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "method=${method}, uri=${uri}, status=${status}\n",
}))
```

Example above uses a `Format` which logs request method and request URI.

*Sample Output*

`method=GET, uri=/hello, status=200`
