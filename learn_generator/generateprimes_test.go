package main

import (
	"math/big"
	"testing"
)

func TestBoth(t *testing.T) {

	primes := MakeGeneratorChannel()
	getprime := MakeGeneratorClosure()
	i := 0
	var tmp big.Int

	for prime := range primes {
		otherprime := getprime()

		if tmp.SetInt64(int64(prime)); !tmp.ProbablyPrime(1) {
			t.Fatalf("[%d]: %d is not prime", i, tmp.Int64())
		}

		if otherprime != prime {
			t.Fatalf("[%d]: %d != %d", i, prime, otherprime)
		}

		if i++; i >= BenchmarkCount/2 {
			break
		}
	}

	t.Logf("%d primes checked", i)
}

const BenchmarkCount = 30000

func BenchmarkMakeGeneratorChannel(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		primes := MakeGeneratorChannel()

		i := 0
		sum := 0
		for prime := range primes {
			sum += prime
			i++
			if i >= BenchmarkCount {
				break
			}
		}
	}

}

func BenchmarkMakeGeneratorClosure(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		getprime := MakeGeneratorClosure()

		sum := 0
		for i := 0; i < BenchmarkCount; i++ {
			prime := getprime()
			sum += prime
		}
	}
}
