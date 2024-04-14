package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println("a:", a)
	fmt.Printf("type a: %T\n", a)
	fmt.Println("b:", b)
	fmt.Printf("type b: %T\n", b)
	a = int(b)
	fmt.Println("a:", a)
	fmt.Printf("type a: %T\n", a)
}
