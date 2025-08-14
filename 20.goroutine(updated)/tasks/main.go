package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 1: Basic WaitGroup Example
// Demonstrates basic goroutine creation and synchronization with WaitGroup
func task1BasicWaitGroup() {
	fmt.Println("=== Task 1: Basic WaitGroup Example ===")
	
	var wg sync.WaitGroup
	wg.Add(1)
	
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Printf("Hello from ayush - %d", i+1)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	wg.Wait()
	fmt.Println("Task 1 completed!")
}

// Task 2: Basic Goroutine with Named Function
// Shows how to use goroutines with named functions and timing
func task2BasicGoroutine() {
	fmt.Println("=== Task 2: Basic Goroutine with Named Function ===")
	
	sayHello := func(name string) {
		for i := 0; i < 3; i++ {
			fmt.Printf("Hello from %s - %d", name, i+1)
			time.Sleep(100 * time.Millisecond)
		}
	}

	go sayHello("Goroutine")
	sayHello("Main")

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Task 2 completed!")
}

// Task 3: Multiple Goroutines
// Demonstrates launching multiple goroutines concurrently
func task3MultipleGoroutines() {
	fmt.Println("=== Task 3: Multiple Goroutines ===")
	
	worker := func(id int) {
		for i := 0; i < 3; i++ {
			fmt.Printf("Worker %d: task %d", id, i+1)
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Start 10 workers
	for i := 0; i < 10; i++ {
		go worker(i) // i is the worker id
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Task 3 completed!")
}

// Task 4: Using WaitGroup Properly
// Shows correct usage of WaitGroup with multiple workers
func task4WaitGroupProper() {
	fmt.Println("=== Task 4: Using WaitGroup Properly ===")
	
	var wg sync.WaitGroup
	
	worker := func(id int) {
		defer wg.Done() // Important: defer ensures Done() is called even if panic occurs
		for i := 0; i < 3; i++ {
			fmt.Printf("Worker %d: task %d", id, i+1)
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Start workers with proper WaitGroup management
	for i := 1; i <= 3; i++ {
		wg.Add(1)   // Increment counter before starting goroutine
		go worker(i) // i is the worker id
	}

	wg.Wait() // Wait for all workers to complete
	fmt.Println("All workers completed")
	fmt.Println("Task 4 completed!")
}

// Task 5: Simple Channel Communication
// Demonstrates basic channel usage for goroutine communication
func task5SimpleChannels() {
	fmt.Println("=== Task 5: Simple Channel Communication ===")
	
	ch := make(chan string) // Create a channel

	go func() {
		ch <- "Hello from goroutine!"
		ch <- "Second message"
		close(ch) // Always close channels when done sending
	}()

	for msg := range ch {
		fmt.Println("Received:", msg)
	}
	fmt.Println("Task 5 completed!")
}

// Task 6: Goroutine with Channel Return
// Shows how to use channels to get results from goroutines
func task6ChannelReturn() {
	fmt.Println("=== Task 6: Goroutine with Channel Return ===")
	
	calculate := func(a, b int, result chan<- int) {
		time.Sleep(100 * time.Millisecond) // Simulate computation time
		result <- a + b
	}

	resultCh := make(chan int)

	go calculate(10, 20, resultCh)

	fmt.Println("Doing other work...")
	time.Sleep(50 * time.Millisecond)

	result := <-resultCh
	fmt.Printf("Calculation result: %d", result)
	fmt.Println("Task 6 completed!")
}

// Task 7: Producer-Consumer Pattern
// Demonstrates the classic producer-consumer pattern using channels
func task7ProducerConsumer() {
	fmt.Println("=== Task 7: Producer-Consumer Pattern ===")
	
	jobs := make(chan int, 5) // Buffered channel for jobs
	done := make(chan bool)   // Signal channel for completion

	// Producer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
			fmt.Printf("Produced job %d", i)
			time.Sleep(200 * time.Millisecond)
		}
		close(jobs) // Signal no more jobs
	}()

	// Consumer goroutine
	go func() {
		for job := range jobs {
			fmt.Printf("Processing job %d", job)
			time.Sleep(200 * time.Millisecond)
		}
		done <- true // Signal processing complete
	}()
	
	<-done // Wait for consumer to finish
	fmt.Println("All jobs processed!")
	fmt.Println("Task 7 completed!")
}

// Task 8: Goroutine Race Condition (Unsafe)
// Demonstrates race conditions and why they're dangerous
func task8RaceConditionUnsafe() {
	fmt.Println("=== Task 8: Goroutine Race Condition (Unsafe) ===")
	
	counter := 0
	var wg sync.WaitGroup
	
	increment := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter++ // This is unsafe! Multiple goroutines accessing shared data
		}
	}
	
	// Start 3 goroutines
	wg.Add(3)
	go increment()
	go increment()
	go increment()
	
	wg.Wait()
	fmt.Printf("Final counter (unsafe): %d (should be 3000)", counter)
	
	// 👉 "counter++ isn't atomic, so multiple goroutines cause a race condition and we get the wrong result.
	// 👉 A mutex ensures only one goroutine updates counter at a time, making the operation safe." ✅
	// Atomic means an operation happens completely or not at all — it can't be interrupted in the middle.
	fmt.Println("Task 8 completed!")
}

// Task 9: Safe Concurrent Counter with Mutex
// Shows how to fix race conditions using mutex for thread safety
func task9SafeCounterMutex() {
	fmt.Println("=== Task 9: Safe Concurrent Counter with Mutex ===")
	
	counter := 0
	var mu sync.Mutex // Mutex to protect shared data
	var wg sync.WaitGroup

	safeInc := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()   // Lock before accessing shared data
			counter++   // Safe operation - only one goroutine can access at a time
			mu.Unlock() // Unlock after operation
		}
	}

	wg.Add(2)
	go safeInc()
	go safeInc()

	wg.Wait()
	
	fmt.Printf("Final counter (safe): %d (should be 2000)", counter)
	fmt.Println("Task 9 completed!")
}

func main() {
	fmt.Println("🚀 Goroutine Learning Tasks")
	fmt.Println("===========================")
	
	// Execute all tasks in sequence
	task1BasicWaitGroup()
	task2BasicGoroutine()
	task3MultipleGoroutines()
	task4WaitGroupProper()
	task5SimpleChannels()
	task6ChannelReturn()
	task7ProducerConsumer()
	task8RaceConditionUnsafe()
	task9SafeCounterMutex()
	
	fmt.Println("🎉 All goroutine tasks completed successfully!")
	fmt.Println("📚 Key Concepts Learned:")
	fmt.Println("- Basic goroutine creation with 'go' keyword")
	fmt.Println("- Synchronization with sync.WaitGroup")
	fmt.Println("- Channel communication and patterns")
	fmt.Println("- Race conditions and thread safety")
	fmt.Println("- Mutex for protecting shared data")
}

// 📝 Notes on sync.WaitGroup:
// A sync.WaitGroup is like a counter that tracks how many goroutines are still running.
//
// Add(n) → increases the counter by n
// Done() → decreases the counter by 1  
// Wait() → waits until the counter goes back to 0
//
// Best Practices:
// - Always call Add() before starting the goroutine
// - Use defer wg.Done() to ensure Done() is called even if panic occurs
// - Never call Add() inside the goroutine itself (leads to race conditions)
