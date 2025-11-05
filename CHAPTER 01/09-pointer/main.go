package main

import "fmt"

type Data struct {
	value int
	data  [100]int
}

func ChangeData(arg *Data) {
	arg.value = 10
	arg.data[0] = 100
}

func main() {
	var a int = 100
	var p *int // 포인터 변수

	p = &a
	*p = 20

	fmt.Println(a)

	var x int = 10
	var y int = 20

	var p1 *int = &x
	var p2 *int = &x
	var p3 *int = &y

	// == 연산을 사용해 포인터가 같은 메모리 공간을 가리키는지 확인
	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p1 == p3 : %v\n", p1 == p3)

	var data Data
	ChangeData(&data)
	fmt.Println(data.value)

	//var data Data
	//var q *Data = &data

	// 구조체를 생성해 초기화
	//q := &Data{}

	// new()를 사용해 초기화
	// 타입을 메모리에 할당하고 기본값으로 채워 주소 반환
	var q = new(Data)
	ChangeData(q)
	fmt.Println(q.data[0])
}
