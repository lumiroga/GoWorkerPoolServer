package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("5", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int)
	m["Seis"] = 6

	s := []int{1, 2, 3}
	s = append(s, 16)

	for index, value := range s {
		fmt.Println("index=>", index, "value=>", value)
	}

	c := make(chan int)
	//go doSomething(c)
	fmt.Println(<-c)

	g := 25
	fmt.Println(g)

	h := &g
	fmt.Println(h)
}

func doSomething(c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
	c <- 1
}
