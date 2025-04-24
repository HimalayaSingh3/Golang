package main

import "fmt"

func main(){

	var arr[10] int
	var i int

	for i=0; i<10; i++{
		fmt.Println("Enter No",i)
		fmt.Scanln(&arr[i])
	}

	fmt.Println(arr)
}