package main

import "fmt"

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

	fmt.Println("Task 7 output:")
	task7()
	fmt.Println()

	fmt.Println("Task 8 output:")
	task8()
	fmt.Println()
}

func task1() {
	s := ""
	var p *string = &s
	fmt.Println(p)
}

func task2() {
	i := 1
	fmt.Printf("%d, %d", i, &i)
	fmt.Println()
}

func task3() {
	b := false
	p := &b
	fmt.Println(b)
	*p = true
	fmt.Println(b)
}

func task4() {
	task2()
}

func change(p *int) {
	*p *= 2
}

func task5() {
	i := 2
	fmt.Println(i)
	change(&i)
	fmt.Println(i)
}

func task6() {
	type square int
	var s square = 25
	fmt.Println(s)
}

func task7() {
	type square int
	var s square = 30
	s += 15
	fmt.Println(s)
}

type square int

func strSquare(s square) string {
	return fmt.Sprintf("%d м%c", s, 0x00B2)
}

func task8() {
	var s square = 34
	s += 10
	fmt.Println(strSquare(s))
}
