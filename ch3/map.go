package main

import "fmt"

func main() {
	m := map[string]int {
		"first" : 1,
	}
	var m1 map[string]int
	var m2 = make(map[string]int)
	fmt.Println(m, m1,m2)
}
