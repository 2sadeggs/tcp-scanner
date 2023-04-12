package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

/*https://www.infoq.cn/article/YfRpv2IpRNh2WXDF6PDv?source=app_share*/
//我们 TCP 扫描器第一步先实现单个端口的测试。使用标准库中的 net.Dial 函数，该函数接收两个参数：协议和测试地址（带端口号）。

/*我们来看下如何实现并行。第一步先把扫描功能拆分为一个独立函数。这样会使我们的代码看起来清晰。*/
func isOpen(host string, port int) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err == nil {
		_ = conn.Close()
		return true
	} else {
		return false
	}
}

func main() {
	//声明一个端口切片 用于保存想要的测试端口结果 也就是打开的端口
	ports := []int{}

	wg := &sync.WaitGroup{} //声明一个wg 用于控制并发
	for port := 1; port < 100; port++ {
		wg.Add(1) //每个for循环开始 wg计数加1
		go func() {
			fmt.Println(port) //加个打印 查看port是否是并发安全
			//答案是否定 哈哈 有点意思
			opened := isOpen("google.com", port)
			if opened {
				ports = append(ports, port)
			}
			wg.Done() //每个匿名函数完成 wg计数减1
		}() //一下子开了99个并发匿名函数 每个匿名函数内会检测相应端口是否开放 并把开放端口追加到大切片内
		//不太确定port是否是并发安全 待验证
	}
	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
}

/*PS D:\GolandProjects\tcp-scanner> go run .\main.go
opened ports: [100]
*/
