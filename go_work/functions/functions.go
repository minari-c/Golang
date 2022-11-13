package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다~", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

//	구간 소수 판정 프로그램 v1.2
func main() {
	var a, b int

	fmt.Print("정수 입력: ")
	_, err := fmt.Scanln(&a, &b)

	if err != nil {
		log.Fatal(err)
	}

	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		if p {
			fmt.Print(i, " ")
		}
	}
}

// package main

// import "fmt"

// func processScore(name string, kor int, eng int, math int) (int, int) {
// 	totalScore := kor + eng + math
// 	average := totalScore / 3

// 	return totalScore, average
// 	// fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", name, totalScore, average)
// }

// func main() {
// 	var t int
// 	var a int

// 	t, a = processScore("홍길동", 100, 90, 93)
// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길동", t, a)

// 	t, a = processScore("홍길순", 89, 91, 92)
// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길순", t, a)
// }

// package main

// import "fmt"

// func sayHello() {
// 	fmt.Println("Hello!~")
// }

// func main() {
// 	sayHello()
// }
