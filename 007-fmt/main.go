package main

import (
  "fmt"
)

var y int = 42

func main() {
  fmt.Println("y:", y)
  fmt.Printf("type y: %T\n", y)
  fmt.Printf("%b\n", y)
  fmt.Printf("%x\n", y)
  fmt.Printf("%#x\n", y)
  y = 911
  fmt.Println("y:", y)
  fmt.Printf("%#x\t%b\t%x\n", y, y, y)

  s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
  fmt.Println(s)
  fmt.Printf("%v", y)
}