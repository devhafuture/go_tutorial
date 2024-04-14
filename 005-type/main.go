package main

import (
	"fmt"
)

var y int = 42
// DECLARED the VARIABLE with IEDNTIFIER "Z"
// is of TYPE string
// and I ASSIGN the value "James Bond"

var z string = "James Bond"
var a string = `James said "Shaken, not stirred"`

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYANMIC programming language

func main() {
  fmt.Println("y:", y)
  fmt.Printf("type: %T\n", y)
  fmt.Println("z:", z)
  fmt.Printf("type: %T\n", z)
  // z = 43 -> syntax error
  // fmt.Println("z:", z)
  // fmt.Printf("type: %T\n",  z)
  fmt.Println("a:", a)
  fmt.Printf("type: %T\n", a)
}