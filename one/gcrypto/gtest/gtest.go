package gtest

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//GenRandomString 测试同一个module不同包的引用，以及不同module下不同包的引用
func GenRandomString(strLength int) string {
	if strLength < 1 {
		strLength = 1
	}
	if strLength > 256 {
		strLength = 256
	}

	var build strings.Builder        //字符串拼劲使用它效率较高
	rand.Seed(time.Now().UnixNano()) //for循环内的随机数较随机，如果在外面调用两次GenRandomString方法可能得到相同的随机字符串，此时应该使用crypto/rand包里的随机函数
	for build.Len() < strLength {
		x := rand.Intn(123)
		//0-9
		if x > 47 && x < 58 {
			build.WriteString(strconv.Itoa(x))
		}
		//a-z
		if x > 96 && x < 123 {
			build.WriteString(string(x))
		}
	}
	fmt.Println("Random String：" + build.String())
	return build.String()
}
