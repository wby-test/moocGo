package main

import (
	"fmt"
	"regexp"
)

const text = `my email is wby@qq.com
he email is wby@wangyi.com
she emial is wby@gmail.com.cn
`

func main() {
	//re, err := regexp.Compile(text)
	//if err != nil {
	//	panic(err)
	//}
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`) //.+一个或多个 .* 零个或多个,()提取内容
	//match := re.FindString(text)
	match := re.FindAllString(text, -1)
	match1 := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	fmt.Println(match1)
}
