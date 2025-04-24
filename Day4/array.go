package main

import "fmt"

func main(){

	arr := [5]int{2,7,1,9,4}


	for i:=0; i<len(arr); i++{
		if arr[i] > arr[0]{
		arr[0] = arr[i]
		}
	}

	fmt.Println(arr[0])
}