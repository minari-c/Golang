/*

OPP: JAVA, C++, C#

Go: receiver

 */

package main

import "fmt"

type Number int
type MyType string

func (n *Number) Square() {		// method
	*n = *n * *n
}

func pointerSquare(n *int) {	// function
	*n = *n * *n
}

func (m MyType) sayHi(n int, f bool) {
	fmt.Println("Test: " + m)
}

func (m MyType) myReturn() int {
	return len(m)
}

func main() {
	var t1 = MyType("메서드 테스트 중")
	t2 := MyType("단축 연산자도 사용가능")
	t2.sayHi(9, true)
	t1.sayHi(-7, false)

	fmt.Println(t2.myReturn())
	fmt.Println(t1.myReturn())


	n1 := Number(5)	// method의 argument는 자동으로 포인터 타입으로 형변환 해줌
	n1.Square()
	n2 := 4
	pointerSquare(&n2)
	fmt.Println(n1, n2)
}

//type Gallons float64
//type Milliliters float64
//
//
//func ToLiters(g Gallons) Liters {
//	return Liters(g * 3.785)
//}
//
//func ToGallons(l Liters) Gallons {
//	return Gallons(l * 0.264)
//}
//
//func ToGallons(ml Milliliters) Gallons {
//	return Gallons(ml * 0.000264)
//}
//
//func main() {
//	var carfue1 Liters = 10
//	var busfue1 Gallons = 60
//
//	carfue1 = Liters(15.9)
//
//	carfue1 = Liters(Gallons(2.5) * 0.264)
//	busfue1 = Gallons(Liters(240.0) * 3.785)
//
//	fmt.Println(carfue1)
//	fmt.Println(busfue1)
//
//	fmt.Println(Gallons(1.2) + Gallons(5.5))
//
//	fmt.Println("defined types")
//}
