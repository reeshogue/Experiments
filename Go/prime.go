package main

import "fmt"

/*
Sift the Two's and Sift the Three's,
The Sieve of Eratosthenes.
When the multiples sublime,
The numbers that remain are Prime.
*/

func sieve(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for j := i * i; j < N; j += i {
			b[j] = true
		}

	}
	return
}

func main() {
	primes := sieve(100)
	for _, i := range primes {
		fmt.Println(i)
	}
}
