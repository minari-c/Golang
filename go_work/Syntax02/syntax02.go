package main

import (
	"fmt"
	// "reflect"
	"strings"
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)

	r := strings.NewReplacer("@", "o") // fomatter?
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

	/*
		variable naming convention:
			must:
				-first character start with an english letter
				-when variable can accessed from other package than, use a capital letter

			recommend:
				-use camelcase
				-use simple character[i, j, k, ...] for iteration count variable
	*/
	// zero value -> defualt value all type [bool is false, string is blank]
	// var E string // variables exposed externally use capital letter
	// var d bool
	// var c rune
	// var b float64
	// var a int
	// // a := 7

	// // naming convention
	// var studentId string
	// var i int32
	// fmt.Println(studentId)
	// fmt.Println(i)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(E)
	// fmt.Println()

	// fmt.Printf("%T\n", E)
	// fmt.Printf("%T\n", d)
	// fmt.Printf("%T\n", c)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n\n", a)

	// fmt.Println(reflect.TypeOf(E))
	// fmt.Println(reflect.TypeOf(d))
	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(b))
	// fmt.Println(reflect.TypeOf(a))
}

// terminal command: go run [go file name]
