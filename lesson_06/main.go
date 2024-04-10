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
}

type contract struct {
	ID     int
	Number string
	Date   string
}

func task1() {
	c := contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v \n", c)
}

func task2() {
	c := contract{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v \n", c)
}

func (c contract) print() string {
	return fmt.Sprintf("Договор № %s от %s", c.Number, c.Date)
}

func task3() {
	c := contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Println(c.print())
}

func task4() {
	type (
		contacts struct {
			Addss string
			Phone string
		}

		user struct {
			ID   int
			Name string
			contacts
		}

		employee struct {
			ID   int
			Name string
			contacts
		}
	)

	u := user{
		contacts: contacts{
			Addss: "UserAdress",
			Phone: "UserPhone",
		},
	}

	e := employee{
		contacts: contacts{
			Addss: "EmpAdress",
			Phone: "EmpPhone",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
