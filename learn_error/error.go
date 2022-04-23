package main

import "fmt"

type CustomError int

// func (err CustomError) Error() string {
//  	return fmt.Sprintf("Error(%v)", int(err))
// }

func (err CustomError) String() string {
	return fmt.Sprintf("Stringer %d", int(err))
}

func main() {
	var e CustomError = CustomError(1)

	// fmt first checks to see if the value implements error, then checks for Stringer
	fmt.Println(e)
}

