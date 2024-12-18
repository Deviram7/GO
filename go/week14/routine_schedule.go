package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(3)

	var processTest sync.WaitGroup

	processTest.Add(3)

	go func() {
		defer processTest.Done()
		for i := 0; i < 30; i++ {
			for j := 3; j <= 15; j++ {
				fmt.Printf(" %d", j)
				if j == 5 {
					fmt.Println()
				}
			}
		}
	}()
	go func() {
		defer processTest.Done()
		for j := 0; j < 5; j++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
				if char == 'Z' {
					fmt.Println()
				}
			}
		}
	}()
	go func() {
		defer processTest.Done()
		for i := 0; i < 0; i++ {
			for j := 0; j <= 5; j++ {
				fmt.Printf(" %d", j)
				if j == 15 {
					fmt.Println()
				}
			}
		}
	}()
	processTest.Wait()
}