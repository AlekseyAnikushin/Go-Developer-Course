package main

import (
	"errors"
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
}

func task1() {
	err := errors.New("Ошибка1")
	err = fmt.Errorf("Ошибка2:%w", err)
	err = fmt.Errorf("Ошибка3:%w", err)
	fmt.Println(err)
}

func task2() {
	err := errors.New("Ошибка1")
	err = fmt.Errorf("Ошибка2:%w", err)
	err = fmt.Errorf("Ошибка3:%w", err)
	fmt.Println(errors.Unwrap(err))
}

func task3() {
	err1 := errors.New("Ошибка1")
	err := fmt.Errorf("Ошибка2:%w", err1)
	err = fmt.Errorf("Ошибка3:%w", err)
	if errors.Is(err, err1) {
		fmt.Printf("Ошибка %v была\n", err1)
	} else {
		fmt.Printf("Ошибки %v не было\n", err1)
	}
}

type myFirstError struct {
	message string
}

func (e myFirstError) Error() string {
	return e.message
}

func task4() {
	err := errors.New("Ошибка1")
	err = fmt.Errorf("Ошибка2:%w", err)
	err = fmt.Errorf("Ошибка3:%w", err)

	fmt.Println("Ошибка myFirstError была:", errors.As(err, new(myFirstError)))
}
