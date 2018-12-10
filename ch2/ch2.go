package main

import (
	"fmt"
	"io/ioutil"
)

func printFile(name string) {
	value, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("readfile error")
	} else {
		fmt.Printf("%s/n", value)
	}
}

func main()  {
	value := "//GoProject//src//videoCourse//ch2//ch2.go"
	printFile(value)
}