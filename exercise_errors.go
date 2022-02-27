package main

import (
	"fmt"
	"math"
)

/*
	fmt.Sprint(e) will call e.Error() to convert the value e to a string.
	If the Error() method calls fmt.Sprint(e), then the program recurses until out of memory.
	We can break the recursion by converting the e to a value without a String or Error method.
*/

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// fmt.Sprintf("%v", e) will cause infinite loop
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
