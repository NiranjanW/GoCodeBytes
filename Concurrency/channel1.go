package Concurrency

import (
	"fmt"
	"time"
)



func chan1(){
	resCh := make(chan bool)

	go func(ch chan<- bool) {
		time.Sleep(time.Second)
		ch <- true
	}(resCh)

	res := <-resCh
	fmt.Print(res)
}