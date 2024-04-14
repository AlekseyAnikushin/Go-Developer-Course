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
	var s5 [5]string
	fmt.Println(s5)
}

func task2() {
	var s4 [4]string = [4]string{"яблоко", "груша", "слива", "абрикос"}
	fmt.Println(s4)
}

func task3() {
	var s4 [4]string = [4]string{"яблоко", "груша", "помидор", "абрикос"}
	s4[2] = "персик"
	fmt.Println(s4)
}

func task4() {
	s := []int{5, 2, 8, 3, 1, 9}
	fmt.Println(s)
}

func task5() {
	s := make([]int, 10)
	fmt.Println(s)
}

func task6() {
	s := make([]int, 10)
	s = append(s, 4, 1, 8, 9)
	fmt.Println(s)
}

func task7() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func task8() {
	s := []int{1, 2, 3, 4, 5, 6}
	s = append(s[:3], s[4:]...)
	fmt.Println(s)
}
