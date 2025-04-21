package main

import "fmt"

func main() {
	var n int = 100

	for i := 2; i < n; i++ {
		var isPrime bool = true
		for j := 2; j*j <= i; j++ {
			if i % j == 0{
				isPrime = false
				break
			}
		}
		if isPrime{
			fmt.Println(i)
		}
	}

}