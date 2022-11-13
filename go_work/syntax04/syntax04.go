package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수: ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break // stop unnecessary operations
		}
		// fmt.Print(i, " ")
	}

	if isPrime { // remove equal operator
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다!")
	}
}

// // random_number_is_prime_number? compact v6.0

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수: ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break // stop unnecessary operations
// 		}
// 		// fmt.Print(i, " ")
// 	}

// 	if isPrime { // remove equal operator
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다!")
// 	}
// }

// // random_number_is_prime_number? compact v3.0

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수: ", number)

// 	for i := 2; i <= number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 		}
// 		fmt.Print(i, " ")
// 	}

// 	if isPrime { // remove equal operator
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다!")
// 	}
// }

// // random_number_is_prime_number? compact v2.0

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수: ", number)

// 	for i := 2; i <= number; i++ {	// not looping in case 1 or number
// 		if number%i == 0 {
// 			isPrime = false	// number is not prime number
// 		}
// 	}

// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다!")
// 	}
// }

// // random_number_is_prime_number?

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // for refrash random seed
// )

// func main() {
// 	seed := time.Now().Unix() // get current time for unix format
// 	rand.Seed(seed)           // call by new_time(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 // [2 ~ 151]
// 	fmt.Println("임의로 추출된 수: ", number)

// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}

// 	if count == 2 {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다!")
// 	}
//}

// // seed of random number

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // for refrash random seed
// )

// func main() {
// 	seed := time.Now().Unix() // get current time for unix format
// 	rand.Seed(seed)           // call by new_time(seed)

// 	for i := 1; i < 6; i++ {
// 		dice := rand.Intn(6) + 1 // return random int number
// 		fmt.Println(dice)
// 	}
// }
