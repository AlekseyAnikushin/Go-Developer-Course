package main

import (
	"context"
	"fmt"
	"sync"
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
}

func task1() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Завершена горутина ", i)
					return
				default:
					fmt.Println("Работает горутина ", i)
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cancel()
	}()

	wg.Wait()
}

func task2() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Завершена горутина ", i)
					return
				default:
					fmt.Println("Работает горутина ", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}

	wg.Wait()
}

func task3() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Завершена горутина ", i)
					return
				default:
					fmt.Println("Работает горутина ", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}

	wg.Wait()
}

func task4() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "some key1", "some value1")
	ctx = context.WithValue(ctx, "some key2", "some value2")
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(ctx.Value("some key1"))
		fmt.Println(ctx.Value("some key2"))
	}()

	wg.Wait()
}

func task5() {
	var ch chan int
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if ch != nil {
					fmt.Println("stop горутина: ", i)
					return
				}
				time.Sleep(time.Second)
				fmt.Println("сложные вычисления горутины: ", i)
			}
		}()
	}

	wg.Add(1)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		ch = make(chan int)
		wg.Done()
	}()

	wg.Wait()
}
