package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가'
	// var a int16 = 7
	// var a = 7
	// a := 7
	a := 7
	var b float64 = 5.3
	a = int(b) // Type Conversion
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c)        // unicode(UTF-8) value print
	fmt.Printf("%c\n", c) // one character print
	fmt.Printf("%T\n", c) // rune is int 32's another name

	fmt.Println(math.Round(2.71))
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java"))
}
