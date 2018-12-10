//指针接收者只可以调用指针，值接收者指针和值都可以调用，跟函数相反，因为接口内置含有指针
//查看接口值和类型方法：intreface{}表示任何方法， type assertion, type switch
//接口变量自带指针，接口变量同样采用值传递， 几乎不需要使用接口的指针
//指针接收者实现只能以指针方式使用， 值接收者都可以用

package main

import (
	"fmt"
	"videoCourse/ch5/easyInterface/real"
)

type Retriever interface {
	Get(url string) string
}

func main() {
	var r Retriever
	r = real.Retriever{}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(r.Get("http://www.imooc.com"))

	//type assertion
	fmt.Println(r.(real.Retriever).Timeout)
}
