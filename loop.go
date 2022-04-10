package main

import "fmt"

func main() {

	// Normal for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// While loop using for loop
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}
	
}
