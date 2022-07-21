package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Please enter your name.")
	var name string
	scanner(&name)
	fmt.Printf("Hi, %s! I'm Go!", name)
}
