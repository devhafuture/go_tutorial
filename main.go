package main

import "fmt"

func main() {

	// map: 키에 기반해 값을 저장하는 방식으로 효율적이고 빠르다.
	// map은 key와 value를 쌍으로 묶어서 저장한다.
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)             // map[James:32 Miss Moneypenny:27]
	fmt.Println(m["James"])    // 32
	fmt.Println(m["Barnabas"]) // 0

	// comma ok 관용구 : 요소가 들어 있는지 확인할 때 사용
	v, ok := m["Barnabas"]
	fmt.Println(v)  // 0
	fmt.Println(ok) // false

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("This is the if print", v)
	}

	// map에 요소 추가하기
	m["Todd"] = 33
	fmt.Println(m) // map[James:32 Miss Moneypenny:27 Todd:33]

	for i, v := range m {
		fmt.Println(i, v)
	}

	// map에서 요소 삭제하기 -> delete(map, key)
	delete(m, "James")
	fmt.Println(m) // map[Miss Moneypenny:27 Todd:33]

	delete(m, "Miss Moneypenny")
	fmt.Println(m) // map[Todd:33]

  delete(m, "Iam Fleming00")
  fmt.Println(m) // map[Todd:33]  없는 키를 제거해도 에러가 나지 않는다.

  if v, ok := m["Todd"]; ok {
    fmt.Println("value:", v)
    delete(m, "Todd")
  }
  fmt.Println(m) // map[]
}