package main

import (
	"fmt"
	"time"
)


type student struct {
	id      int
	name    string
	subject string
	seconds int
}

func (s student) study() {
	fmt.Printf("%s은 %s과목을 공부합니다.\n", s.name, s.subject)
}

func (s student) countDown() {
	for s.seconds > 0 {
		fmt.Println(s.seconds)
		time.Sleep(time.Second)
		s.seconds--
	}
}

func funcStudy(fs student) {
	fmt.Printf("%s은 %s과목을 공부합니다.\n", fs.name, fs.subject)
}

func main() {
	s1 := student{1000, "홍길동", "Go", 5}

	fmt.Println(s1)
	s1.study()

	s1.countDown()

	funcStudy(s1)
}


//type Liters float64
//type Gallons float64
//type Milliliters float64
//
//func (g Gallons) GallonsToLiters() Liters {
//	return Liters(g * 3.785)
//}
//
//func (g Gallons) GallonsToMilliliters() Liters {
//	return Liters(g * 3.785 * 1000.0)
//}
//
//func (l Liters) ToGallons() Gallons {
//	return Gallons(l * 0.264)
//}
//
//func (ml Milliliters)ToGallons() Gallons {
//	return Gallons(ml * 0.000264)
//}
//
//func main() {
//	var coke Liters = Liters(2)
//	fanta := Milliliters(750)
//
//	fmt.Printf("%.3f리터는 %0.3f갤론과 같습니다.\n", coke, coke.ToGallons())
//	fmt.Printf("%.3f밀리리터는 %0.3f갤론과 같습니다.\n", fanta, fanta.ToGallons())
//}
