package main

import (
	"fmt"
	"os"
	"strconv"
)

func sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: primes <limit>")
		os.Exit(1)
	}
	limit, err := strconv.Atoi(os.Args[1])
	if err != nil || limit < 2 {
		fmt.Println("Please provide a valid integer >= 2")
		os.Exit(1)
	}
	primes := sieve(limit)
	fmt.Printf("Found %d primes up to %d:\n", len(primes), limit)
	for i, p := range primes {
		if i > 0 && i%10 == 0 {
			fmt.Println()
		}
		if i%10 == 9 {
			fmt.Printf("%d\n", p)
		} else {
			fmt.Printf("%d ", p)
		}
	}
	fmt.Println()
}
