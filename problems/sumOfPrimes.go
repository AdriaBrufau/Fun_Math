package problems

import (
	"fmt"
)

func SumOfPrimes() {
	var limit int = 1000000
	varr := generatePrime(limit)
	returnedPrime, returnedarrayOfPrimes := checkSumOfPrimes(varr)
	// 	The prime 41, can be written as the sum of six consecutive primes:

	// 41 = 2 + 3 + 5 + 7 + 11 + 13
	// This is the longest sum of consecutive primes that adds to a prime below one-hundred.

	// The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

	// Which prime, below one-million, can be written as the sum of the most consecutive primes?
	fmt.Printf("under %v\n", limit)
	fmt.Printf("returnedPrime %v,\nreturnedArray: %v\n", returnedPrime, returnedarrayOfPrimes)
}

func checkSumOfPrimes(arr []int) (int, []int) {
	var sum, currentPrime int
	currentComposition, composition := make([]int, 0), make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			sum += arr[j]
			composition = append(composition, arr[j])

			if arr[i] == sum && sum > currentPrime {
				currentPrime = arr[i]
				currentComposition = composition
			}
			if arr[i] == 953 {
				fmt.Println("sum: ", sum)
				fmt.Println("composition", composition)
			}

		}
		composition = make([]int, 0)
		sum = 0
	}

	return currentPrime, currentComposition
}

func generatePrime(limit int) []int {
	//sieve of atkin implementation
	var primes = make([]bool, limit)
	var returnedPrimes []int

	for i := 0; i < limit; i++ {
		primes[i] = false
	}
	fmt.Print("starting sieve....")
	for x := 1; x < limit; x++ {
		for y := 1; y < limit; y++ {
			var n int = (4 * x * x) + (y * y)
			if n < limit && (n%12 == 1 || n%12 == 5) {
				primes[n] = true
			}
			n = (3 * x * x) + (y * y)
			if n < limit && (n%12 == 7) {
				primes[n] = true
			}
			if x > y {
				n = (3 * x * x) - (y * y)
				if n < limit && (n%12 == 11) {
					primes[n] = true
				}
			}
		}
		fmt.Println("x: ", x, "prime: ", primes[x])
	}
	fmt.Print("completed sieve....")

	primes[2] = true
	primes[3] = true
	for l := 5; l*l <= limit; l++ {
		if primes[l] {
			for i := l * l; i < limit; i += l * l {
				primes[i] = false
			}
		}
	}
	fmt.Print("assigning numbers....\n\n\n")
	for g := 0; g < limit; g++ {
		if primes[g] {
			fmt.Printf("g: %v\n", g)
			returnedPrimes = append(returnedPrimes, g)
		}
	}

	return returnedPrimes
}
