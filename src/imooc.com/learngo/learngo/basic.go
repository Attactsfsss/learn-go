package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "kkk"
var bb = true

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q/n", a, s)
}
func variableInitialZeroValue() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, s, b)
}
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	c = true
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	c =true
	fmt.Println(a, b, c, s)
}
func euler() {
	fmt.Printf("%.3f\n", //保留小数点后三位//
		cmplx.Exp(1i*math.Pi)+1) //Exp是e的函数//    //cmp表示复数类型//
}
func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func consts() {
	const filename = "abc.txt"
	const a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c)
}
func enums() { //enums是指枚举类型//
	const (
		cpp    = 0
		java   = 1
		python = 2
		goland = 3
	)
	fmt.Println(cpp, java, python, goland)

}
func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialZeroValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss,bb)
	euler()
	triangle()
	consts()
	enums()
}
