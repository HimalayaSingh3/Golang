package main


import "fmt"

func main(){

	var a int = 0
	var b int = 1
	var n int = 50

	for i:=0; i<=n; i++{
		fmt.Println(a)
		var next int = a + b
		a = b
		b = next
	}
}