package main

import "fmt"

func fa(a2 [3]int) {
	a2[1] = 7
}

func fs(s2 []int) {
	s2[1] = 7
}

func main() {
	a := [...]int{2, 3, 5}
	fmt.Printf("%T\n", a)

	a1 := [3]int{2, 3, 5}
	s1 := []int{2, 3, 5}

	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", s1)
	fa(a1)
	fs(s1)
	// a[1] = 7
	// s[1] = 7

	fmt.Println(a1, s1)
}

// package main

// import "fmt"

// func main() {
// 	// var s []int
// 	// s = make([]int, 5)
// 	s := make([]int, 5)

// 	// var s []int = []int{0, 0, 0, 0, 0}
// 	// s := []int{0, 0, 0, 0, 0}

// 	if len(s) == 0 {
// 		fmt.Println("슬라이스")
// 		fmt.Println(s)
// 	}
// 	s[4] = 11
// 	s = append(s, -9)
// 	s = append(s, 10, 2)

// 	fmt.Println(s)
// 	fmt.Println(len(s))
// 	fmt.Println(s[5])
// 	// const an int = 5
// 	// var a [an]int = [5]int{2, 3, 1, 5, 4}
// 	// a[4] = 77
// 	// fmt.Println(a)
// }
