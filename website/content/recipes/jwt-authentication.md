---
title: JWT Authentication
menu:
  side:
    parent: recipes
    weight: 11
---

- JWT authentication using HS256 algorithm.
- JWT is retrieved from `Authorization` request header.

### Server

`server.go`

{{< embed "jwt-authentication/server.go" >}}

### Client

`curl`

#### Login

Login using username and password to retrieve a token.

```sh
curl -X POST -d 'username=jon' -d 'password=shhh!' localhost:1323/login
```

*Response*

```js
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"
}
```

#### Request

Request a restricted resource using the token in `Authorization` request header.

```sh
curl localhost:1323/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"
```

*Response*

```
Welcome Jon Snow!
```

### Maintainers

- [vishr](https://github.com/vishr)
- [axdg](https://github.com/axdg)

### [Source Code]({{< source "jwt-authentication" >}})
