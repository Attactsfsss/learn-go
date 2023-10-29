package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) { //eval是获取返回值
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			("unsupported operation: %s" + op))
	}
}

// 13 / 3 = 4...1
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function $s with args"+"(%d,%d)", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))

}
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
func swap(a, b int) (int, int) {
	b, a = a, b
	return a, b
}
func main() {
	if result, err := eval(13, 3, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 3, 4)) //pow是指a的b次幂
	fmt.Println(sum(1, 2, 3, 4, 5))
	a, b := 3, 4
	a, b = swap(3, 4)
	fmt.Println(a, b)
}
