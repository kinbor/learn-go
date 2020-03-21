package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/kinbor/learn-gomod/one/gcrypto/gaes"
	"github.com/pkg/profile"
)

func main() {
	// 开始性能分析, 返回一个停止接口
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	// 在main()结束时停止性能分析
	defer stopper.Stop()

	content := []byte("1234567890")
	key16 := []byte("1234567891234567")

	data, err := gaes.Encrypt(content, key16)
	if err != nil {
		fmt.Println(err)
	} else {
		tmpData := base64.StdEncoding.EncodeToString(data)
		fmt.Println(tmpData)
	}

	// 让程序至少运行1秒
	time.Sleep(time.Second)
}
