package main

import "fmt"

// 첫 글자가 대문자인 함수는 패키지 외부로 공개
func Add(a int, b int) int {
	return a + b
} // 값 반환 후 로컬 변수 삭제

//// 함수는 값을 여러 개 반환할 수 있음
//func Divide(a, b int) (int, bool) {
//	if b == 0 {
//		return 0, false
//	}
//	return a / b, true
//}

// 반환할 변수에 이름을 지정할 경우 모든 반환 변수에 이름을 지정해야 함
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return // 명시적으로 반환할 값을 지정하지 않아도 됨
	}
	return a / b, true
}

func Multiply(a, b int) int {
	return a * b
}

// 재귀 호출
func printNo(n int) {
	// 재귀 호출을 사용할 때는 항상 탉출 조건을 정해야 함
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	// Go에서는 변수 전달이 항상 복사로만 이뤄짐
	x := Add(1, 2)
	fmt.Println(x)

	fmt.Println()

	a, success := Divide(9, 3)
	fmt.Println(a, success)
	b, success := Divide(9, 0)
	fmt.Println(b, success)

	fmt.Println()

	fmt.Println(Multiply(3, 5))

	fmt.Println()

	printNo(3)

	fmt.Println()

	fmt.Println(fibo(3))
}
