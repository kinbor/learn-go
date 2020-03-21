package gtest

import "fmt"

//PrintText 测试同一个module不同包的引用，以及不同module下不同包的引用
func PrintText() {
	fmt.Println("package:gaes ref package:gtest in the same module:gcrypto!")
}
