package Concurrency

import (
	"log"
	"time"
)

type Item struct {
	name string
}
func channel0(){

for i :=0; i < 1000; i++ {
	go doSomething(nil)
}
itemChan := make(chan Item)
	for {
		errChan := make(chan error)
		select {
	case item := <- itemChan:
		doSomething(item)
	case err := <- errChan:
		handleError(err)
	case <- time.After(1 * time.Minute):
		log.Println("nothing happened for 1 min")
	}
}
}

func handleError(err error) string{
	return err.Error()
}
func doSomething(interface{}) string {
	return "Hello World"
}





