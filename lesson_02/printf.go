package main

import "fmt"

func main() {
	fmt.Println("Делим 16 на 3")
	fmt.Printf("Результат: %v, остаток от деления: %v, тип результата: %[1]T", 16/3, 16%3)
}
