package main

import (
	parent "family/father"
	child "family/father/son"
	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("Howard Stark"))
	c := new(child.Son)
	fmt.Println(c.Data("Tony Stark"))
}