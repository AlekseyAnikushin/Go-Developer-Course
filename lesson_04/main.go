package main

import (
	"fmt"
)

func main() {
	fmt.Println("Task 1 output:")
	task1()
	fmt.Println()

	fmt.Println("Task 2 output:")
	task2()
	fmt.Println()

	fmt.Println("Task 3 output:")
	task3()
	fmt.Println()

	fmt.Println("Task 4 output:")
	task4()
	fmt.Println()

	fmt.Println("Task 5 output:")
	task5()
	fmt.Println()

	fmt.Println("Task 6 output:")
	task6()
	fmt.Println()
}

// task 1
func hello() {
	fmt.Println("Hello, Go!")
}

func task1() {
	hello()
}

// task 2
func task2() {
	func() {
		fmt.Println("Hello, Go!")
	}()
}

// task 3
func hello3(f func()) {
	f()
}

func task3() {
	f := func() {
		fmt.Println("Hello, Go!")
	}

	hello3(f)
}

// task 4
func hello4() (f func()) {
	f = func() {
		fmt.Println("Hello, Go!")
	}

	return
}

func task4() {
	f := hello4()

	f()
}

// task 5
func task5() {
	defer fmt.Println("Завершение работы")

	hello()
}

// task 6
func printFibonacci(x int, y int, i int) {
	if i > 23 {
		return
	}

	fmt.Printf("%d: %d", i, x+y)
	fmt.Println()

	i++

	printFibonacci(y, x+y, i)
}

func task6() {
	fmt.Println("1: 0")
	fmt.Println("2: 1")
	printFibonacci(0, 1, 3)
}
