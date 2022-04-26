package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func MakeGeneratorChannel() chan int {
	primes := make(chan int)

	go func() {

		primes <- 2

		for candidate := 3; true; candidate++ {

			i := 2
			for ; i < candidate/2+1; i++ {
				if candidate%i == 0 {
					break
				}
			}

			if i == candidate/2+1 {
				primes <- candidate
			}

		}
	}()

	return primes
}

func MakeGeneratorClosure() func() int {

	candidate := 1

	getprime := func() int {

		if candidate == 1 {
			candidate = 2
			return 2
		}

		for {
			candidate++

			i := 2
			for ; i < candidate/2+1; i++ {
				if candidate%i == 0 {
					break
				}
			}

			if i == candidate/2+1 {
				return candidate
			}
		}
	}

	return getprime
}

func DoChannel(N int) int {
	primes := MakeGeneratorChannel()

	var sum int

	i := 0
	for prime := range primes {
		sum += prime
		i++
		if i >= N {
			break
		}
	}

	return sum
}

func DoClosure(N int) int {
	var sum int

	getprime := MakeGeneratorClosure()
	for i := 0; i < N; i++ {
		prime := getprime()
		sum += prime
	}

	return sum
}

func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// runtime.GOMAXPROCS(1)
	fmt.Printf("Limited to %d cpus\n", runtime.GOMAXPROCS(0))

	var sum int

	switch os.Args[2] {
		case "channel": {
			sum = DoChannel(N)
			break
		}

		case "closure": {
			sum = DoClosure(N)
			break
		}

		default:
			panic(fmt.Sprintf("unknown implementation %s", os.Args[2]))
	}

	fmt.Printf("sum=%d\n", sum)
}
