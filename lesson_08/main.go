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
}

func task1() {
	m := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	fmt.Println(m)
}

func task2() {
	m := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	i, ok := m["слон"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)", "слон", i, ok)
	fmt.Println()

	i, ok = m["бегемот"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)", "бегемот", i, ok)
	fmt.Println()

	i, ok = m["осьминог"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)", "осьминог", i, ok)
	fmt.Println()
}

func task3() {
	m := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	delete(m, "бегемот")
	fmt.Println(m)
}

func task4() {
	m := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	m["выдра"] = struct{}{}
	fmt.Println(m)
}

func task5() {
	m := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	m["бегемот"] = 2
	fmt.Println(m)
}
