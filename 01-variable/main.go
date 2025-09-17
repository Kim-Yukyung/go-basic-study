package main

import "fmt"

var globalVariable string = "global variable"

func main() {
	fmt.Println(globalVariable)

	// 변수
	var var1 int = 1
	var var2 int
	var var3 = 3
	var4 := 5

	fmt.Println(var1, var2, var3, var4)

	// 타입 변환
	a := 3
	var b float64 = 3.5

	var c int = int(b) // 소수점은 버려짐
	result := float64(a * c)
	fmt.Println(result)

	// 실수
	var d float32 = 1234.523
	var e float32 = 3456.123
	var f float32 = d * e
	var g float32 = f * 3

	fmt.Println(d, e, f, g)
}
