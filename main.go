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

/*执行时间很长 卡死在等待页面 需要处理超时*/
/*PS D:\GolandProjects\tcp-scanner> go run .\main.go
dial tcp 142.251.43.14:80: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or establ
ished connection failed because connected host has failed to respond.*/
