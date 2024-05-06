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
}

// --------------------- TASK 1 ------------------------------------------
func do1(v any) {
	a := 10

	/*
		Здесь нужно увеличить значение a на v.
		В случае невозможности приведения к int
		необходимо сообщить об этом и немедленно
		завершить полнение программы.
	*/

	i, ok := v.(int)

	if !ok {
		panic("Невозможно привести значение к Int")
	}

	a += i

	fmt.Println(a)
}

func task1() {
	a := 1
	do1(a)
}

// --------------------- TASK 2 ------------------------------------------
type Bird interface {
	Fly()
}

type Duck struct{}

func (d Duck) Fly() {
	fmt.Println("Утка - Умею летать!")
}

func (d Duck) Swim() {
	fmt.Println("Утка - Умею плавать!")
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Воробей - Умею летать!")
}

func Do2(b Bird) {
	b.Fly()
	//здесь нужно дописать код
	switch b.(type) {
	case Duck:
		b.(Duck).Swim()
	default:
	}
}

func task2() {
	var d, s Bird
	d = Duck{}
	Do2(d)
	s = Sparrow{}
	Do2(s)
}

// --------------------- TASK 3 ------------------------------------------
type Reporter interface {
	Format()
}

type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}

func Report(x Reporter) {
	x.Format()
}

func task3() {
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}
