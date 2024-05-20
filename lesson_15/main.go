package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
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
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("горутина: ", i)
		}()
	}
	wg.Wait()
}

func task2() {
	var n int32

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&n, 1)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(n)
}

func task3() {
	var mu sync.Mutex

	var n int

	for i := 0; i < 200; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			n++
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(n)
}

func start() {
	fmt.Println("Start")
}

func task4() {
	var once sync.Once

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			once.Do(start)
			fmt.Println(i)
		}()
	}

	time.Sleep(2 * time.Second)
}

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type temp struct {
	tmp string
	mu  sync.RWMutex
}

func (t *temp) ReadTemp() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.tmp
}

func (t *temp) ChangeTemp(v string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tmp = v
}

func task5() {
	t := temp{tmp: "empty"}
	for i := 10; i < 16; i++ {
		go func() {
			t.ChangeTemp(strconv.Itoa(i))
		}()
		go func() {
			fmt.Println(t.ReadTemp())
		}()
	}
	time.Sleep(2 * time.Second)
}
