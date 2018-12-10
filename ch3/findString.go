package main

import "fmt"

func findShortestString(src string) (des map[int]string) {
	des1 := map[int]string{}
	for i := 0; i < len(src)-1; i++ {
		flag := src[i]
		for j := i+1; j < len(src); j++ {
			if flag == src[j] {
				des1[j-i] = src[i:j]
				break
			}
		}
	}
	return des1
}

func main() {
	a := "dfdfadff"
	fmt.Printf("%s ", a[2:6])
	des := map[int]string{}
	des = findShortestString(a)
	fmt.Println(des)
}
