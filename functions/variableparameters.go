package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func getValues(x int) (db int, tp int, qd int) {
	db = 2 * x
	tp = 3 * x
	qd = 4 * x
	//Don't need to specify the return variables
	return
}

func main() {

	fmt.Println(sum(2, 3, 4, 6, 5, 43, 2, 13, 12, 3, 23, 3, 2))

	fmt.Println(getValues(73))
}
