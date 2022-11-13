package main

import (
	"fmt"
	"go_project_home/chap10/calendar"
	"log"
)

func main() {
	date := calendar.Date{}

	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(2)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(20)
	if err != nil {
		log.Fatal(err)
	}

	//err = date.SetYear(-500)
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println(date)

	fmt.Println(date.GetYear())
	fmt.Println(date.GetMonth())
	fmt.Println(date.GetDay())

}
