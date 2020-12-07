package Concurrency

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func channel3() {

	bufferSize := 5
	wordCh := make(chan string , bufferSize)
	words := strings.Split("the quick brown fox jumped over the lazy dog" , " ")

	var waitGroup sync.WaitGroup

	for _ , word := range words {
		waitGroup.Add(1)

		go func(ch chan <- string , wg *sync.WaitGroup,someWord string) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			ch <- word
		}(wordCh , &waitGroup,word)
	}
	done := make(chan bool)
	go func(d chan bool) {
		defer close(d)
		for res := range wordCh {
			fmt.Print(fmt.Sprintf("%s ", res))
		}
	}(done)

	waitGroup.Wait()
	close(wordCh)
	<-done

}