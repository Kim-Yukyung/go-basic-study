package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//// 0~99 랜덤 숫자 반환
	//target := (int)(rand.Float64() * 100)

	//rand.Seed(time.Now().UnixNano()) // deprecated

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	target := rng.Intn(100)

	var number int
	cnt := 1

	for {
		fmt.Print("숫자값을 입력하세요: ")

		// 사용자 입력
		_, err := fmt.Scanln(&number)
		if err != nil {
			return
		}

		if number < target {
			fmt.Println("입력하신 숫자가 더 작습니다.")
		} else if number > target {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		} else {
			fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", cnt)
			break
		}
		cnt++
	}
}
