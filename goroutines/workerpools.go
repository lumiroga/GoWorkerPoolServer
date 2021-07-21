package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started Fib with number: %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d finished Fib with number: %d and result %d\n", id, job, fib)

		results <- fib
	}

}

func main() {
	tasks := []int{2, 3, 4, 5, 6, 10, 31, 32}
	nWorkers := 3 //qty of operations at the same time
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}

}
