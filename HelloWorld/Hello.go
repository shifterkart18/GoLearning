package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Hello world!")

	fmt.Println("The time is currently: ", time.Now(), time.Now(), time.Now())

	fmt.Printf("This is a random number %v blah %v: ", rand.Intn(10), 1123)

	fmt.Println("I will now count to ten")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("This %v is  3 * 321", MultiplyBy321(3))
}

func MultiplyBy321(val int) int {

	result := val * 321
	return result
}
