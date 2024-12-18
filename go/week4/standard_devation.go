package main

import (
	"fmt"
	"math"
)

func main() {
	var num [5]float64
	var sum, mean, sd float64

	fmt.Println("Enter 5 elements:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&num[i])
		sum += num[i]
	}

	mean = sum / 5

	for i := 0; i < 5; i++ {
		sd += math.Pow(num[i]-mean, 2)
	}
	sd = math.Sqrt(sd / 5)

	fmt.Printf("The Standard Deviation is: %.2f\n", sd)
}