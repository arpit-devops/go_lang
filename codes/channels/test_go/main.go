package main

import (
	"fmt"
)

func main() {

	// for i := 1; i < 9; i++ {
	// 	fmt.Print(strconv.Itoa(fiboRec(i)) + " ")
	// }

	a, b := 0, 1

	for a < 100 {
		fmt.Println(a)
		b = a + b
		a = b - a
	}
}

func fiboRec(n int) int {

	if n <= 1 {
		return n
	}

	return fiboRec(n-1) + fiboRec(n-2)
}

//func fiboLoop()
