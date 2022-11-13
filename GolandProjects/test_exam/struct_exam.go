package main

import "fmt"

type Point struct {
	x int
	y int
}


func main() {
	p := Point{x: 1, y: 2}

	fmt.Println(p)
}
