package main

import "fmt"

// 메서드란 특정 타입(struct 등)에 속해 동작을 정의하는 함수

type account struct {
	balance int
}

// 일반 함수
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메서드
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

// 사용자 정의 별칭 타입
type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

type student struct {
	age       int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a *student) addAgePointer(age int) {
	a.age += age
}

// 값 타입 메서드
func (a student) addAgeValue(age int) {
	a.age += age
}

// 변경된 값을 반환하는 값 타입 메서드
func (a student) returnAddAgeValue(age int) student {
	a.age += age
	return a
}

func main() {
	a := &account{10} //balance가 100인 account 포인터 변수 생성

	// 함수 형태 호출
	withdrawFunc(a, 30)

	// 메서드 형태 호출
	a.withdrawMethod(30)

	fmt.Println(a.balance)

	var i myInt = 10
	fmt.Println(i.add(20))
	fmt.Println()

	var studentA *student = &student{15, "Joe", "Park"}
	// 1) pointer receiver → 원본 변경
	studentA.addAgePointer(1)
	fmt.Println(studentA.age) // 16

	// 2) value receiver → 복사본 변경 → 원본은 그대로
	studentA.addAgeValue(1)
	fmt.Println(studentA.age) // 16

	var studentB student = studentA.returnAddAgeValue(1)
	fmt.Println(studentB.age) // 17

	studentB.addAgePointer(1)
	fmt.Println(studentB.age) // 18
}
