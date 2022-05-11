package main

import (
	// "math/rand"
	// "fmt"
)

func EmptyFunc() {
	return
}

func OneOutputFunc() int {
	return 16
}

func OneIntInputFunc(x int) {
	return
}

func OneIntInputFuncRetOneInt(x int) int {
	return x + x
}

func TwoIntInputFuncRetOneInt(x, y int) int {
	return x + y
}

func TwoIntInputFuncRetTwoInt(x, y int) (int, int) {
	return x * y, x + y
}

func ThreeIntInputFuncRetOneInt(x, y, z int) int {
	return x + y + z
}

func main() {
	// r := rand.Intn(10)

	EmptyFunc()

	x := OneOutputFunc()

	OneIntInputFunc(x)
	
	y := OneIntInputFuncRetOneInt(x)

	TwoIntInputFuncRetOneInt(x, y)

	TwoIntInputFuncRetTwoInt(x, y)


}
