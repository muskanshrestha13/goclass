package main

import "fmt"

var a int8 = -1

const PI = 3.14

func main() {
	var testingUint8 uint8 = 1
	c := "a"
	a += 2      // a= a+2
	b := 10 % 2 //remainder
	x := 2 << 2 //0010(binary of 2)
	testingUint8 = 2
	fmt.Println(a)
	fmt.Println(" VAL OF b =", b)
	fmt.Println(" VAL OF x =", x)
	fmt.Println(testingUint8)
	fmt.Printf("%T\n", c)
	add()

}
func add() {
	c := 5
	fmt.Println(c)
}
