package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) //Adds one 'buffer' to the wait group
		go doSomething(i, &wg)
	}

	wg.Wait() //stays active until counter is equals to zero
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done() //Decrements counter ('buffer') by one
	fmt.Println("Beginning of process....", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End process...", i)
}
