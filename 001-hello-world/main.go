package main

import (
	"fmt"
)

func main() {

  fmt.Println("Hello everyone, this is the most basic program in the world! and learning the Go programming language.")
  
  foo()
  
  for i := 0; i < 100; i++ {
    if i % 2 == 0 {
      fmt.Println(i, "is even")
    }
  }

  bar()
}

func foo() {
  fmt.Println("I'm in foo")
}

func bar() {
  fmt.Println("and then we exited")
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional