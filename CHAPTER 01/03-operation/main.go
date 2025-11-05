package main

import (
	"fmt"
	//"math"
	"math/big"
)

//fucn equal(a, b float64) bool {
//	return math.Nextafter(a, b) == b
//}

func main() {
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t)

	fmt.Println()

	var x1 int8 = 34   // 00100010
	var x2 int16 = 34  // 00000000 00100010
	var x3 uint8 = 34  // 00100010
	var x4 uint16 = 34 // 00000000 00100010

	fmt.Printf("^%d = %5d, \t %08b\n", x1, ^x1, uint8(^x1))
	fmt.Printf("^%d = %5d, \t %016b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("^%d = %5d, \t %08b\n", x3, ^x3, ^x3)
	fmt.Printf("^%d = %5d, \t %016b\n", x4, ^x4, ^x4)

	fmt.Println()

	var a int8 = 4
	var b int8 = 64

	fmt.Printf("a: %08b a<<2: %08b a<<2: %d\n", a, a<<2, a<<2)
	fmt.Printf("b: %08b b<<2: %08b b<<2: %d\n", b, b<<2, b<<2)

	fmt.Println()

	var c int8 = 16
	var d int8 = -128
	var e int8 = -1
	var f uint8 = 128

	fmt.Printf("c: %08b c>>2: %08b c>>2: %d\n", c, c>>2, c>>2)
	fmt.Printf("d: %08b d>>2: %08b d>>2: %d\n", uint8(d), uint8(d>>2), d>>2)
	fmt.Printf("e: %08b e>>2: %08b e>>2: %d\n", uint8(e), uint8(e>>2), e>>2)
	fmt.Printf("f: %08b f>>2: %08b f>>2: %d\n", uint8(f), uint8(f>>2), f>>2)

	fmt.Println()

	a1, _ := new(big.Float).SetString("0.1")
	a2, _ := new(big.Float).SetString("0.2")
	a3, _ := new(big.Float).SetString("0.3")

	a4 := new(big.Float).Add(a1, a2)
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(a3.Cmp(a4))

	fmt.Println()

	var r int = 10
	var q int = 20

	r, q = q, r

	fmt.Println(r, q)
}
