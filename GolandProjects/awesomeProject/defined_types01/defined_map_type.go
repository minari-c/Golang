/*
 map 타입의 사용자 정의 타입을 만들고 자기만의 적절한 메서드를 추가하여 코드를 작성하고 설명하시오.
*/

package main

import (
	"fmt"
	"time"	// 수행 시간을 위한 time package
)

type Schedule map[time.Time]string	// 내 할일과 할일을 수행할 시간을 정의한다.

func (s Schedule) workInfo() {	// method
	for _, work := range s {	// 오늘의 할일만 출력한다.
		fmt.Println(work)
	}
}

func main() {
	mySchedule := Schedule{
		time.Now(): "밥 먹기",	// map type의 초기화 형식을 따른다.
	}

	mySchedule.workInfo()		// method 사용
}