// пример работы горутин, каналов и waitgroup
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gofer(i int, start chan int, wg *sync.WaitGroup, gc int) {
	defer wg.Done()
	sl := time.Duration(rand.Intn(gc))
	time.Sleep(sl * time.Second)
	fmt.Printf("Я горутина № %v ,работала %v sec \n", i+1, int(sl))
	start <- i

}

func main() {
	gc := 100
	var start = make(chan int, gc)
	wg := &sync.WaitGroup{}

	for i := 0; i < gc; i++ {
		wg.Add(1)
		go gofer(i, start, wg, gc)

	}
	wg.Wait()
	close(start)
	for i := range start {

		fmt.Printf("Горутина № %v закончила работу \n", i+1)

	}

}
