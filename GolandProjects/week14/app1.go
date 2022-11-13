package main

import (
	"fmt"
	"time"
)
func a() {
	for i := 0; i < 20; i++ {
		fmt.Print("A")
	}
}

func b() {
	for i := 0; i < 20; i++ {
		fmt.Print("B")
	}
}

func c() {
	fmt.Print("C")
}

func d() {
	fmt.Print("D")
}

func main() {
	go a()	// go 루틴을 작업 쓰레드에 할당

	go b()

	// go c()
	defer c()	// 속해있는 함수 종료 시점에 실행 file close 할 때 많이 사용
	defer d()

	time.Sleep(time.Second * 2)

	fmt.Println("메인함수 끝")

	// main Thread와 다른 Thread가 경쟁을 한다. -> 일반적으로는 main Thread에 우선 할당 된다.
}
