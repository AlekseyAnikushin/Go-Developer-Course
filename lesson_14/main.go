package main

import (
	"fmt"
	"time"
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

	fmt.Println("Task 5 output:")
	task5()
	fmt.Println()

	fmt.Println("Task 6 output:")
	task6()
	fmt.Println()
}

func task1() {
	go func() {
		fmt.Println("Привет из дочерней горутины!")
	}()

	time.Sleep(100 * time.Millisecond)
}

func task2() {
	ch := make(chan string)

	go func() {
		ch <- "Привет, строковый канал!"
	}()

	fmt.Println(<-ch)
}

func task3() {
	ch := make(chan string, 4)

	ch <- "Привет"
	ch <- "буферизованный канал!"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func task4() {
	ch := make(chan int)
	close(ch)
	stop := make(chan struct{})
	go func() {
		<-ch
		stop <- struct{}{}
	}()
	<-stop
	fmt.Println("happy end")
}

func task5() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)

	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()

	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}

	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()

	time.Sleep(time.Second)

	fmt.Println("завершение работы главной горутины")
}

func task6() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)

			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}
