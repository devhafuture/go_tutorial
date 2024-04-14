package main

import "fmt"

// x := 43 -> syntax error: non-declaration statement outside function body.
// DECALRE the variable "y"
// ASSIGN the value 43
// declare & assign = initilization
var y int = 43 // -> type 생략 가능

// DECLARE there is a VARIABLE with whte IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE "int"
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for interger, 0.0 fo r floats, "" for strings, nil for pointers, functions, interfaes, slices, channels, and maps.
var z int

func main() {
  // short declaration operator -> only use in function
  // DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
  x := 42
  fmt.Println("x:", x)

  fmt.Println("y:", y)

  fmt.Println("z:", z)

  foo()
}

func foo() {
  fmt.Println("again y:", y)
}

// the scope of y would be package level