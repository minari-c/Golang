package main

import "fmt"

func main() {
	// var primes [3]int

	primes := [3]int{2, 3, 5}

	// test := [5]bool{true, true, true}
	// fmt.Println(test[3])

	// fmt.Print("Arrays: ")
	// fmt.Println(primes[2])

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(primes[i])
	// }

	// for i := 0; i < len(primes); i++ {
	// 	fmt.Println(primes[i])
	// }

	// for index, prime := range primes {
	// 	fmt.Println(index, prime)
	// }

	// for prime := range primes {
	// 	fmt.Println(prime)
	// }

	// for _, prime := range primes {
	// 	fmt.Println(prime)
	// }

	var sum int = 0
	for _, prime := range primes {
		sum = sum + prime
	}
	fmt.Println(sum)
	fmt.Println(float64(sum) / float64(len(primes)))
	fmt.Printf("%.2f\n", float64(sum)/float64(len(primes)))
	// i := 0
	// // for i < 3 {
	// for i < len(primes) {
	// 	fmt.Println(primes[i])
	// 	i++
	// }
}
