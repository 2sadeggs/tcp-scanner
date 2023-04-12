package main

import (
	"fmt"
	"net"
)

/*https://www.infoq.cn/article/YfRpv2IpRNh2WXDF6PDv?source=app_share*/
//我们 TCP 扫描器第一步先实现单个端口的测试。使用标准库中的 net.Dial 函数，该函数接收两个参数：协议和测试地址（带端口号）。

func main() {
	//_, err := net.Dial("tcp", "www.baidu.com:80")
	_, err := net.Dial("tcp", "google.com:80")
	if err == nil {
		fmt.Println("Connection successful!")
	} else {
		fmt.Println(err)
	}
}
