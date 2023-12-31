package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string { //二进制函数
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result //strconv.Itoa(）将数转为字符串
	}
	return result
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}
}
func forever() {
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
		convertToBin(72387885),
		convertToBin(0),
	)
	printFile("abc.txt")
	forever()

}
