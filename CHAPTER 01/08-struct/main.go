package main

import "fmt"

type Student struct {
	Name  string // 대문자로 시작하는 필드는 외부로 공개
	Class int
	No    int
	Score float64
}

type User struct {
	Name string
	ID   int
	Age  int
}

type VIPUser struct {
	//UserInfo User
	User     // 포함된 필드
	VIPLevel int
	Price    int
}

func main() {
	var a Student
	a.Name = "Mike"
	a.Class = 1
	a.No = 5
	a.Score = 100
	//var a Student = Student{"Mike", 1, 5, 100}
	fmt.Println(a)

	var b Student = Student{Name: "Emma", Score: 75.5}
	fmt.Println(b)

	user := User{"Nick", 1, 20}
	vip := VIPUser{
		User{"Chloe", 2, 25},
		3,
		100,
	}
	fmt.Println(user)
	//fmt.Println("VIP Name:", vip.UserInfo.Name)
	fmt.Println("VIP Name:", vip.Name)

	var student1 Student = Student{"Alice", 3, 2, 80}
	var student2 Student = Student{"Mark", 3, 3, 50.5}
	fmt.Println(student1, student2)
	student2 = student1
	fmt.Println(student1, student2)
}
