package main

import (
	"fmt"
	"log"
)

func main() {
	var start, last int
	fmt.Print("범위 입력(int int): ")
	_, err := fmt.Scanln(&start, &last)

	if err != nil {
		log.Fatal(err)
	}

	for i := start; i <= last; i++ {
		isPrime := true
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i, ' ')
		}
	}
}
