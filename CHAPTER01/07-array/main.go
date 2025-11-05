package main

import "fmt"

func main() {
	//var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	t := [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	//for i := 0; i < 5; i++ {
	//	fmt.Println(t[i])
	//}

	// range는 인덱스와 요솟값 반환
	for _, v := range t {
		fmt.Println(v)
	}

	// 인덱스 1에는 10, 인덱스 3에는 30
	var s = [5]int{1: 10, 3: 30}
	fmt.Println(s)

	// 배열 요소 개수는 초기값에 따라 결정
	x := [...]int{10, 20, 30}
	fmt.Println(x)

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{10, 20, 30, 40, 50}
	b = a // 배열 복사
	fmt.Println(b)

	var q = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	for _, i := range q {
		for _, v := range i {
			fmt.Print(v, " ")
		}
	}
}
