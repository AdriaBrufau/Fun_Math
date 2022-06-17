package problems

import (
	"fmt"
)

func SumOfPrimes() {
	// 	The prime 41, can be written as the sum of six consecutive primes:

	// 41 = 2 + 3 + 5 + 7 + 11 + 13
	// This is the longest sum of consecutive primes that adds to a prime below one-hundred.

	// The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

	// Which prime, below one-million, can be written as the sum of the most consecutive primes?
	
	fmt.Print(generatePrime(100))
	
}

func generatePrime(limit int) [] int{
	//sieve of atkin implementation
	var primes = make([]bool, limit)
	var returnedPrimes []int

	for i := 0; i < limit; i++{
		primes[i] = false
	}

	for x := 1; x < limit; x++{
		for y := 1; y < limit; y++{
			var n int = (4 * x * x) + (y * y)
			if n < limit && (n % 12 == 1 || n % 12 == 5){
				primes[n] = true
			}
			n = (3 * x * x) + (y * y)
			if n < limit && (n % 12 == 7){
				primes[n] = true
			}
			if x > y {
				n = (3 * x * x) - ( y * y)
				if n < limit && (n % 12 == 11){
					primes[n] = true
				}
			}
		}
	}
	primes[2] = true
	primes[3] = true
	for l := 5; l * l <= limit; l++ {
		if primes[l]{
			for i := l * l; i < limit; i += l * l{
				primes[i] = false;
			}
		}
	}

	for g := 0; g  < limit; g++{
		if primes[g]{
			fmt.Sprintf("g: %d", g)
			returnedPrimes = append(returnedPrimes, g)
		}
	}

	return returnedPrimes
}
