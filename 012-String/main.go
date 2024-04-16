package main

import "fmt"

func main() {
  s := "Hello, World!"
  fmt.Println("s: ", s)
  fmt.Printf("type s: %T\n", s)

  bs := []byte(s)
  fmt.Println(bs)
  fmt.Printf("type bs: %T\n", bs)

  for i := 0; i < len(bs); i++ {
    fmt.Printf("%#U", s[i])
  }
  fmt.Println("")

  for i, v := range s {
    fmt.Println(i, v)
  }
}