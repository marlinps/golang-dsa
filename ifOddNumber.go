package main

import "fmt"

func main(){
	number := 20

	if(number%2 == 0){
		fmt.Println("Even Numbers", number)
	}else{
		fmt.Println("Odd Numbers", number)
	}
}