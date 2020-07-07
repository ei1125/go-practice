package main

import "fmt"

//関数の外でも宣言できる、型を宣言したいとき
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	xi = 2
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)

}

func main() {

	fmt.Println(i, f64, s, t, f)
	//関数の外では宣言できない
	foo()
}
