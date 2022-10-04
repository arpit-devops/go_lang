package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
		fmt.Printf("Input was: %q\n", name)
	}

	fmt.Printf("name is  %q", name)

}
