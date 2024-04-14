package main

import (
	"fmt"
)

var y string
var z int

func main() {
	// DECLARE a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable

	fmt.Println("printing the value of y:", y, "ending")
	fmt.Printf("type y: %T\n", y)
	y = "Bond, James Bond"

	fmt.Println("printing the value of y:", y, "ending")
	fmt.Printf("type y: %T\n", y)

  fmt.Println("z:", z)
  fmt.Printf("type z: %T\n", z)
}
