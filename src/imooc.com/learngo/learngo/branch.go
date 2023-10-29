package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default: //超过3个条件的用switch，case，default，3个以内用if，else语句；default是指所有的case都不成立时所要执行的语句//
		panic(fmt.Sprintf("Wrong score: %d", score)) //Sprintf是字符串格式化命令；panic是当系统发现无法运行下去的故障时将调用它//
	}
	return g
}
func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
	)
}
