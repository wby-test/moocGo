package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() funInterface {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}


type funInterface func() int

func (g funInterface) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)
	if next > 1000 {
		return 0,io.EOF
	}
	return strings.NewReader(s).Read(p)
}

func printFileContents(read io.Reader) {
	scanner := bufio.NewScanner(read)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	printFileContents(f)

}
