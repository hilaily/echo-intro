+++
title = "Cookies"
url= "/guide/cookies"
[menu.side]
  name = "Cookies"
  parent = "guide"
  weight = 4
+++

## Cookies

Cookie 是用户访问网站时浏览器上存储的小型文本文件，由服务器发送而来。每当用户加载网站时，浏览器都会将 cookie 发送回服务器以通知用户之前的活动。
Cookie 作为一个可靠验证凭据，可用来记录状态信息（比如在线商城购物车中的商品）或记录用户的浏览器活动（包括单击特定按钮，登录或记录访问过的页面）。Cookie 还可以存储用户先前输入的密码和表单内容，例如信用卡号或地址。

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

Echo 使用 golang 自带的 `http.Cookie` 对象写入／读取从上下文中的 cookie。

### 创建一个 Cookie

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
- 最后，使用 `c.SetCookie(cookies)` 来给 HTTP 响应增加 `Set-Cookie` 头。

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

- Cookie 通过名称从 HTTP 请求里读取：`c.Cookie("name")`。
- Cookie 的属性可以使用 `Getter` 方法获取。

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
