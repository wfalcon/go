// пример работы горутин, каналов и waitgroup
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gofer(i int, ch chan int, wg *sync.WaitGroup, gc int) {
	defer wg.Done()
	sl := time.Duration(rand.Intn(gc))
	time.Sleep(sl * time.Second)
	fmt.Printf("Я горутина № %v ,работала %v sec \n", i+1, int(sl))
	ch <- i

}

func main() {
	gc := 100
	var ch = make(chan int, gc)
	wg := &sync.WaitGroup{}

	for i := 0; i < gc; i++ {
		wg.Add(1)
		go gofer(i, ch, wg, gc)

	}
	wg.Wait()
	close(ch)
	for i := range ch {

		fmt.Printf("Горутина № %v закончила работу \n", i+1)

	}

}
