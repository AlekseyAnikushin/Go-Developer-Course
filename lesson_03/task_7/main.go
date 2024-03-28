package main

import "fmt"

func main() {
	const (
		twoPow0 = 1 << iota
		twoPow1
		twoPow2
		twoPow3
		twoPow4
	)

	fmt.Println(twoPow0, twoPow1, twoPow2, twoPow3, twoPow4)
}
