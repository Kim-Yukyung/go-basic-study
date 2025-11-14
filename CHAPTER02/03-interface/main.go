package main

import "fmt"

// 인터페이스는 메서드 규칙만 정의하고,
// 메서드를 구현한 타입이면 어떤 타입이든 해당 인터페이스로 사용할 수 있는 구조
// -> 덕 타이핑 (명시적으로 "implements" 선언이 필요 없음)

// 인터페이스 변수의 기본값은 nil
// -> 인터페이스를 사용할 때 항상 nil이 아닌지 확인

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

// 포함된 인터페이스
type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader // Reader의 메서드 집합을 포함
	Writer // Writer의 메서드 집합을 포함
}

// 빈 인터페이스 -> 모든 타입을 받을 수 있음
func PrintValue(s interface{}) {
	switch v := s.(type) {
	case int:
		fmt.Printf("Integer %d\n", v)
	case float64:
		fmt.Printf("Float %f\n", v)
	case string:
		fmt.Printf("String %s\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	student := Student{"Mark", 20}
	var stringer Stringer // Stringer 인터페이스 선언

	// stringer는 Stringer 인터페이스, Student 타입은 String() 메서드 포함
	// -> stringer 값으로 student 대입 가능
	stringer = student

	fmt.Printf("%s\n", stringer.String())
	fmt.Println()

	PrintValue(10)
	PrintValue(3.14)
}
