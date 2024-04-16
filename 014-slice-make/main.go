package main

import "fmt"

func main() {

  // slice 만들고 용량 늘리는 법
	x := make([]int, 10, 12)
	fmt.Println("x:", x)
	fmt.Println("x length:", len(x))
	fmt.Println("x capacity:", cap(x))
  
  x[0] = 42
  x[9] = 999
	fmt.Println("x:", x)
	fmt.Println("x length:", len(x))     // 10
	fmt.Println("x capacity:", cap(x))   // 12

  x = append(x, 3423)
	fmt.Println("x:", x)
	fmt.Println("x length:", len(x))     // 11
	fmt.Println("x capacity:", cap(x))   // 12
  
  x = append(x, 3424)
	fmt.Println("x:", x)
	fmt.Println("x length:", len(x))     // 12
	fmt.Println("x capacity:", cap(x))   // 12

  // capacity 용량을 초과하면 기본 배열을 버리고 capacity 용량은 두배로 늘어난다.
  x = append(x, 3425)
	fmt.Println("x:", x)
	fmt.Println("x length:", len(x))     // 13
	fmt.Println("x capacity:", cap(x))   // 24
}
