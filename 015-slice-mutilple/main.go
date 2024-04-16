package main

import "fmt"

func main() {

  // 1차원 slice
  jb := []string{"James", "Bond", "Chocolate", "martini"}
  fmt.Println(jb)
  mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
  fmt.Println(mp)

  // 2차원 slice
  xp := [][]string{jb, mp}
  fmt.Println(xp)
}
