package Concurrency

import (
	"fmt"
	"sync"
	"time"
)

func channel2(){

	resCh := make(chan string ,2)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func( ch chan <- string , wg *sync.WaitGroup){
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch <- "Hello"
	}(resCh ,&waitGroup)

	go func(ch chan<- string, wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch <- "World"
	}(resCh, &waitGroup)

	waitGroup.Wait()

	for len(resCh) > 0 {
		res := <- resCh
		fmt.Println(res)
	}
}
