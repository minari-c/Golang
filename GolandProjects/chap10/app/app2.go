package main

import (
	"GolandProjects/calendar"
	"fmt"
	"log"
)

func main() {
	ev1 := calendar.Event{}
	ev2 := calendar.Event{}
	// err := ev1.SetTitle("서찬원님의 전역일") // 글자수 제한 6자리 이하

	err := ev2.SetTitle("김대현 생일")
	if err != nil {
		log.Fatal(err)
	}

	ev1.SetYear(1999)
	ev1.SetMonth(9)
	ev1.SetDay(7)

	err = ev1.SetTitle("서찬원전역일")
	if err != nil {
		log.Fatal(err)
	}

	ev1.SetYear(2020)
	ev1.SetMonth(12)
	ev1.SetDay(25)

	fmt.Println(ev2.GetTitle(), ev2.GetYear(), ev2.GetMonth(), ev2.GetDay())
	fmt.Println(ev1.GetTitle(), ev1.GetYear(), ev1.GetMonth(), ev1.GetDay())
}