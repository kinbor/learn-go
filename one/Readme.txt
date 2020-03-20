此项目主要完成的主要目标有以下几个：
1.练习go mod依赖包管理的使用
2.练习go mod同一个module不同包的引用方式
3.练习go mod不同module包的引发方式，以及包与包的引用关系
3.1.注意事项：gaes.go文件里注释掉的gtest引用，以及gtest.PrintText()方法的调用，运行gcrypto包里的main.go不需要注释掉；若运行gcrypto包同一级的main.go需要注释掉，否则报错。
4.回忆go test的使用方法
5.练习aes算法
6.练习base64编解码
