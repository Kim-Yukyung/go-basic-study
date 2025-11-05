package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 출력
	var a int = 10
	var b int = 20
	var f float64 = 3.14

	fmt.Print("a: ", a, "b: ", b)
	fmt.Println("a: ", a, "b: ", b, "f: ", f)
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f)

	// 입력
	var x int
	var y int

	n, err := fmt.Scan(&x, &y) // n은 성공적으로 입력한 값 개수
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, x, y)
	}

	n, err = fmt.Scanf("%d %d\n", &x, &y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, x, y)
	}

	n, err = fmt.Scanln(&x, &y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, x, y)
	}

	stdin := bufio.NewReader(os.Stdin) // 표준 입력을 읽는 객체

	n, err = fmt.Scanln(&x, &y)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, x, y)
	}

	n, err = fmt.Scanln(&x, &y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, x, y)
	}
}
