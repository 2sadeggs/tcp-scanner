# tcp-scanner
简单TCP端口扫描器
思路总结：
0最初 接触net.Dial函数检测一个主机一个端口是否开放 （要有tcp三次握手的基础知识储备） 且发现了测试耗时长的问题
1接下来 增加了端口 添加for循环 可以检测同一个主机下的多个端口 现在端口是增加了 耗时问题没解决 且耗时更长了 因为每个端口检测都耗时 且是串行
2再接下来 既然串行耗时 那改成并行吧 于是单独抽出代码做了一个检测端口是否开放的函数 函数的初衷是并发改善串行耗时太长的问题 但是出现了新问题 那就是并发安全 
3再接下来 就是解决并发安全问题 利用for循环批量生成“令牌” 然后把“令牌”传给匿名函数检测 这样保证了每个匿名函数接受到的令牌不同 解决了上边for循环开始并发抢“循环初始变量”的情况
4再接下来 就是加入了超时机制 添加timeout参数 并且探测主力由net.Dial 改为了net.DialTimeout
5再再接下来 上边大幅度解决了耗时问题 但是主机和端口不灵活 每次换主机名都得重新编译 于是引入了flag包 将主机名、开始端口、结束端口、超时时间做成参数，
但是原来的例子参数值很长 输入命令行内明显不是很好记 所以改成了短参数 那么问题来了 flag支持长短参数一块识别吗
6补充：Windows主机环境的特殊性 在扫描任何主机的时候25端口都开放，这是假象！是假象！可能是本机火绒引起的
