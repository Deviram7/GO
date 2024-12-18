package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{50, 90, 30, 10, 50}
	fmt.Println(num)

	fmt.Println("Ascending Order of Integer's")
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)
	fmt.Println()

	text := []string{"CMRIT", "CBIT", "CVSR", "CVRH", "IARE", "BVRIT", "MVSR"}
	fmt.Println(text)

	fmt.Println("String Reverse Sort")
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)
}