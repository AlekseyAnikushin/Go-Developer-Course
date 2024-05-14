package main

import (
	"encoding/json"
	"encoding/xml"
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

type contract1 struct {
	Number   int
	Landlord string
	Tenat    string `json:"tenat"`
}

func task1() {
	s := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`

	c := contract1{}

	err := json.Unmarshal([]byte(s), &c)

	if err != nil {
		fmt.Println("Ошибка распаковки Json: ", err)
	}

	fmt.Printf("%+v\n", c)
}

type contract struct {
	Number   int
	Landlord string
	Tenat    string `json:"tenat" xml:"tenat"`
}

func task2() {
	c := contract{
		Number:   2,
		Landlord: "Остап",
		Tenat:    "Шура",
	}

	s, err := json.Marshal(c)

	if err != nil {
		fmt.Println("Ошибка формирования Json: ", err)
	}

	fmt.Println(string(s))
}

type contract3 struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts struct {
	List []contract3 `xml:"contract"`
}

func task3() {
	x := `<?xml version="1.0" encoding="UTF-8"?>
	<contracts>
	<contract>
	<number>1</number>
	<sign_date>2023-09-02</sign_date>
	<landlord>Остап</landlord>
	<tenat>Шура</tenat>
	</contract>
	<contract>
	<number>2</number>
	<sign_date>2023-09-03</sign_date>
	<landlord>Бендер</landlord>
	<tenat>Балаганов</tenat>
	</contract>
	</contracts>`

	c := contracts{}

	err := xml.Unmarshal([]byte(x), &c)

	if err != nil {
		fmt.Println("Ошибка распаковки Xml: ", err)
	}

	fmt.Printf("%+v\n", c)
}

type contracts4 struct {
	List []contract `xml:"contract"`
}

func task4() {
	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	}

	c4 := contracts4{}
	c4.List = append(c4.List, c)

	x, err := xml.MarshalIndent(c4, "", "  ")

	if err != nil {
		fmt.Println("Ошибка формирования Xml: ", err)
	}

	fmt.Println(string(x))
}
