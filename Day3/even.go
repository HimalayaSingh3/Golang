package main

import "fmt"

func main() {
	var N int = 6

	if N%2 == 0 {
		fmt.Println("Even Number")
	}
	if N%2 != 0 {
		fmt.Println("Odd Number")
	}

}
