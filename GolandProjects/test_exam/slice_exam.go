package main

import "fmt"

func deep_copy(a []int, b []int) {
	for i, v := range a {
		b[i] = v
	}
}

func change_slice( a *[]int, b *[]int ) {
	temp := *a
	*a = *b
	*b = temp
}

func make_slice() []int {
	return []int {1, 2, 3}
}

func main() {
	var slice = []int { 1, 2, 3 }
	var temp = make_slice()

	change_slice(&slice, &temp)

	fmt.Println(slice, temp)
}

/*
	fmt.Println("slice", slice)

	fmt.Println("slice[:]", slice[:])

	fmt.Println("slice[1:]", slice[1:])

	fmt.Println("slice[:1]", slice[:1])


*/
