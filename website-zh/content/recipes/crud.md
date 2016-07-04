---
title: CRUD
menu:
  side:
    parent: recipes
    weight: 2
---

## CRUD 示例

### 服务端

`server.go`

{{< embed "crud/server.go" >}}

### 客户端

`curl`

#### 创建 User

```sh
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe Smith"}' \
  localhost:1323/users
```

*Response*

```js
{
  "id": 1,
  "name": "Joe Smith"
}
```

#### 获取 User

```sh
curl localhost:1323/users/1
```

*Response*

```js
{
  "id": 1,
  "name": "Joe Smith"
}
```

#### 更新 User

```sh
curl -X PUT \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe"}' \
  localhost:1323/users/1
```

*Response*

```js
{
  "id": 1,
  "name": "Joe"
}
```

#### 删除 User

```sh
curl -X DELETE localhost:1323/users/1
```

*Response*

`NoContent - 204`

### 维护者

- [vishr](https://github.com/vishr)

### [Source Code]({{< source "crud" >}})
