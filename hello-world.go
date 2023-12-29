package main

import (
	"fmt"
	"strconv"
)

func add(x int, y int) int {
	return x + y
}

func main2() {
	// https://go.dev/ref/spec

	// declare variables
	// can be explicit
	var greeting string = "hello world!"
	fmt.Println((greeting))

	// can reassign
	greeting = "hello, again"
	fmt.Println((greeting))

	// can infer
	var greeting2 = "hello #2"
	fmt.Println(greeting2)

	// can declare and assign
	greeting3 := "hello #3"
	fmt.Println((greeting3))

	// can declare as a constant
	const s string = "this is a constant value"
	fmt.Println((s))

	// concat strings
	fmt.Println("go" + "lang")

	// three different ways to declare / set variables
	var x int = 1
	var y = 2
	z := 3
	var sum = x + y + z
	fmt.Println("1+2+3 =", sum)

	// use custom function
	s12 := add(1, 2)
	fmt.Println("s12: ", s12)

	// functions can be variables
	var f = add
	s23 := f(2, 3)
	fmt.Println("s23: ", s23)

	//loops
	// feels like a while loop
	i := 1
	for i <= 3 {
		fmt.Println("while loop: ", i)
		i = i + 1
	}

	// for loop
	for i := 1; i <= 3; i++ {
		fmt.Println("for loop: ", i)
	}

	//supports break / continue
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// & is and, || is or, ! is not
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//what happens when you treat string as boolean
	b := true
	fmt.Println(b)
	// the following would be compile time error, perfect!
	// b = "some string"

	var boolValues = []string{
		"1",
		"t",
		"T",
		"TRUE",
		"true",
		"True",
		"0",
		"f",
		"F",
		"FALSE",
		"false",
		"False",
		"taco",
	}
	for _, v := range boolValues {
		boolValue, err := strconv.ParseBool(v)
		if err != nil {
			fmt.Printf("failed to convert '%s' to boolean \n", v)
		}
		fmt.Printf("convert %s => %t\n", v, boolValue)
	}

	b2 := true
	if b2 {
		println()
	}

}
