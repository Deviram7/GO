package main

import (
	"fmt"
)

func lcm(t1, t2 int) int {
	greater := t1
	if t2 > t1 {
		greater = t2
	}
	for greater%t1 != 0 || greater%t2 != 0 {
		greater++
	}
	return greater
}

func gcd(t1, t2 int) int {
	for t2 != 0 {
		t1, t2 = t2, t1%t2
	}
	return t1
}

func main() {
	var n1, n2, action int
	fmt.Println("Enter two positive integers:")
	fmt.Scanln(&n1, &n2)
	fmt.Println("Enter 1 for LCM or 2 for GCD:")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Printf("LCM of %d and %d is %d\n", n1, n2, lcm(n1, n2))
	case 2:
		fmt.Printf("GCD of %d and %d is %d\n", n1, n2, gcd(n1, n2))
	default:
		fmt.Println("Invalid choice")
	}
}