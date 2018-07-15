package Mult321

import "fmt"

//MultiplyBy321 multiply a number by 321
func MultiplyBy321(val int) int {

	return val<<8 + val<<6 + val
}

//Test321 test Multiply321
func Test321(Start int, End int) {

	bSuccess := true
	for i := Start; i < End; i++ {
		NormalVal := i * 321
		TestVal := MultiplyBy321(i)

		if TestVal != NormalVal {
			bSuccess = false
			fmt.Printf("The output for MultiplyBy321 was not correct for %v\n", i)
			fmt.Printf("NormalVal = %v\n", NormalVal)
			fmt.Printf("TestVal = %v\n", TestVal)
		}
	}

	if bSuccess {
		fmt.Printf("The output for MultiplyBy321 was correct for values %v to %v\n", Start, End)
	}
}
