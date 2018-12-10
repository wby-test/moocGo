package main

import (
	"fmt"
	"unsafe"
)

var arr = [5]int{1, 2, 3, 4, 5}


func printArr(arr [5]int) {
	for _, v := range arr {
		v = v+1
	}
	fmt.Println(arr)
}

func printSlice(slice []int) {
	for index, v := range slice {
		slice[index] = v+1
	}
	fmt.Println(slice)
}

func main() {
	printArr(arr)
	slice := arr[1:4]
	fmt.Println("slice= ", slice)
	//printSlice(slice)
	fmt.Println("exchang:=",slice)
	fmt.Println("arr[1:4} " , arr[1:4])
	fmt.Println(arr)
	for i := 0; i<20; i++ {
		slice = append(slice, i+100)
	}

	fmt.Println("append--slice:=", slice)
	fmt.Println(arr)
	s := slice
	s[0] = 3
	fmt.Println("s" , s)
	fmt.Println("append--arr:=", arr)

	stringa := "q"
	stringb := 'q'
	fmt.Println(unsafe.Sizeof(string(stringa)), "   ",unsafe.Sizeof(stringb), unsafe.Sizeof(int8(2)))
	fmt.Println("arr[:-1]", arr[:])
}