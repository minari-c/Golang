package main

import (
	"bufio"
	"fmt"
	"log" // not mathmetical log, it just a logging, and we use a Fatal
	"os"
	"strconv" // TrimSpace()
	"strings" // ParseInt()
)

func main() {
	fmt.Print("단 입력: ")
	rd := bufio.NewReader(os.Stdin) // standard input stream?
	in, err := rd.ReadString('\n')  // underscore sign is None or Pass

	if err != nil { // you can used small bracket or not used
		log.Fatal(err)
	}

	in = strings.TrimSpace(in) // remove a front-back blank
	// dan, err := strconv.ParseInt(in, 10, 32) // ParseInt(string, ["0b", "0o", ... ], bit)
	dan, err := strconv.Atoi(in) // array to int?
	if err != nil {
		log.Fatal(err)
	}

	// while form
	i := 1
	for i < 10 {
		fmt.Printf("%d * %d = %d\n", dan, i, dan*i)
		i++
	}

	// for i := 1; i < 10; i++ {
	// 	// fmt.Println(dan, " * ", i, " = ", dan * i)
	// 	fmt.Printf("%d * %d = %d\n", dan, i, dan*i)
	// }
}

// package main

// import (
// 	"bufio" // buffer i/o
// 	"fmt"
// 	"os" // os os os ~
// )

// func main() {
// 	fmt.Print("숫자 입력: ")
// 	rd := bufio.NewReader(os.Stdin) // standard input stream?
// 	in, _ := rd.ReadString('\n')    // underscore sign is None or Pass
// 	fmt.Print(in)
// }
