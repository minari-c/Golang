package main

import "fmt"

func main() {
	fruit := map[int]string {
		1: "apple",
		2: "banana",
		3: "grape",
	}

	var b int = 1
	var c int = 0

	for k, v := range fruit {
		fmt.Println("key:", k, "value:", v)
	}


	val, exits := fruit[b]
	if exits {
		fmt.Println("they have a", val)
	} else {
		fmt.Println( b, "is not fruit code")
	}

	val, exits = fruit[c]
	if exits {
		fmt.Println("they have a", val)
	} else {
		fmt.Println(c, "is not fruit code")
	}
}
