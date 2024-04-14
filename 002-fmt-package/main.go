package main

import "fmt"

func main() {
  n, err := fmt.Println("Hello, playground", 42, true)
  fmt.Println(n)  // 26byte
  fmt.Println(err)  // error 0

  n1, _ := fmt.Println("안녕 해퓨쳐")
  fmt.Println(n1)
} 

// go 에서는 변수나 함수를 선언할 때에는 타입을 명시해야 한다.
// go 에서는 변수를 만들고 쓰지 않으면 에러가 난다.
// go 에서는 변수를 사용하지 않으려면 _ 를 붙여서 사용한다.