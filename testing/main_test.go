package main

import "testing"

func TestSum(t *testing.T) {
	/*total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Invalid test - Unit test failed")
	} else {
		t.Log("Test passed")
	}*/
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{10, 12, 22},
		{6, 5, 11},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)

		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 1, 1},
		{1, 2, 2},
		{100, 1, 100},
		{7000, 1231, 7000},
		{2, 1, 2},
	}
	for _, item := range tables {
		max := GetMax(item.a, item.b)

		if max != item.n {
			t.Errorf("Maximum is incorrect %d expected %d", max, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{40, 102334155},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)

		if fib != item.n {
			t.Errorf("Fibonacci sequence is wrong, got %d, expected %d", fib, item.n)
		}
	}
}
