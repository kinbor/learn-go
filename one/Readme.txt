此项目主要完成的主要目标有以下几个：
1.练习go mod依赖包管理的使用
2.练习go mod同一个module不同包的引用方式
3.练习go mod不同module包的引发方式，以及包与包的引用关系
3.1.引用方式：
    a.示例：关于gaes.go文件里注释掉的gtest引用，以及gtest.PrintText()方法的调用
    b.本地模式
        1.one/go.mod配置如下
            module testONE
            go 1.14
            require github.com/kinbor/learn-gomod/one/gcrypto v0.0.0
            replace github.com/kinbor/learn-gomod/one/gcrypto => ./gcrypto
        2.one/gcrypto/go.mod配置如下
            module gcrypto
            go 1.14
        3.one/gcrypto/gaes/gaes.go注释的引用改为“gcrypto/gtest”
        在完成上述三处修改后，运行gcrypto包里的main.go，gaes/gaes.go的引用不需要注释掉；若运行gcrypto包同一级的main.go需要注释掉，否则报错。
    c.域名模式
        在现有文件配置下，直接运行即可。
3.2.版本号
    a.合法格式
        github.com/kinbor/learn-gomod/test.v1 v1.0.0-20141024135613-dd632973f1e7
        github.com/kinbor/learn-gomod/test.v2 v2.9.1
        github.com/kinbor/learn-gomod/test.v2 <=v2.2.1
        github.com/kinbor/learn-gomod/test v0.0.0-20160109021039-d7bb493dee3e
        github.com/kinbor/learn-gomod/test latest
    b.版本升级
        go get -u 将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
        go get -u=patch 将会升级到最新的修订版本
        go get package@version 将会升级到指定的版本号version
3.3.vendor
    go mod vendor 会复制modules下载到vendor中, 只会下载你代码中引用的库，而不是go.mod中定义全部的module。

4.回忆go test的使用方法
5.练习aes算法
6.练习base64编解码