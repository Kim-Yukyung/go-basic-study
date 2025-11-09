package main

import "fmt"

// 인터페이스는 메서드 규칙만 정의하고,
// 메서드를 구현한 타입이면 어떤 타입이든 해당 인터페이스로 사용할 수 있는 구조

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

// Student는 String() string 메서드를 가짐
// -> Student는 Stringer 인터페이스를 만족함
func (s Student) String() string {
	// string 타입 반환
	return fmt.Sprintf("Student{Name: %s, Age: %d}", s.Name, s.Age)
}

func main() {
	student := Student{"Mark", 20}
	var stringer Stringer // Stringer 인터페이스 선언

	// stringer는 Stringer 인터페이스, Student 타입은 String() 메서드 포함
	// -> stringer 값으로 student 대입 가능
	stringer = student

	fmt.Printf("%s\n", stringer.String())
}
