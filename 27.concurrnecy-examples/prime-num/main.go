package main

import (
	"fmt"
	"sort"
	"sync"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	} 
		return true
}

func worker(jobs <- chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		if isPrime(n) {
			results <- n
		}
	}
}

func main() {
	start, end := 2, 100
	numWorkers := 4


	jobs := make(chan int, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i ++{
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		for n := start; n <= end; n++ {
			jobs <- n
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var primes []int
	for p := range results {
		primes = append(primes, p)
	}

	sort.Ints(primes)

	fmt.Printf("Primes between %d and %d:\n", start, end)
	for _, p := range primes {
		fmt.Print(p, " ")
	}
	fmt.Println()
}