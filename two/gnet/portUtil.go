package gnet

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

//ChoosePort 自动为服务器选择未被占用的端口
func ChoosePort(host string, startPort int) int {
	myAddr := "127.0.0.1"
	host = strings.Replace(host, " ", "", -1)
	if len(host) > 0 || host != "" {
		myAddr = host
	}

	myPort := 9527
	if startPort > 1000 && startPort < 60000 {
		myPort = startPort
	}

	for i := myPort; i < 65535; i++ {
		if !CheckPort(myAddr, i) {
			fmt.Printf("端口%d已被占用", i)
			fmt.Println()
		} else {
			myPort = i
			break
		}
	}

	fmt.Printf("已选定的端口号：%d", myPort)
	fmt.Println()
	return myPort
}

// CheckPort 检测指定IP地址的端口是否被占用
func CheckPort(host string, port int) bool {
	timeout := time.Second
	isTimeout := false
	for i := 0; i < 2; i++ {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), 3*timeout)
		if err != nil {
			isTimeout = true //超时
		}
		if conn != nil {
			conn.Close() //OK
		}
	}
	return isTimeout
}
