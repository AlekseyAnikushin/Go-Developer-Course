package main

import "fmt"

func main() {
	const (
		myLocalConst1 = (iota + 1) * 10
		myLocalConst2
		myLocalConst3
		myLocalConst4
		myLocalConst5
	)

	fmt.Println(myLocalConst1, myLocalConst2, myLocalConst3, myLocalConst4, myLocalConst5)
}
