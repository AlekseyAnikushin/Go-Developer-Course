package main

import (
	"fmt"
	"strings"
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
}

func fruitMarket(fruit string, m map[string]int) int {
	v, ok := m[fruit]

	if ok {
		return v
	}

	return -1
}

func task1() {
	m := map[string]int{
		"апельсины": 5,
		"яблоки":    3,
		"сливы":     1,
		"груши":     0,
	}

	var fruit string
	fmt.Println("Введите название фруктов:")
	fmt.Scanf("%s\n", &fruit)

	if quant := fruitMarket(fruit, m); quant >= 0 {
		fmt.Println("Количество: ", quant)
		return
	}

	fmt.Println("Не найдено.")
}

func task2() {
	s := []int{1, 2, 3}
	l := len(s)

	for i1 := 0; i1 < l; i1++ {
		fmt.Printf("v1: %d\n", s[i1])

	STOP:
		for i2 := 0; i2 < l; i2++ {
			fmt.Printf("%sv2: %d\n", strings.Repeat(" ", 4), s[i2])

			for i3 := 0; i3 < l; i3++ {
				fmt.Printf("%sv3: %d\n", strings.Repeat(" ", 8), s[i3])

				for i4 := 0; i4 < l; i4++ {
					fmt.Printf("%sv4: %d\n", strings.Repeat(" ", 12), s[i4])

					if i4 == 1 {
						break STOP
					}
				}
			}
		}
	}
}

func checkFood(food string) {
	switch food {
	case "груша", "яблоко", "апельсин":
		fmt.Println("Это фрукт")
	case "тыква", "огурец", "помидор":
		fmt.Println("Это овощ")
	default:
		fmt.Println("Что-то странное...")
	}
}

func task3() {
	var food string
	fmt.Println("Введите название еды:")
	fmt.Scanf("%s\n", &food)

	checkFood(food)
}
