+++
title = "Recover Middleware"
[menu.side]
  name = "Recover"
  parent = "middleware"
  weight = 5
+++

## Recover Middleware

Recover middleware recovers from panics anywhere in the chain, prints stack trace
and handles the control to the centralized
[HTTPErrorHandler]({{< ref "guide/customization.md#http-error-handler">}}).

### Configuration

```go
RecoverConfig struct {
  // StackSize is the stack size to be printed.
  // Optional, with default value as 4 KB.
  StackSize int

  // DisableStackAll disables formatting stack traces of all other goroutines
  // into buffer after the trace for the current goroutine.
  // Optional, with default value as false.
  DisableStackAll bool

  // DisablePrintStack disables printing stack trace.
  // Optional, with default value as false.
  DisablePrintStack bool
}
```

### Default Configuration

```go
DefaultRecoverConfig = RecoverConfig{
	StackSize:  4 << 10, // 4 KB
	StackAll:   true,
	PrintStack: true,
}
```

*Usage*

`e.Use(middleware.Recover())`

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
  StackSize:  1 << 10, // 1 KB
}))
```

Example above uses a `StackSize` of 1 KB and default values for `DisableStackAll`
and `DisablePrintStack`.
