package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello world!")

	fmt.Println("The time is currently: ", time.Now())

	fmt.Printf("This is how a formatted string works:\nint: %v \nfloat: %v \nstring: %v\n", 18, 3.14, "Andrew\n")

	FizzBuzz(30)

	Test321(-100, 100)
}
