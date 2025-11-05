package main

import "fmt"

const (
	Red   int = iota
	Green int = iota
	Blue  int = iota
)

const (
	C1 uint = iota + 1
	C2
	C3
)

// 타입 없는 상수
const PI = 3.14

func main() {
	const C int = 10

	//var b int = C * 20
	//C = 20 // 에러 발생: 상수는 대입문 좌변에 올 수 없음
	//fmt.Println(&C) // 에러 발생: 상수의 메모리 주소값을 접근할 수 없음
	// -> 컴파일 타임에 리터럴로 전환되어 실행 파일에 값 형태로 쓰이기 때문

	const PI1 float64 = 3.141592653
	var PI2 float64 = 3.141592653

	//PI1 = 4
	PI2 = 4

	fmt.Println("PI1: ", PI1)
	fmt.Println("PI2: ", PI2)

	fmt.Println()

	// 타입 없는 상수는 변수에 복사될 때 타입이 정해짐
	// 314는 정수 타입 변수에 대입할 수 있음
	var a int = PI * 100

	fmt.Println(a)
}
