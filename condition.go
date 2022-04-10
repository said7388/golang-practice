package main

import "fmt"


func main() {
	var marks int
	fmt.Printf("Enter your marks :\t")
	fmt.Scan(&marks)

	if marks <= 100 && marks >= 80 {
		fmt.Println("You have got GPA 5")
	} else if marks <= 79 && marks >= 70 {
		fmt.Println("You have got GPA 4")
	} else if marks <= 69 && marks >= 60 {
		fmt.Println("You have got GPA 3.5")
	} else if marks <= 59 && marks >= 50 {
		fmt.Println("You have got GPA 3")
	} else if marks <= 49 && marks >= 40 {
		fmt.Println("You have got GPA 2")
	} else if marks <= 39 && marks >= 33 {
		fmt.Println("You have got GPA 1")
	} else {
		fmt.Println("You are feiled in this subject")
	}
}
