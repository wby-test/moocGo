package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func (i int) {
			for {
				fmt.Println("hello  ",i )
			}
		}(i)
	}
	time.Sleep(time.Millisecond )
}
