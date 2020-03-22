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
    d.注意事项
        1.不同版本的导入路径
            import (
                "github.com/eddycjy/mquote/v2/example"
            )
            如上示例，导入的路径里多了一个"v2"。Go modules 在主版本号为 v0 和 v1 的情况下省略了版本号，而在主版本号为 v2 及以上则需要明确指定出主版本号，否则会出现冲突。
        2.版本号哪里来的
            平时使用github/gitlab/gitee等版本管理工具的时候，不太在意这个东西，其实就是Tag标签，具体如何操作自行搜索。

1.4.2.版本号
    a.版本号格式
        v(major).(minor).(patch)=v(主版本号).(次版本号).(修订号)
        1.主版本号：当你做了不兼容的 API 修改
        2.次版本号：当你做了向下兼容的功能性新增
        3.修订号：当你做了向下兼容的问题修正
        4.版本号延伸格式
            v(major).(minor).(patch)-xxxx

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

7.回忆go build使用方法
    go build是非常常用的命令，它可以启动编译，把我们的包和相关的依赖编译成一个可执行的文件。
7.1.语法
        go build [-o output] [-i] [build flags] [packages]
    其中所有的参数都可以忽略，直接运行go build 命令，就会把当前目录里的文件编译到当前目录下。
7.2.跨平台编译
    a.使用go build更多的会进行跨平台编译，Go提供了编译链工具，可以让我们在任何一个开发平台上，编译出其他平台的可执行文件。默认情况下，都是根据我们当前的机器生成的可执行文件。
    b.跨平台编译主要涉及到两个环境变量参数的设置：GOOS 和 GOARCH，这两个参数的取值如下表：来自https://golang.org/doc/install/source
        GOOS        GOARCH                      GOOS        GOARCH
        aix         ppc64                       linux	    ppc64le
        android     386                         linux	    mips
        android     amd64                       linux	    mipsle
        android	    arm                         linux	    mips64
        android	    arm64                       linux	    mips64le
        darwin	    386                         linux	    s390x
        darwin	    amd64                       netbsd	    386
        darwin	    arm                         netbsd	    amd64
        darwin	    arm64                       netbsd	    arm
        dragonfly	amd64                       openbsd	    386
        freebsd	    386                         openbsd	    amd64
        freebsd	    amd64                       openbsd	    arm
        freebsd	    arm                         openbsd	    arm64
        illumos	    amd64                       plan9	    386
        js	        wasm                        plan9	    amd64
        linux	    386                         plan9	    arm
        linux	    amd64                       solaris	    amd64
        linux	    arm                         windows	    386
        linux	    arm64                       windows	    amd64
        linux	    ppc64
    c.使用方法
        以32位windows系统为例：GOOS=windows GOARCH=386 go build -o myCompiler_x86.exe main.go
        以64位windows系统为例：GOOS=windows GOARCH=amd64 go build -o myCompiler_x64.exe main.go

8.回忆go install使用方法
    go install指令和go build类似，且大部分参数也通用。它只是将编译的中间文件放在GOPATH/pkg 目录下，以及固定地将编译结果放在GOPATH/bin目录下。
8.1.语法
        go install [-i] [build flags] [packages]
    它和go build参数比较仅少了 -o 参数。
8.2.用法
    按照go build语法使用即可。

9.其他指令
    a.go env        查看环境变量
    b.go list       查看包和模块
    c.go version    查看golang版本号
    d.go tool       工具链


10.golang开发环境搭建
10.1.下载最新的golang安装包：https://golang.google.cn/dl/
10.2.配置环境变量：控制面板->系统->高级系统设置->环境变量
    a.GO111MODULE=on开启go module
    b.GOPATH=xxxx设置gopath
    c.GOPROXY=https://goproxy.io或https://goproxy.cn配置golang代理
    d.Path=golang安装包的安装bin目录
10.3.VSCode环境配置
    a.安装Go for Visual Studio Code插件
    b.VSCode会提示安装golang开发调试等相关的程序，按照提示安装即可。大致包含以下程序：
        01.dlv.exe
        02.fillstruct.exe
        03.go-outline.exe
        04.go-symbols.exe
        05.gocode-gomod.exe
        06.gocode.exe
        07.godef.exe
        08.godoctor.exe
        09.golint.exe
        10.gomodifytags.exe
        11.gopkgs.exe
        12.goplay.exe
        13.gorename.exe
        14.goreturns.exe
        15.gotests.exe
        16.guru.exe
        17.impl.exe
    c.调试程序：在VSCode里调试程序与VS中调试相差不大。首先在代码行的开头部分点击加断点，F5运行调试，调试依赖第二步里的相关插件。

11.练习aes算法
12.练习base64编解码
13.对外暴露的方法应当增加注释使用说明，具体格式如下
    // Add result=num1 +num2
    // format:MethodName+WhiteSpace+...
    func Add(num1 int, num2 int) result int{
        return num1 + num2
    }
