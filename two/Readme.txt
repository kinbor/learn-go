1.检测指定IP地址的端口是否被占用
    a.函数：gnet.CheckPort(host string, port string) bool { ... }
    b.参数：
        b1.host主机名或IPv4，不许为空
        b2.port指定的端口号
    c.结果：若为TRUE，表示端口已开放；若为FALSE，表示端口未开放
2.自动选择一个未被占用的端口
    a.函数：gnet.ChoosePort(host string, startPort int) int { ... }
    b.参数：
        b1.host主机名或IPv4，允许为空
        b2.startPort指定的端口号[1001,60000]，默认9527
    c.结果：被选定的端口号