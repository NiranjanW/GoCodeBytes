package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	//
	//	// --
	//	// The topic is passed as a pointer to the Producer, so we can't
	//	// use a hard-coded literal. A:q:nd a variable is a nicer way to do
	//	// it anyway ;-)
	topic := "order.fulfillment"
	//
	//	// --
	//	// Create Producer instance
	//	// https://docs.confluent.io/current/clients/confluent-kafka-go/index.html#NewProducer
	//	//
	//	// Variable p holds the new Producer instance. There's a bunch of config that we _could_
	//	// specify here but the only essential one is the bootstrap server.
	//	//
	//	// Note that we ignore any error that might happen, by passing _ as the second variable
	//	// on the left-hand side. This is, obviously, a terrible idea.
	//
	p, _ := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092"})
	//
	//	// --
	//	// Send a message using Produce()
	//	// https://docs.confluent.io/current/clients/confluent-kafka-go/index.html#Producer.Produce
	//	//
	//	// Only essential values are specified here - the topic, partition, and value
	//	//
	//	// There is NO handling of errors, timeouts, etc - we just fire & forget this message.
	//	// Did it work? ¯\_(ツ)_/¯

	var lines = readFromFile()
	//fmt.Println(lines)
	a, _ := regexp.Compile(`\d*:`)
	for _, line := range lines{
		if len(line) == 0 {
			 errors.New("blank line")
		}
		key :=(strings.Split(line ,":"))[0]
			//}else {
		//	value := a.Split(line, 2)[1]
		//}
		second := a.Split(line, 2)
		if second == nil {
			errors.New("blank line")
			os.Exit(1)
		}


		value := a.Split(line, 2)[1]

		//fmt.Printf("key %v \n, value %v" ,key,value)
		//fmt.Printf("key %v , value %v" , key, value)

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic,
				Partition: 0},
				Key: []byte(key),
			Value: []byte(value)}, nil)
	}

	//
	//	//// --
	//	//// Close the producerG
	p.Flush(1)
	p.Close()
	//

	//
	//
}

func readFromFile() []string {
	absPath, _ := filepath.Abs("./Data/order.fulfillment.txt")
	file ,err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	//kafka:= make([]string,600)
	t := []string{}
	for {
		line ,err := reader.ReadString('\n')
		t = append(t,line)
		if err != nil {
			if err ==io.EOF {
				break
			} else {
				fmt.Print(err)
			}
		}
		//row := strings.Split(line, ":")


		//value :=  row[1]
		//fmt.Printf(" kafka  value %v ", value )

	}
	return t
}
