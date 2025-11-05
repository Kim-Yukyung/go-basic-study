package main

import (
	"fmt"
	"unsafe"
)

// string 구조
type StringHeader struct {
	Data uintptr // 문자열의 데이터가 있는 메모리 주소
	Len  int     // 문자열의 길이
}

func main() {
	str1 := "Hello\t'World'\n"
	str2 := `Go is "awesome"!\nGo is simple and\tpowerful`
	fmt.Println(str1, str2)
	fmt.Println()

	str3 := `백쿼드에서는 여러 줄 표현에
특수 문자가 필요 없음`
	fmt.Println(str3)
	fmt.Println()

	var char rune = '가'
	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)
	fmt.Println()

	a := "가나다"
	b := "abc"
	fmt.Println(len(a), len(b)) // 문자열이 차지하는 메모리 크기
	fmt.Println()

	str := "Hello 월드"
	runes := []rune(str)
	fmt.Println(str)
	fmt.Println(string(runes))
	fmt.Println(len(runes))
	fmt.Println()

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", str[i], str[i], str[i])
	}
	fmt.Println()

	for i := 0; i < len(runes); i++ {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", runes[i], runes[i], runes[i])
	}
	fmt.Println()

	for _, v := range str {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", v, v, v)
	}
	fmt.Println()

	x := "I want go home!"
	y := x
	stringHeader1 := (*StringHeader)(unsafe.Pointer(&x))
	stringHeader2 := (*StringHeader)(unsafe.Pointer(&y))
	fmt.Println(stringHeader1, stringHeader2)
}
