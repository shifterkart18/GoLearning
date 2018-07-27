package main

import (
	"GoLearning/myutil"
	"fmt"
	"math"
	"runtime"
	"time"
)

const Name string = "Andrew Martz"

var Number int32 = 18

func main() {

	fmt.Println("Hello world!")

	fmt.Println("The time is currently: ", time.Now())

	fmt.Println("My name is: ", Name)

	fmt.Println("A package variable: ", Number)
	Number := 21
	fmt.Println("A package variable with a new value: ", Number)

	fmt.Printf("This is how a formatted string works:\nint: %v \nfloat: %v \nstring: %v\n", 18, math.Pi, "Andrew\n")

	FizzBuzz(30)

	Test321(-100, 100)

	MyName := "Andrew Martz"
	fmt.Printf("the reverse of %v is %v", MyName, blah.ReverseLetters(MyName))

	//"First Second Third"
}

func PrintOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
