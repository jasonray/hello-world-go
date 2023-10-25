package main

import "fmt"

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
}