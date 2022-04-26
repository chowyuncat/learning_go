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

func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sum int

	// runtime.GOMAXPROCS(1)
	fmt.Printf("Limited to %d cpus\n", runtime.GOMAXPROCS(0))

	loops := 0

	if false {
		primes := MakeGeneratorChannel()

		i := 0
		for prime := range primes {
			loops++
			sum += prime
			i++
			if i >= N {
				break
			}
		}
	}

	if true {
		getprime := MakeGeneratorClosure()
		for i := 0; i < N; i++ {
			loops++
			prime := getprime()
			sum += prime
		}
	}

	fmt.Printf("sum=%d, loops=%d\n", sum, loops)
}
