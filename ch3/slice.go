package main

import "fmt"

//delete element of slice
func deleteSliceElement(slice []int, index int) []int{
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[3:6]
	s2 := s1[2:7]
	fmt.Println("s1 ", s1)
	fmt.Println("s2 ", s2)

	s := []int{1, 2, 3}
	s[1] = 100
	s3 := make([]int, 10, 100)
	fmt.Println(s3, len(s3), cap(s1))

	//delete
	delet := arr[1:7]
	fmt.Println("begin delete: ", delet)
	delet = deleteSliceElement(delet, 5)
	fmt.Println("after delete: ", delet)
}
