package main

import "fmt"

func getMyAge() (int, bool) {
	return 20, true
}

type ColorType int // ColorType을 선언하고 const 열거값 정의
const (
	Red ColorType = iota
	Yellow
	Blue
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "red"
	case Yellow:
		return "yellow"
	case Blue:
		return "blue"
	default:
		return "unknown"
	}
}

func main() {
	light := "red"

	if light == "green" {
		fmt.Println("길을 건넌다")
	} else if light == "red" {
		fmt.Println("기다린다")
	}

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}
	//fmt.Println(age) // 초기문에 선언한 변수의 범위는 if문 안으로 한정

	// Go에서는 case 하나를 실행 후 switch문을 빠져나감
	// break 문을 명시할 필요 없음
	switch age, ok := getMyAge(); ok {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	default:
		fmt.Println("My age is", age)
	}

	fmt.Println("My favorite color is", colorToString(Blue))

	// faalthrough 키워드는 다음 case까지 같이 실행
	switch a := 3; a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	default:
		fmt.Println("a >= 3")
	}

	//for i := 0; i < 10; i++ {
	//	fmt.Print(i, ", ")
	//}
	for i := range 10 {
		fmt.Print(i, ", ")
	}
	fmt.Println()

	x := 1
	y := 1

OuterFor:
	for ; x < 10; x++ {
		for y = 1; y < 10; y++ {
			if x*y == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", x, y, x*y)

	//for i := 5; i > 0; i-- {
	//	for j := i; j > 0; j-- {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
}
