package main

import "fmt"

func main() {
  x := []int{1, 2, 3}
  y := []int{4, 5, 6}
  fmt.Println("x: ", x)
  fmt.Println("y: ", y)

  // append([]T, elements ...T)
  x = append(x, 7, 8, 9)
  fmt.Println("x:", x)

  // append([]T, slice ...T))
  x = append(x, y...)
  fmt.Println("x:", x)

  // remove element(7,8,9) from slice
  x = append(x[:3], x[6:]...)
  fmt.Println("x:", x)
}
