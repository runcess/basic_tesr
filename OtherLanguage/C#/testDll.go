package main

import (
	"C"
	"fmt"
)

//export test1
func test1() {
	fmt.Println("==============================")
	fmt.Println("테스트1")
	a := 3
	b := 5
	fmt.Println("a+b=", a+b, "<testDLL.dll>")
}

//export test2
func test2(a int, b float64) float64 {
	fmt.Println("===============================")
	fmt.Println("테스트2")
	fmt.Println(a, "+", b, "=", float64(a)+b, "<testDLL.dll>")
	return (float64(a) + b)
}

//export test3
func test3(arg *C.char) *C.char {
	fmt.Println("===============================")
	fmt.Println("테스트3")
	input := C.GoString(arg)
	fmt.Println("입력값은", input, "입니다.", "<testDLL.dll>")
	return C.CString("반환값은 " + input + " 입니다" + "<testDLL.dll>")
}

func main() {
}

// https://linsoo.pe.kr/archives/26773
// >go build -o testDLL.dll -buildmode=c-shared .\testDLL.go
