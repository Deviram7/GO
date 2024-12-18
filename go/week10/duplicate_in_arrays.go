package main

import (
	"fmt"
)

func main() {
	myarray := []string{"CMRIT", "cmrit", "CmRit", "cmrIT", "CMRit", "cmrit"}
	fmt.Println("Original Array is: ", myarray)
	uniqueArray := removeDuplicate(myarray)
	fmt.Println("After removing duplicate: ", uniqueArray)
}

func removeDuplicate(myarray []string) []string {
	visit := make(map[string]bool)
	result := []string{}
	for _, value := range myarray {
		if !visit[value] {
			visit[value] = true
			result = append(result, value)
		}
	}
	return result
}
