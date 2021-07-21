package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 2) //Buffer size two limits the capacity to two goroutines at the same time
	var wg sync.WaitGroup

	for i := 0; i < 7; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Println("Intitiating process...", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizing process", i)
	<-c
}
