package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// 它包含了一个 scheme、 认证信息、主机名、端口、路径、查询参数以及查询片段。
	s := "postgres://user:pass@host.com:5432/testdb?name=yukki#segment"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// 这里我们提取路径和 # 后面的查询片段信息。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)

	query, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(query)
	fmt.Println(query["name"][0])
}
