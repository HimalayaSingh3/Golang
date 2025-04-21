package main

import "fmt"

func  main(){
	var fact int = 1
	var num = 5

	for i:= 1; i<= num ; i++ {
		fact = fact * i
	}

	fmt.Println("Factorial of ", num, " is ", fact)
}