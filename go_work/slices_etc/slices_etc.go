package main

import (
	"fmt"
	"sort"
)

func main() {
	p := []int{111, 19, 1, 17, 15}
	sort.Ints(p)
	fmt.Println(p)

	idx := 2
	// for k := idx + 1; k < len(p); k++ {
	// 	p[k-1] = p[k]
	// }
	// p = p[:len(p)-1]
	// fmt.Println(p)

	p = append(p[:idx], p[idx+1:]...)
	fmt.Println(p)
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func numbers(n1 int, n2 ...int) {
// 	fmt.Println(n1, n2)
// }

// func circles(radii ...float64) {
// 	for _, radius := range radii {
// 		fmt.Println(radius * radius * math.Pi)
// 	}
// }

// func prims(numbers ...int) {
// 	for _, number := range numbers {
// 		if number <= 2 {
// 			continue
// 		}

// 		var isPrime = true
// 		for i := 2; i < number; i++ {
// 			if number%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			fmt.Print(number, " ")
// 		}
// 	}
// 	fmt.Println()
// }

// func main() {
// 	p := []int{111, 19, 1, 17}

// 	prims(111, 19, 1, 17)
// 	// prims(p)
// 	prims(p...)

// 	circles(10.0, 100.0, 5.0)

// 	numbers(5, 10)
// 	numbers(5, 10, 3, 2)

// 	// fmt.Println("a")
// 	// fmt.Println("b", "c", "d")
// }
