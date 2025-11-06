package main

import (
	"fmt"
	"sort"
)

// 슬라이스 내부 정의
// 포인터의 메모리 주소값이 복사
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {
	var slice []int // 길이가 0인 슬라이스

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}
	fmt.Println()

	//var array = [...]int{10, 20, 30, 40, 50} // 길이가 5인 고정 배열
	var slice1 = []int{10, 20, 30, 40, 50} // 슬라이스 선언
	var slice2 = []int{1, 5: 2, 10: 3}     // [1 0 0 0 0 2 0 0 0 0 3]
	fmt.Println(slice1, slice2)

	//// make(type, length)
	//var slice3 = make([]int, 3)

	for i := 0; i < len(slice1); i++ {
		slice1[i] += 10
	}

	for i, v := range slice1 {
		slice1[i] = v * 2
	}

	// 슬라이스 요소 추가
	// 요소가 추가된 새로운 슬라이스 반환
	slice1 = append(slice1, 1000)
	fmt.Println(slice1)
	slice1 = append(slice1, 2000, 3000, 4000, 5000)
	fmt.Println(slice1)
	fmt.Println()

	// append() 문제
	a := make([]int, 3, 5) // len: 3, cap: 5
	b := append(a, 4, 5)   // len: 5, cap: 5

	fmt.Println("a:", a, len(a), cap(a))
	fmt.Println("b:", b, len(b), cap(b))
	fmt.Println()

	a[1] = 100 // b까지 바뀜

	fmt.Println("a:", a, len(a), cap(a))
	fmt.Println("b:", b, len(b), cap(b))
	fmt.Println()

	a = append(a, 500)

	fmt.Println("a:", a, len(a), cap(a)) // len: 4, cap: 5
	fmt.Println("b:", b, len(b), cap(b)) // len: 5, cap: 5
	fmt.Println()

	// 슬라이싱하면 배열의 일부를 포인터로 가리키는 슬라이스를 만듦
	array := [5]int{1, 2, 3, 4, 5}
	slice = array[1:2]
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
	fmt.Println()

	//// 슬라이스를 슬라이싱
	//x := []int{1, 2, 3, 4, 5} // len: 5, cap: 5
	//y := x[1:2]               // len: 1, cap: 4
	//front := x[:3]            // 처음부터 슬라이싱
	//back := x[2:len(x)]       // 끝까지 슬라이싱 (== x[2:])
	//all := x[:]               // 전체 슬라이싱

	// 슬라이스 복제
	x := []int{1, 2, 3, 4, 5}
	//y := make([]int, len(x))
	//
	//for i, v := range x {
	//	y[i] = v
	//}

	//y := append([]int{}, x...)

	y := make([]int, len(x))
	copy(y, x)

	x[0] = 100
	fmt.Println(x, y)
	fmt.Println()

	// 요소 삭제
	idx := 0
	//for i := idx + 1; i < len(x); i++ {
	//	x[i-1] = x[i]
	//}
	//x = x[:len(x)-1] // 슬라이스 마지막 요소 삭제

	x = append(x[:idx], x[idx+1:]...)

	fmt.Println(x)
	fmt.Println()

	// 요소 추가
	//x = append(x, 0)
	//for i := len(x) - 1; i > idx; i-- {
	//	x[i] = x[i-1]
	//}
	//x[idx] = 1

	//x = append(x[:idx], append([]int{1}, x[idx:]...)...)

	// 메모리 낭비 줄이기
	x = append(x, 0)
	copy(x[idx+1:], x[idx:])
	x[idx] = 1

	fmt.Println(x)
	fmt.Println()

	// 슬라이스 정렬
	s := []int{5, 4, 3, 2, 1}
	sort.Ints(s)
	fmt.Println(s)
}
