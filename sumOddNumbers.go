package  main

import "fmt"

func main(){
	n := 5
	sum := 0;

	for i:=0; i<=n; i++{
		if i%2 != 0{
			sum += i
		}
	}

	fmt.Println("Sum of Odd Numbers :", sum)
}