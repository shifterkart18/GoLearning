package FizzBuzz

import "fmt"

//FizzBuzz test if a programmer can code themselves out of a wet paper bag
func FizzBuzz(Num int) {
	for i := 1; i <= Num; i++ {
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
}
