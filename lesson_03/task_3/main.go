package main

import "fmt"

const myGlobalConst = 10

func main() {
	const myGlobalConst = 20

	fmt.Println(myGlobalConst)
}
