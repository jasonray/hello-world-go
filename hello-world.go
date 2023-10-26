package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	var greeting = "hello world!"
	fmt.Println(greeting)

	fmt.Println("go" + "lang")

	var x = 1
	var y = 2
	var sum = x + y
	fmt.Println("1+2 =", sum)

	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println("constant outside of function:", s)

	const cn = 500000000
	const cd = 3e20 / cn
	fmt.Println("calculated value:", cd)

	const vn = 500000000
	const vd = 3e20 / vn
	fmt.Println("calculated value:", vd)

	fmt.Println(int64(cd))

	fmt.Println(math.Sin(cn))
}
