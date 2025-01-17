package  main

import "fmt"

func main(){
	n := 3
	sum := 0;

	for i:=0; i<=n; i++{
		sum += i
	}

	fmt.Println("Sum of Numbers :", sum)
}