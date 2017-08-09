+++
title = "Cookies"
url= "/guide/cookies"
[menu.side]
  name = "Cookies"
  parent = "guide"
  weight = 8
+++

## Cookies

Cookie 是用户在访问网站时服务器发送过来存储在浏览器上的一小段数据。每次用户访问网页，浏览器都把 Cookies 发送回服务器以提醒服务器这个用户以前干过什么。
Cookie 用来提供一个可靠的途径让服务器记住一些状态信息（比如在线商城中添加物品到购物车）或者记录用户的浏览器行为（比如点击了某个按钮，登录，哪个页面被访问过）。
Cookie 也可以用来存储用户输入过的表单内容像电话号码，地址等等。

### Cookie 属性

| 属性       | 可选   |
| :------- | :--- |
| Name     | No   |
| Value    | No   |
| Path     | Yes  |
| Domain   | Yes  |
| Expires  | Yes  |
| Secure   | Yes  |
| HTTPOnly | Yes  |

### 创建一个 Cookie

Echo 使用 golang 自带的 `http.Cookie` 对象来从处理函数的上下文里写入／读取 cookies。
```go
func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}
```

- 使用 `new(http.Cookie)` 创建Cookie。
- cookie 的属性值会被赋值给 `http.Cookie` 的可导出属性。
- 最后，使用 `c.SetCookie(cookies)` 来给响应添加 `Set-Cookie` 头。

### 读取 Cookie

```go
func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}
```

- Cookie 通过名称从 HTTP 请求里读取: `c.Cookie("name")`。
- Cookie 的属性可以使用`Getter` 方法获取。

### 读取所有 Cookies

```go
func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all cookie")
}
```
