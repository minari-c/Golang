package main

import (
	"date"
	"fmt"
	"log"
)


// 현재는 복사본 리시버, 리시버는 파라미터?
// 원본을 수정하고 싶을 때는 포인터 리시버를 사용해라
// 포인터 리시버: ( 메소드 )
// 리시버: ( 생성자? )


func main() {
	date := Date{Year: 2021, Month: 11, Day: 15 }

	fmt.Println(date)

	date.Year = -500
	err := date.SetYear(-2022)
	if err != nil {
		log.Fatal(err)
	}
	date.Month = 13
	date.Day = 133

	fmt.Println(date)
}
