+++
title = "Recover 恢复"
url = "/middleware/recover"
[menu.side]
  name = "Recover 恢复"
  parent = "middleware"
  weight = 5
+++

## Recover 中间件

Recover 中间件从 panic 链中的任意位置恢复程序， 打印堆栈的错误信息，并将错误集中交给 
[HTTPErrorHandler](https://echo.labstack.com/guide/customization#http-error-handler) 处理。

*使用*

```go
e.Use(middleware.Recover())
```

### 自定义配置

*使用*

```go
e := echo.New()
e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
  StackSize:  1 << 10, // 1 KB
}))
```

上面的示例使用 1 kb 的 `StackSize`，`DisableStackAll` 和 `DisablePrintStack` 使用默认值。

### 配置

```go
RecoverConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Size of the stack to be printed.
  // Optional. Default value 4KB.
  StackSize int `json:"stack_size"`

  // DisableStackAll disables formatting stack traces of all other goroutines
  // into buffer after the trace for the current goroutine.
  // Optional. Default value false.
  DisableStackAll bool `json:"disable_stack_all"`

  // DisablePrintStack disables printing stack trace.
  // Optional. Default value as false.
  DisablePrintStack bool `json:"disable_print_stack"`
}
```

*默认配置*

```go
DefaultRecoverConfig = RecoverConfig{
  Skipper:           defaultSkipper,
  StackSize:         4 << 10, // 4 KB
  DisableStackAll:   false,
  DisablePrintStack: false,
}
```

