package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Hello world!")

	fmt.Println("The time is currently: ", time.Now())

	fmt.Printf("This is a random number %v and a float %v\n", rand.Intn(10), 3.14)

	fmt.Println("I will loop ten times")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Printf("\n")

	fmt.Printf("This %v is  3 * 321", MultiplyBy321(3))

	testfizzbuzz()
}

func MultiplyBy321(val int) int {

	result := val * 321.0
	return result
}

func testfizzbuzz() {
	for i := 1; i <= 100; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int) {
	if i%3 == 0 {
		fmt.Println("fizz")
	} else if i%5 == 0 {
		fmt.Println("buzz")
	} else if i%15 == 0 {
		fmt.Println("buzz")
	} else {
		fmt.Println(i)
	}
}
