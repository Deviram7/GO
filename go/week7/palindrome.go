package main

import "fmt"

func main() {
	fmt.Println("Golang program to check for palindrome")
	var number, rem, temporary int
	var reverse int = 0
	number = 5454
	temporary = number
	for {
		rem = number % 10
		reverse = reverse*10 + rem
		number /= 10
		if number == 0 {
			break
		}
	}
	if temporary == reverse {
		fmt.Printf("Number %d is a Palindrome", temporary)
	} else {
		fmt.Printf("Number %d is not a Palindrome", temporary)
	}
}
