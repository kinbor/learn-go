特别说明：本文档中所有文字描述涉及的代码均在go version 1.4环境下执行。

此项目主要完成的主要目标有以下几个：
1.练习go mod依赖包管理与使用
1.1.module下对象关系
    a.a repository contains one or more modules(1:N)
    b.each module contains one or more packages(1:N)
    c.each package consists of one or more source file in a single directory(1:N)
1.2.包的存放位置：$GOPATH/pkg/mod，所以一定要注意磁盘空间大小
1.3.练习go mod相同module包的引用方式
1.4.练习go mod不同module包的引发方式，以及包与包的引用关系
1.4.1.引用方式：
    a.示例：关于gaes.go文件里注释掉的gtest引用，以及gtest.PrintText()方法的调用
    b.本地模式
        1.one/go.mod配置如下
            module testONE
            go 1.14
            require github.com/kinbor/learn-go/one/gcrypto v0.0.0
            replace github.com/kinbor/learn-go/one/gcrypto => ./gcrypto
        2.one/gcrypto/go.mod配置如下
            module gcrypto
            go 1.14
        3.one/gcrypto/gaes/gaes.go注释的引用改为“gcrypto/gtest”
        在完成上述三处修改后，运行gcrypto包里的main.go，gaes/gaes.go的引用不需要注释掉；若运行gcrypto包同一级的main.go需要注释掉，否则报错。
    c.域名模式
        在现有文件配置便是域名模式，直接运行即可。
1.4.2.版本号
    a.版本号格式
        v(major).(minor).(patch)=v(主版本号).(次版本号).(批次号)
    b.合法格式
        github.com/kinbor/learn-go/test.v1 v1.0.0-20141024135613-dd632973f1e7
        github.com/kinbor/learn-go/test.v2 v2.9.1
        github.com/kinbor/learn-go/test.v2 <=v2.2.1
        github.com/kinbor/learn-go/test v0.0.0-20160109021039-d7bb493dee3e   //版本号-日期-CommitId
        github.com/kinbor/learn-go/test latest
        github.com/kinbor/learn-go/test 分支名称
    c.版本升级
        go get -u 将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)，仅当前目录
        go get -u ./...将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)，当前目录及子目录，不包含测试依赖
        go get -u -t ./...将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)，当前目录及子目录，包含测试依赖
        go get -u=patch 将会升级到最新的修订版本
        go get package@version 将会升级到指定的版本号version
    d.查看依赖包版本更新信息
        go list -u -m all
1.5.常见问题
    1.5.1.权限问题，这主要指私有仓库，windows下的解决方法是：控制面板>用户账户>凭据管理：添加上凭证即可
    1.5.2.路径问题，常见的仓库域名可能并不在80或443端口下，这时可以搭建个中间代理，匹配上go get的正则表达式即可

2.学习go mod下的指令
2.1.go mod downlaod：下载依赖的module到cache目录
2.2.go mod graph：打印module依赖图
2.3.go mod init：在当前目录初始化一个新module
2.4.go mod tidy：增加丢失的module，删掉未使用的module
2.5.go mod vendor：从cache目录里复制代码引用的module库到vendor目录里
2.6.go mod verify：验证依赖module
2.7.go mod why：解释为什么需要依赖
2.8.go mod edit：通过命令行或终端工具设置go.mod

3.回忆go test的使用方法
3.1.运行当前目录及所有子目录下的测试用例：go test ./...
3.2.运行指定目录及所有子目录下的测试用例：go test foo/...
3.3.运行指定前缀的测试用例：go test foo...
3.4.运行GOPATH下的所有测试用例：go test ...

4.回忆go fmt的使用方法
4.1.go fmt相当于gofmt -l -w，前者是后者的封装，实际运行的是后者
4.2.go fmt参数
    a.-n参数，它会告诉go fmt把需要进行代码格式优化的文件打印出来，但是不会执行格式化
    b.-x参数，它会告诉go fmt去执行代码格式优化作业，完成格式化后将文件名称打印出来
4.3.使用方法
    a.go fmt -x会将当前目录下的*.go文件代码格式化
    b.go fmt -x ./...会将当前目录以及子目录（同一个module）下的*.go文件代码格式化。如果子目录属于另外一个module则不会对其*.go文件进行代码格式化
    c.go fmt xxxx.go会将当前指定的文件代码格式化

5.回忆go clean的使用方法
5.1.语法：go clean [clean flags] [build falgs] [packages]
5.2.参数：
    a.-i参数，主要清理由go install创建的安装包和可运行文件
    b.-n参数，主要把内部的清理过程打印出来，但并不执行清理作业
    c.-r参数，循环清理import中引用的包
    d.-x参数，主要把内部的清理过程打印出来，同时执行清理作业
    e.-cache参数，清理所有go buil产生的缓存
    f.-testcache参数，清理当前包所有的测试结果
    g.-modcache参数，清理module下载的所有缓存包，包含已经解压的源码包
5.3.常用方法
    a.go clean -i -r
    b.go clean -n或go clean -x清理编译文件
    c.go clean -cache清理缓存
    d.go clean -modcache清理module缓存

6.回忆go tool pprof性能分析
6.1.环境搭建
    a.安装Graphviz，官网：http://www.graphviz.org/
    b.配置环境变量，控制面板->系统->高级系统设置->环境变量，Path=安装包的安装bin目录
6.2.开发测试
    a.在main方法的文件中引用github.com/pkg/profile包，他人封装好的包
    b.在main方法的开始部位增加如下两行代码（根据传参可分析不同方面的性能）
        stopper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))  //开始性能分析, 返回一个停止接口
        defer stopper.Stop()                                                    //在main()结束时停止性能分析
    c.在main方法的结束部位增加如下一行代码
        time.Sleep(time.Second)                                                 //为了保证性能分析数据的合理性，分析的最短时间是 1 秒，使用 time.Sleep() 在程序结束前等待 1 秒。
6.3.执行测试
    a.go build -o cpu.exe main.go                       //将 main.go 编译为可执行文件 cpu。
    b../cpu.exe                                         //运行可执行文件，在当前目录输出 cpu.pprof 文件。
    c.结果输出到PDF文件（依赖Graphviz）
        go tool pprof -pdf cpu.exe cpu.pprof > cpu.pdf    //使用 go tool 工具链输入 cpu.pprof 和 cpu 可执行文件，生成 PDF 格式的输出文件，将输出文件重定向为 cpu.pdf 文件。
    d.结果输出到SVG文件（依赖Graphviz）
        go tool pprof -svg cpu.exe cpu.pprof > cpu.svg    //使用 go tool 工具链输入 cpu.pprof 和 cpu 可执行文件，生成 PDF 格式的输出文件，将输出文件重定向为 cpu.svg 文件。
    d.结果输出到PNG文件（依赖Graphviz）
        go tool pprof -png cpu.exe cpu.pprof > cpu.png    //使用 go tool 工具链输入 cpu.pprof 和 cpu 可执行文件，生成 PDF 格式的输出文件，将输出文件重定向为 cpu.png 文件。
    d.结果输出到TXT文件（或终端）
        go tool pprof -text cpu.exe cpu.pprof > cpu.txt   //使用 go tool 工具链输入 cpu.pprof 和 cpu 可执行文件，生成 txt 格式的输出文件，将输出文件重定向为 cpu.txt 文件。
        或
        go tool pprof -text cpu.exe cpu.pprof
6.4.Web环境
    a.如果你的go程序是用http包启动的web服务器，想要查看自己的web服务器的状态。这个时候就可以选择net/http/pprof。
    b.编码：import _ "net/http/pprof"。在浏览器中使用http://localhost:port/debug/pprof/ 直接看到当前web服务的状态，包括CPU占用情况和内存使用情况等。
    c.非Web程序也可以使用Web形式启动，编码如下：
        func main() {
            go func() {
                http.ListenAndServe("localhost:6060", nil)
            }()
        }

7.常用指令
    a.go build      编译包和依赖
    b.go install    编译并安装包和依赖
    c.go env        查看环境变量
    d.go list       查看包和模块
    e.go version    查看golang版本号
    f.go tool       工具链


8.golang开发环境搭建
8.1.下载最新的golang安装包：https://golang.google.cn/dl/
8.2.配置环境变量：控制面板->系统->高级系统设置->环境变量
    a.GO111MODULE=on开启go module
    b.GOPATH=xxxx设置gopath
    c.GOPROXY=https://goproxy.io或https://goproxy.cn配置golang代理
    d.Path=golang安装包的安装bin目录

9.练习aes算法
10.练习base64编解码