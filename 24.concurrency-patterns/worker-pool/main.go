package main

import (
	"fmt"
	"sync"
	"time"
)

// A Worker Pool is a fixed number of goroutines (workers) all pulling tasks from a single input channel. When a task arrives, any available worker can pick it up and process it.

// Worker Pool consists of:

// A fixed number of worker goroutines
// A shared input channel for tasks
// Workers that continuously pull tasks from the channel
// Optional result channel for processed outputs


type Task struct {
	ID int 
	Data string
}

type Result struct {
	TaskID int
	Output string
}

func worker(id int, tasks <- chan Task, results chan <- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		time.Sleep(100 * time.Millisecond)

		//proces and the task and create result
		result := Result{
			TaskID: task.ID,
			Output: fmt.Sprintf("Worker %d processed task %d: %s", id, task.ID, task.Data),
		}

		results <- result

		fmt.Printf("Worker %d completed task %d \n", id, task.ID)
	}
	fmt.Printf("Worker %d finished\n", id)
}


func main () {
	const numWorkers = 3
	const numTasks = 8

	// bidirectional channels
	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	var wg sync.WaitGroup

	// Start workers - channels are converted to directional automatically
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i,tasks, results,  &wg)
	}

	// send tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i, Data: fmt.Sprintf("data-%d", i)}
	}

	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	//results
	fmt.Println("\nResults:")
	for result := range results {
		fmt.Println(result.Output)
	}
}