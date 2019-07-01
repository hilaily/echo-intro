+++
title = "Casbin 认证 "
url = "/middleware/casbin-auth"
[menu.side]
  name = "Casbin 认证"
  parent = "middleware"
  weight = 7
+++

[Casbin](https://github.com/casbin/casbin) 是 Go 下的强大而高效的开源访问控制库，它为基于各种模型的授权提供支持。到目前为止，Casbin 支持的访问控制模型如下：

- ACL (访问控制列表)
- 超级用户下的ACL
- 没有用户的 ACL： 对于没有身份验证或用户登录的系统尤其有用。
- 没有资源的ACL：过使用 write-article , read-log 等权限，某些方案可以应对一类资源而不是单个资源，它不控制对特定文章或日志的访问。
- RBAC (基于角色的访问控制)
- 具有资源角色的 RBAC： 用户和资源可以同时具有角色 (或组)。
- 具有域 / 租户 (tenants) 的 RBAC ：用户可以针对不同的域 / 租户 (tenants) 具有不同的角色集。
- ABAC (基于属性的访问控制)
- RESTful
- Deny-override：支持允许和拒绝授权，否认覆盖允许。

> Echo 社区贡献

## 依赖

```
import (
  "github.com/casbin/casbin"
  casbin_mw "github.com/labstack/echo-contrib/casbin"
)
```

*Usage*

```
e := echo.New()
e.Use(casbin_mw.Middleware(casbin.NewEnforcer("casbin_auth_model.conf", "casbin_auth_policy.csv")))
```

有关语法，请参阅：[Model.md](https://github.com/casbin/casbin/blob/master/Model.md)。

## 自定义配置

*用法*

```
e := echo.New()
ce := casbin.NewEnforcer("casbin_auth_model.conf", "")
ce.AddRoleForUser("alice", "admin")
ce.AddPolicy(...)
e.Use(casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
  Enforcer: ce,
}))
```

## 配置

```
// Config defines the config for CasbinAuth middleware.
Config struct {
  // Skipper defines a function to skip middleware.
  Skipper middleware.Skipper

  // Enforcer CasbinAuth main rule.
  // Required.
  Enforcer *casbin.Enforcer
}
```

*Default Configuration*

```
// DefaultConfig is the default CasbinAuth middleware config.
DefaultConfig = Config{
  Skipper: middleware.DefaultSkipper,
}
```

