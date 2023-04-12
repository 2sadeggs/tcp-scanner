package main

/*注意：windows主机有可能telnet任何主机的25端口都通 不过这都是幻觉 幻觉 实际不同
有可能是火绒等一些杀毒软件造成的*/
import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

/*https://www.infoq.cn/article/YfRpv2IpRNh2WXDF6PDv?source=app_share*/
//我们 TCP 扫描器第一步先实现单个端口的测试。使用标准库中的 net.Dial 函数，该函数接收两个参数：协议和测试地址（带端口号）。

/*我们来看下如何实现并行。第一步先把扫描功能拆分为一个独立函数。这样会使我们的代码看起来清晰。*/
func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	//conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)

	if err == nil {
		_ = conn.Close()
		return true
	} else {
		return false
	}
}

func main() {
	/*flag 包可以帮助我们编写命令行程序。我们可以配置每个字符串或数字。我们为主机名及要测试的端口范围和连接超时添加参数。*/
	hostname := flag.String("h", "", "hostname to test")
	startPort := flag.Int("s", 80, "the port on which the scanning starts")
	endPort := flag.Int("e", 100, "the port from which the scanning ends")
	timeout := flag.Duration("t", time.Millisecond*200, "timeout")
	flag.Parse()

	//声明一个端口切片 用于保存想要的测试端口结果 也就是打开的端口
	ports := []int{}

	wg := &sync.WaitGroup{} //声明一个wg 用于控制并发
	/*唯一的问题就是，现在这个程序会有竞争条件。
	  在只扫描少数端口时，速度比较慢，可能不会出现，
	  但确实存在这个问题。所以我们需要使用 mutex 来修复它。*/
	mutex := &sync.Mutex{}
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1) //每个for循环开始 wg计数加1
		go func(p int) {
			//fmt.Println(p) //加个打印 查看port是否是并发安全
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done() //每个匿名函数完成 wg计数减1
		}(port)
		//改变匿名函数的调用方式 改用传参 经验证解决了并发安全问题
		/*解决问题的关键思路 是把for循环生成以后的值依次传递给匿名函数 这很重要 至少能保证每次传给匿名函数的值不同
		如果按上一次的思路 在for循环开始时 99个go并发已经出去 这样99个并发很用可能取到相同的值 这不是预期的运行方式
		预期的运行方式是遍历for循环的每一个值 是每一个值 且每一个值本身也不相同*/
	}
	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
}

/*
测试结果
PS D:\GolandProjects\tcp-scanner> .\main.exe -hostname www.baidu.com -startPort 1 -endPort 500
opened ports: [110 25]
*/
